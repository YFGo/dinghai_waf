package biz

import (
	"context"
	"errors"
	"math"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"wafconsole/utils/code"
	"wafconsole/utils/const/user"
	up "wafconsole/utils/plugin"
)

type WafUserCommonRepo interface {
	SaveKVToRs(ctx context.Context, userEmail, systemCode string, expiration time.Duration) error
	SendSingEmailCode(ctx context.Context, content, userEmail string) error
	DeleteKvFromRs(ctx context.Context, userEmail string) error
}

type WafUserCommonUsecase struct {
	captchaStore sync.Map
	repo         WafUserCommonRepo
	wafUserRepo  WafUserRepo
	log          *log.Helper
}

func NewWafUserCommonUsecase(repo WafUserCommonRepo, wafUserRepo WafUserRepo, logger log.Logger) *WafUserCommonUsecase {
	return &WafUserCommonUsecase{
		repo:         repo,
		wafUserRepo:  wafUserRepo,
		captchaStore: sync.Map{},
		log:          log.NewHelper(logger),
	}
}

func (w *WafUserCommonUsecase) RefreshAccessToken(refreshToken string) (string, string, int64, error) {
	// 通过 refreshToken 刷新accessToken
	jwts := up.InitNewJWTUtils()
	parserRefreshToken, isUpd, err := jwts.ParseRefreshToken(refreshToken)
	if err != nil {
		return "", "", 0, status.Error(codes.Unknown, err.Error())
	}
	if isUpd {
		accessToken, refreshTokenNew, expiresAt := jwts.GetToken(parserRefreshToken.UserClaims)
		return accessToken, refreshTokenNew, expiresAt, nil
	}
	return "", "", 0, status.Error(codes.Canceled, "refreshToken is expired")
}

// GetCaptcha 获取图片验证码
func (w *WafUserCommonUsecase) GetCaptcha(ctx context.Context) (string, string, string, error) {
	rotateCaptcha := code.NewRotateCaptchaStrategy()
	captchaID, masterImg, thumbImg, captchaCache, err := rotateCaptcha.Generate()
	if err != nil {
		w.log.WithContext(ctx).Error(err)
		return "", "", "", err
	}
	// 存储正确角度
	w.captchaStore.Store(captchaID, captchaCache)
	return captchaID, masterImg, thumbImg, nil
}

// VerifyCaptchaInfo 校验验证码逻辑
func (w *WafUserCommonUsecase) VerifyCaptchaInfo(ctx context.Context, captchaId string, userAngle float64) error {
	// 从缓存中获取
	captchaCache, exists := w.captchaStore.Load(captchaId)
	if !exists {
		w.log.WithContext(ctx).Error(codes.NotFound)
		return status.Error(codes.NotFound, "<UNK>")
	}
	captchaCacheValue, ok := captchaCache.(*code.CaptchaCacheValue)
	if !ok {
		w.log.WithContext(ctx).Error(codes.Unknown)
		return status.Error(codes.Unknown, "<UNK>")
	}
	w.captchaStore.Delete(captchaId)
	if time.Now().After(captchaCacheValue.ExpireTime) { // 验证码过期
		return status.Error(codes.Canceled, "expired")
	}
	success := math.Abs(userAngle-captchaCacheValue.CorrectAngle) <= 5
	if !success { // 验证未通过
		w.log.WithContext(ctx).Error(codes.PermissionDenied)
		return status.Error(codes.PermissionDenied, "captcha code is error")
	}
	return nil
}

// SendCode 发送验证码
func (w *WafUserCommonUsecase) SendCode(ctx context.Context, userEmail, sendAction string) error {
	// 判断如果是登录行为发送的验证码 , 验证邮箱是否存在
	if sendAction == user.LoginActionSend {
		_, err := w.wafUserRepo.GetUserInfoByEmail(ctx, userEmail)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) { // 用户不存在
				return status.Error(codes.NotFound, "user not found")
			}
			w.log.WithContext(ctx).Error(err)
			return err
		}

	}
	// 1. 生成验证码
	emailCode := code.NewEmailCode()
	codeInfo := emailCode.Generate()
	// 2. 先把验证码存入redis
	if err := w.repo.SaveKVToRs(ctx, userEmail, codeInfo, 5*time.Minute); err != nil {
		w.log.WithContext(ctx).Error(err)
		return err
	}
	// 3. 发送邮件
	emailBody, err := code.GenerateEmailBody(codeInfo) // 邮件模板
	if err != nil {
		w.log.WithContext(ctx).Error(err)
		return err
	}
	if err := w.repo.SendSingEmailCode(ctx, emailBody, userEmail); err != nil {
		w.log.WithContext(ctx).Error(err)
		return err
	}
	return nil
}

// SseConnect 实现SSE连接
func (w *WafUserCommonUsecase) SseConnect(ctx context.Context, userId int64) error {
	// 这个方法只返回成功，实际的SSE处理由HTTP层完成
	return nil
}

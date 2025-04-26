package code

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
	"log/slog"
	"time"

	"github.com/wenlng/go-captcha-assets/resources/images"
	"github.com/wenlng/go-captcha/v2/rotate"
)

// CaptchaCacheValue 验证码缓存信息
type CaptchaCacheValue struct {
	CorrectAngle float64
	ExpireTime   time.Time
}

// CaptchaStrategy 验证码策略接口
type CaptchaStrategy interface {
	Generate() (interface{}, string, string, error)
}

// RotateCaptchaStrategy 旋转类验证码
type RotateCaptchaStrategy struct {
	capt rotate.Captcha
}

func NewRotateCaptchaStrategy() *RotateCaptchaStrategy {
	builder := rotate.NewBuilder()
	imgs, err := images.GetImages()
	if err != nil {
		slog.Error("get images is failed: ", err)
		return nil
	}

	builder.SetResources(
		rotate.WithImages(imgs),
	)
	return &RotateCaptchaStrategy{
		capt: builder.Make(),
	}
}

func (r *RotateCaptchaStrategy) Generate() (string, string, string, *CaptchaCacheValue, error) {
	captData, err := r.capt.Generate()
	if err != nil {
		return "", "", "", nil, err
	}
	blockData := captData.GetData()
	if blockData == nil {
		return "", "", "", nil, errors.New("<UNK>")
	}

	masterImg, err := captData.GetMasterImage().ToBase64() // 主题图片生成
	if err != nil {
		return "", "", "", nil, err
	}
	thumbImg, err := captData.GetThumbImage().ToBase64() // 缩略图生成
	if err != nil {
		return "", "", "", nil, err
	}
	captchaID := uuid.New().String()
	return captchaID, masterImg, thumbImg, &CaptchaCacheValue{
		CorrectAngle: float64(blockData.Angle),
		ExpireTime:   time.Now().Add(5 * time.Minute),
	}, nil
}

// EmailCode 邮件验证码
type EmailCode struct {
	numberCode string
}

func NewEmailCode() *EmailCode {
	return &EmailCode{
		numberCode: "",
	}
}

func (e *EmailCode) Generate() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	code := fmt.Sprintf("%06d", rand.Intn(1000000))
	return code
}

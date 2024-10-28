package biz

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"wafconsole/app/user/internal/server/plugin"
)

type WafUserCommonRepo interface {
}

type WafUserCommonUsecase struct {
	repo WafUserCommonRepo
}

func NewWafUserCommonUsecase(repo WafUserCommonRepo) *WafUserCommonUsecase {
	return &WafUserCommonUsecase{repo: repo}
}

func (w *WafUserCommonUsecase) RefreshAccessToken(refreshToken string) (string, string, int64, error) {
	// 通过 refreshToken 刷新accessToken
	jwts := plugin.InitNewJWTUtils()
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

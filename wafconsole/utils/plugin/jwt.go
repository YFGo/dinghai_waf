package utils

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"log/slog"
	"time"
)

const (
	AccessTokenExpire  = 3600
	RefreshTokenExpire = 7200
	AccessTokenKey     = "waf"
	RefreshToken       = "refresh"
)

type UserClaims struct {
	UserId   uint64 `json:"user_id"`
	Username string `json:"username"`
}

type CustomClaims struct {
	UserClaims
	jwtv4.RegisteredClaims //内嵌的标准声明
}

type JWTUtils struct {
	AccessSecret  []byte
	RefreshSecret []byte
	Timeout       int
	MaxRefresh    int
}

func InitNewJWTUtils() *JWTUtils {
	authMiddleware := &JWTUtils{
		AccessSecret:  []byte(AccessTokenKey),
		RefreshSecret: []byte(RefreshToken),
		Timeout:       AccessTokenExpire,
		MaxRefresh:    RefreshTokenExpire,
	}
	return authMiddleware
}

// GetToken 获取accessToken 和 RefreshToken
func (jm *JWTUtils) GetToken(userClaims UserClaims) (string, string, int64) {
	//accessToken 数据
	aT := CustomClaims{
		UserClaims: userClaims,
		RegisteredClaims: jwtv4.RegisteredClaims{
			Issuer:    "zyfLink",
			IssuedAt:  jwtv4.NewNumericDate(time.Now()),
			ExpiresAt: jwtv4.NewNumericDate(time.Now().Add(time.Hour * time.Duration(jm.Timeout))), //一小时过期
		},
	}
	//refreshToken数据
	rT := CustomClaims{
		UserClaims: userClaims,
		RegisteredClaims: jwtv4.RegisteredClaims{
			Issuer:    "zyfLink",
			IssuedAt:  jwtv4.NewNumericDate(time.Now()),
			ExpiresAt: jwtv4.NewNumericDate(time.Now().Add(time.Hour * time.Duration(jm.MaxRefresh))),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, aT)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rT)
	accessTokenSigned, err := accessToken.SignedString(jm.AccessSecret)
	if err != nil {
		fmt.Println("获取AccessToken失败 , Secret错误")
		return "", "", 0
	}
	refreshTokenSigned, err := refreshToken.SignedString(jm.RefreshSecret)
	if err != nil {
		fmt.Println("获取RefreshToken失败 , Secret错误")
		return "", "", 0
	}
	return accessTokenSigned, refreshTokenSigned, aT.ExpiresAt.Time.Unix()
}

// ParseRefreshToken 解析ParserRefreshToken
func (jm *JWTUtils) ParseRefreshToken(refreshTokenStr string) (*CustomClaims, bool, error) {
	refreshToken, err := jwt.ParseWithClaims(refreshTokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jm.RefreshSecret, nil
	})
	if err != nil {
		//生成refreshToken失败
		slog.Error("refreshToken解析失败")
		return nil, false, err
	}
	if claims, ok := refreshToken.Claims.(*CustomClaims); ok && refreshToken.Valid {
		return claims, true, nil
	}
	return nil, false, errors.New("invaild token")
}

// ParseAccessToken 解析ParseAccessToken
func (jm *JWTUtils) ParseAccessToken(accessTokenStr string) (*CustomClaims, bool, error) {
	accessToken, err := jwt.ParseWithClaims(accessTokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jm.AccessSecret, nil
	})
	if err != nil {
		var v *jwt.ValidationError
		errors.As(err, &v)
		if v.Errors == jwt.ValidationErrorExpired {
			slog.Error("accessToken已过期")
		} else {
			slog.Error("accessToken解析失败")
			return nil, false, err
		}

	}
	if claims, ok := accessToken.Claims.(*CustomClaims); ok && accessToken.Valid {
		return claims, false, nil
	}
	return nil, false, errors.New("invaild token")
}

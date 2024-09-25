package utils

import (
	"ConfessionWall/config/config"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type UserClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

var ErrTokenHandlingFailed = errors.New("token handling failed")

// GenerateToken 用于生成 JWT token
func GenerateToken(userID uint) (string, error) {
	lifespan, err := strconv.Atoi(config.Config.GetString("jwt.lifespan"))
	if err != nil {
		zap.L().Error("解析 JWT Lifespan 失败", zap.Error(err))
		return "", err
	}

	claims := UserClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(lifespan) * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                          // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                          // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Config.GetString("jwt.secret")))
	return tokenString, err
}

// ExtractToken 用于从 JWT token 中提取 user_id
func ExtractToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(config.Config.GetString("jwt.secret")), nil
	})
	if err != nil {
		zap.L().Error("解析 Token 失败", zap.String("token", tokenString), zap.Error(err))
		return 0, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, ErrTokenHandlingFailed
}

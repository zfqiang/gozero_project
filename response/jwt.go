package response

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// GenToken 生成JWT
func GenerateAccessToken(secret string, expire int64, userId int64, username, role string) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"userId":   userId,
		"username": username,
		"role":     role,
		"iat":      now.Unix(),
		"exp":      now.Add(time.Duration(expire) * time.Second).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(secret))
}

func GenerateRefreshToken(secret string, expire int64, userId int64, username string) (string, error) {
	claims := jwt.MapClaims{
		"userId":   userId,
		"username": username,
		"exp":      time.Now().Add(time.Duration(expire) * time.Second).Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(secret))
}

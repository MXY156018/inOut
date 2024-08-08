package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

// 用户 JWT
type UserClaims struct {
	//用户ID
	UserID int
	//缓存时间
	BufferTime int64
	jwt.StandardClaims
}

// 生成用户 JWT
func GetUsrToken(signingKey string, claims UserClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

// 解码游戏 JWT token
//
// token JWT token
func ParseUserToken(tokenString string, signingKey string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(_ *jwt.Token) (i interface{}, e error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

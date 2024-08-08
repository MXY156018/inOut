package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

// 用户 JWT
type MerchantClaims struct {
	//商户ID
	MchID int
	//是否有商品管理权限 0 无，按位与
	Privilege int
	//缓存时间
	BufferTime int64
	jwt.StandardClaims
	//是否登陆
	IsLogin int
}

// 生成用户 JWT
func GetMerchantToken(signingKey string, claims MerchantClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

// 解码游戏 JWT token
//
// token JWT token
func ParseMerchantToken(tokenString string, signingKey string) (*MerchantClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MerchantClaims{}, func(_ *jwt.Token) (i interface{}, e error) {
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
		if claims, ok := token.Claims.(*MerchantClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

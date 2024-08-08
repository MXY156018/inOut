package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type AdminClaims struct {
	//用户ID
	UserID int
	//权限ID
	AuthorityId string
	//缓存时间
	BufferTime int64
	// 商户ID(后台使用 7 为平台)
	MerchantId int
	jwt.StandardClaims
}

// 生成管理员 Token
func GetAdminToken(signingKey string, claims AdminClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signingKey))
}

// 解析 token
func ParseAdminToken(tokenString string, signingKey string) (*AdminClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AdminClaims{}, func(_ *jwt.Token) (i interface{}, e error) {
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
		if claims, ok := token.Claims.(*AdminClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}

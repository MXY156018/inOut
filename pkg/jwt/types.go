package jwt

import "errors"

var (
	//过期
	TokenExpired = errors.New("Token is expired")
	// 非法token
	TokenNotValidYet = errors.New("Token not active yet")
	// 非法token
	TokenMalformed = errors.New("That's not even a token")
	// 非法token
	TokenInvalid = errors.New("Couldn't handle this token:")
)

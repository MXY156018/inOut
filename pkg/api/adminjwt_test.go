package api_test

import (
	"fmt"
	"mall-pkg/api"
	mjwt "mall-pkg/jwt"
	"net/http"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type monitorHttpWrite struct {
}

func (l *monitorHttpWrite) Header() http.Header {
	return http.Header{}
}
func (l *monitorHttpWrite) Write(data []byte) (int, error) {
	fmt.Printf("%s", string(data))
	return len(data), nil
}

func (l *monitorHttpWrite) WriteHeader(statusCode int) {
	fmt.Println("status code", statusCode)
}

func Test_AdminJWT(t *testing.T) {
	var secret = "mall_jwt_admin:9nmjh3829hfsdflwe/<./de"
	claims := mjwt.AdminClaims{
		UserID:      int(1),
		AuthorityId: "888",
		BufferTime:  86400,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer: "mall",
		},
	}
	token, err := mjwt.GetAdminToken(secret, claims)
	if err != nil {
		t.Fatal(err)
	}
	
	dstClaims, err := mjwt.ParseAdminToken(token, secret)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(dstClaims.UserID, dstClaims.AuthorityId)

	jwtMidle := api.AdminJwt{
		Secret:      secret,
		ExpiresTime: 604800,
	}

	callNext := false
	m := jwtMidle.Middleware(func(w http.ResponseWriter, r *http.Request) {
		t.Log("next")
		callNext = true
	})
	req := &http.Request{
		Header: make(http.Header),
	}
	req.Header.Add(api.Req_Header_Token1, token)
	m(&monitorHttpWrite{}, req)
	if !callNext {
		t.Fatal("not call next")
	}
}

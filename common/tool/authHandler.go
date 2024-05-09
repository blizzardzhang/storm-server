package tool

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JwtPayLoad struct {
	UserId  string `json:"user_id"`
	Account string `json:"account"`
	Role    string `json:"role"`
}

type CustomerClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

// GenToken 创建token
func GenToken(user JwtPayLoad, accessSecret string, expires int64) (string, error) {
	claims := CustomerClaims{
		JwtPayLoad: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * time.Duration(expires))),
		},
	}
	//claims := make(jwt.MapClaims)
	//claims["exp"] = iat + seconds
	//claims["iat"] = iat
	//claims[ctxdata.CtxKeyJwtUserId] = userId
	//token := jwt.New(jwt.SigningMethodHS256)
	//token.Claims = claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(accessSecret))
}

// ParseToken 解析token
func ParseToken(tokenStr string, accessSecret string) (*CustomerClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomerClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomerClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token is invalid")
}

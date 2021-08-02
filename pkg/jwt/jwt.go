package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

const (
	ErrorreasonServerbusy   = "服务器繁忙"
	ErrorreasonLoginOutTime = "登录过期，请重新登录"
)

var (
	secret     string // 加盐
	expireTime int    // token有效期
	issuer     string
	subject    string
)

func init() {
	secret = viper.GetString("jwtConfig.Secret")
	expireTime = viper.GetInt("jwtConfig.ExpireTime")
	issuer = viper.GetString("jwtConfig.Issuer")
	subject = viper.GetString("jwtConfig.Subject")
}

type AuthJwt struct {
}
type Claims struct {
	UserInfo interface{}
	jwt.StandardClaims
}

func (*AuthJwt) CreateClaims(data interface{}) (claims *Claims) {
	claims = &Claims{
		UserInfo: data,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * time.Duration(expireTime)).Unix(),
			Issuer:    issuer,
			Subject:   subject,
		},
	}
	return
}
func (j *AuthJwt) CreateToken(data interface{}) (token string, err error) {
	claims := j.CreateClaims(data)
	token, err = j.GetToken(claims)
	return
}
func (*AuthJwt) GetToken(claims *Claims) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(secret))
	if err != nil {
		err = errors.New(ErrorreasonServerbusy)
	}
	return
}
func (*AuthJwt) Parse(strToken string) (claims *Claims, err error) {
	token, errs := jwt.ParseWithClaims(strToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if errs != nil {
		return nil, errs
	}
	claims, ok := token.Claims.(*Claims)
	if !ok {
		err = errors.New(ErrorreasonLoginOutTime)
		return nil, err
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, err
	}
	return claims, nil
}
func (j *AuthJwt) Refresh(token string) (newToken string, err error) {
	claims, err := j.Parse(token)
	if err != nil {
		return
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)
	newToken, err = j.GetToken(claims)
	if err != nil {
		return
	}
	return
}

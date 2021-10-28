package main

import (
	"fmt"
	"github.com/zngue/go_helper/pkg/config"
	"github.com/zngue/go_helper/pkg/jwt"
	"testing"
	"time"
)

type UserInfoToken struct {
	ID int
	Name string
	Password string
	Age int
	Sex int

}
func TestCreateJwt(t *testing.T) {

	config.NewConfig(config.Path("eg/conf"))

	j := new(jwt.AuthJwt)

	token, err := j.CreateToken(UserInfoToken{
		ID: 10,
		Name: "zhangsan",
		Password: "13230",
		Age: 12,
		Sex: 1,
	})
	parse, err := j.Parse(token)
	unixTime := time.Now().Unix()
	if unixTime > parse.StandardClaims.ExpiresAt {
		fmt.Println("expire")
	}

	if user, ok := parse.UserInfo.(map[string]interface{}); ok {
		fmt.Println(user)
	}else {
		fmt.Println(ok)
	}
	fmt.Println(token, err, parse, unixTime)

}

package main

import (
	"fmt"
	"github.com/zngue/go_helper/pkg/config"
	"github.com/zngue/go_helper/pkg/jwt"
	"testing"
	"time"
)

func TestCreateJwt(t *testing.T) {

	config.NewConfig(config.Path("eg/conf"))

	j := new(jwt.AuthJwt)

	token, err := j.CreateToken(map[string]string{
		"username": "zhangsan",
		"age":      "56",
		"password": "123456",
	})
	parse, err := j.Parse(token)
	unixTime := time.Now().Unix()
	if unixTime > parse.StandardClaims.ExpiresAt {
		fmt.Println("expire")
	}

	if user, ok := parse.UserInfo.(map[string]string); ok {
		fmt.Println(user)
	}
	fmt.Println(token, err, parse, unixTime)

}

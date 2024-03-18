package server

import (
	"context"
	"github.com/astaxie/beego/logs"
	"github.com/golang-jwt/jwt/v4"
	"way/lx/client"
)

type LoginRes struct {
	Token    string
	UserInfo interface{}
}

func GenToken(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	return token.SignedString([]byte("123123123"))
}
func UserLogin(ctx context.Context, username, password string) (*LoginRes, error) {
	getUsername, err := client.GetUsername(ctx, username, password)
	if err != nil {
		return nil, err
	}
	logs.Info(getUsername, "======getusername")
	token, err := GenToken(1)
	if err != nil {
		return nil, err
	}
	return &LoginRes{
		Token:    token,
		UserInfo: getUsername,
	}, nil
}

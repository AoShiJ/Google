package server

import (
	"context"
	demo "demo/lx/message"
	"demo/lx/model"
	"fmt"
)

func Login(ctx context.Context, username, password string) (int64, error) {
	user := new(model.User)
	u, _ := user.Select(username)
	if u.Password != password {
		return 0, fmt.Errorf("用户密码不正确")
	}
	return int64(u.ID), nil
}
func mToPd(user *model.User) *demo.LoginRequest {
	return &demo.LoginRequest{
		Username: user.Username,
		Password: user.Password,
	}
}

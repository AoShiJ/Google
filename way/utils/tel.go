package utils

import (
	"demo/lx/framework/redis"
	"github.com/astaxie/beego/logs"
	redis1 "github.com/go-redis/redis"
	"time"
)

func SendMobile(string2 string) {
	err := redis.WithClient(func(cli *redis1.Client) error {
		e := cli.Set(string2, 1000, time.Second*10).Err()
		logs.Info(e)
		return nil
	})
	if err != nil {
		return
	}
}

package redis

import (
	"demo/lx/framework/nacos"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	redis1 "github.com/go-redis/redis"
	"time"
)

type T struct {
	Redis struct {
		Ip   string `json:"Ip"`
		Port int    `json:"Port"`
	} `json:"Redis"`
}

func WithClient(hand func(cli *redis1.Client) error) error {
	config, err := nacos.GetConfig()
	if err != nil {
		return err
	}
	var t T
	json.Unmarshal([]byte(config), &t)
	logs.Info(config)
	cli := redis1.NewClient(&redis1.Options{Addr: fmt.Sprintf("%v:%v", t.Redis.Ip, t.Redis.Port)})
	defer cli.Close()
	err = hand(cli)
	if err != nil {
		return err
	}
	return nil
}
func SetNx(key string, val interface{}, time time.Duration) error {
	err := WithClient(func(cli *redis1.Client) error {
		err := cli.SetNX(key, val, time).Err()
		if err != nil {
			return err
		}
		return err
	})
	if err != nil {
		return err
	}
	return nil
}
func DelNx(key string) error {
	err := WithClient(func(cli *redis1.Client) error {
		return cli.Del(key).Err()
	})
	if err != nil {
		return err
	}
	return nil
}

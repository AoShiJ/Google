package app

import (
	"demo/lx/framework/nacos"
	"github.com/spf13/viper"
)

func Init(s ...string) {
	nacos.ClientConfig(viper.GetString("Nacos.IpAddr"), viper.GetInt("Nacos.Port"))
}

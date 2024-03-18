package nacos

import (
	"github.com/astaxie/beego/logs"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
)

var Client config_client.IConfigClient

func ClientConfig(IpAddr string, Port int) error {
	clientConfig := constant.ClientConfig{
		NamespaceId:         "", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "log",
		CacheDir:            "cache",
		LogLevel:            "debug",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: IpAddr,
			Port:   uint64(Port),
		},
	}
	var err error
	Client, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
func GetConfig() (string, error) {
	dataid := viper.GetString("NaCos.DataID")
	group := viper.GetString("NaCos.Group")
	content, err := Client.GetConfig(vo.ConfigParam{
		DataId: dataid,
		Group:  group})
	if err != nil {
		return "", err
	}
	logs.Info(1)
	return content, nil
}

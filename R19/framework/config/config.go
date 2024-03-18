package config

import "github.com/spf13/viper"

type NaCosConf struct {
	DataID string `yaml:"DataID"`
	Group  string `yaml:"Group"`
	IpAddr string `yaml:"IpAddr"`
	Port   int64  `yaml:"Port"`
}

func InitViper(path string) error {
	viper.SetConfigFile(path)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

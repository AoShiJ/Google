package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql(fc func(db *gorm.DB) error) error {
	viper.SetConfigFile("./etc/order.yaml")
	viper.ReadInConfig()
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("MysqlConf.User"),
		viper.GetString("MysqlConf.Password"),
		viper.GetString("MysqlConf.Host"),
		viper.GetString("MysqlConf.Port"),
		viper.GetString("MysqlConf.DataBaseName"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Errorf(err.Error())
		return err
	}
	err = fc(db)
	if err != nil {
		return err
	}
	r, err := db.DB()
	if err != nil {
		return err
	}
	defer r.Close()
	return nil
}

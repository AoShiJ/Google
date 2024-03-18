package mysql

import (
	"demo/lx/framework/nacos"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func FuncMysql(SqlFc func(db *gorm.DB) error) error {
	var err error
	config, err := nacos.GetConfig()
	if err != nil {
		return err
	}
	var t T
	json.Unmarshal([]byte(config), &t)
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", t.Mysql.Username, t.Mysql.Password, t.Mysql.Host, t.Mysql.Port, t.Mysql.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	err = SqlFc(db)
	if err != nil {
		return err
	}
	d, _ := db.DB()
	d.Close()
	return nil
}

type T struct {
	Mysql struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
		Host     string `json:"Host"`
		Port     string `json:"Port"`
		Database string `json:"Database"`
	} `json:"Mysql"`
}

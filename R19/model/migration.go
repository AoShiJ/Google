package model

import (
	"demo/lx/framework/mysql"
	goods1 "demo/lx/model/goods"
	"demo/lx/model/order"
	"gorm.io/gorm"
)

func Migration() error {
	err := mysql.FuncMysql(func(db *gorm.DB) error {
		return db.AutoMigrate(new(User), new(goods1.Goods), new(order.Order))
	})
	if err != nil {
		return err
	}
	return nil
}

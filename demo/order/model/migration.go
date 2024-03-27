package model

import "gorm.io/gorm"

func Migration() {
	err := InitMysql(func(db *gorm.DB) error {
		return db.AutoMigrate(new(Order), new(OrderGoods))
	})
	if err != nil {
		panic(err)
	}
}

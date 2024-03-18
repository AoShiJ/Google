package order

import (
	"demo/lx/framework/mysql"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Title       string
	OrderNumber string
	TotalPrice  string
	ToTalNumber int
}

func (o Order) Create(order *Order) (Order *Order, err error) {

	err = mysql.FuncMysql(func(db *gorm.DB) error {
		if err = db.Create(&order).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	return
}

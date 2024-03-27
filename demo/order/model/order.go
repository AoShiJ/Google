package model

import (
	"fmt"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserId  int64
	OrderOn string
	Amount  string
	State   int
}
type OrderGoods struct {
	gorm.Model
	OrderID   int64 `gorm:"column:order_id"` // 对应数据库表中的字段名为 order_id
	GoodsID   int64 `gorm:"column:goods_id"` // 对应数据库表中的字段名为 goods_id
	UnitPrice string
	GoodName  string
	Num       int
}

func (o Order) Insert(order *Order) (*Order, error) {
	err := InitMysql(func(db *gorm.DB) error {
		return db.Create(&order).Error
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (o Order) UpdateByOrderNO(order *Order) (*Order, error) {
	sql := fmt.Sprintf("UPDATE orders SET state = ? WHERE order_on = ?")
	err := InitMysql(func(db *gorm.DB) error {
		return db.Exec(sql, order.State, order.OrderOn).Error
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (o Order) Update(order *Order) (*Order, error) {
	err := InitMysql(func(db *gorm.DB) error {
		return db.Where("id=?", order.ID).Updates(&order).Error
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o Order) GetOrderGoodsById(id int64) ([]*OrderGoods, error) {
	var order []*OrderGoods
	err := InitMysql(func(db *gorm.DB) error {
		return db.Where("id=?", id).Find(&order).Error
	})
	if err != nil {
		return nil, err
	}
	return order, nil
}
func (o OrderGoods) CreateOrderGoods(OrderGoods []*OrderGoods) ([]*OrderGoods, error) {
	err := InitMysql(func(db *gorm.DB) error {
		return db.Create(&OrderGoods).Error
	})
	if err != nil {
		return nil, err
	}
	return OrderGoods, nil
}
func NewOrder() *Order {
	return new(Order)
}

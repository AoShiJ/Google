package server

import (
	"context"
	"demo/lx/message/order"
	o "demo/lx/model/order"
	"gorm.io/gorm"
)

func CreateOrder(ctx context.Context, in *order.OrderInfo) {
	order := o.Order{}
	order.Create(msp(in))
}
func msp(in *order.OrderInfo) *o.Order {
	return &o.Order{
		Model: gorm.Model{
			ID: uint(in.Id),
		},
		Title:       in.Title,
		OrderNumber: in.OrderNumber,
		TotalPrice:  in.TotalPrice,
		ToTalNumber: int(in.TotalNumber),
	}
}

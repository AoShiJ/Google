package model

import (
	"fmt"
	"goods/rpc/goods"
	"gorm.io/gorm"
	"strings"
)

type Goods struct {
	Id    int64
	Name  string
	Price string
	Stock int64
}

func NewGoods() *Goods {
	return new(Goods)
}
func (g Goods) GetGoodsByIDs(id []int64) ([]*Goods, error) {
	var in []string
	var args []interface{}
	for _, i := range id {
		in = append(in, "?")
		args = append(args, i)
	}
	sprintf := fmt.Sprintf("id in (%v)", strings.Join(in, ","))
	var goods []*Goods
	err := InitMysql(func(db *gorm.DB) error {
		return db.Where(sprintf, args...).Find(&goods).Error
	})
	if err != nil {
		return nil, err
	}
	return goods, nil
}

func (g Goods) UpdateStock(req []*goods.UpdateStockReq) ([]*Goods, error) {
	var ids []int64
	for _, i := range req {
		ids = append(ids, i.ID)
		sql := fmt.Sprintf("UPDATE goods SET stock = stock - %d WHERE id = ?", i.Num)
		err := InitMysql(func(db *gorm.DB) error {
			return db.Exec(sql, i.ID).Error
		})
		if err != nil {
			return nil, err
		}
	}

	updatedGoods, err := g.GetGoodsByIDs(ids)
	if err != nil {
		return nil, err
	}

	return updatedGoods, nil
}

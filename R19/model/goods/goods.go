package goods

import (
	"demo/lx/framework/mysql"
	"github.com/astaxie/beego/logs"
	"gorm.io/gorm"
)

type Goods struct {
	gorm.Model
	Name  string
	Price string
	Num   int
}

func (g Goods) CreateGoods(goodsInfo *Goods) (*Goods, error) {
	err := mysql.FuncMysql(func(db *gorm.DB) error {
		logs.Info(goodsInfo)
		e := db.Create(goodsInfo).Error
		return e
	})
	if err != nil {
		return nil, err
	}
	return goodsInfo, nil
}
func (g Goods) DeleteGoods(id int) error {
	n := new(Goods)
	err := mysql.FuncMysql(func(db *gorm.DB) error {
		e := db.Where("id=?", id).Delete(n).Error
		return e
	})
	if err != nil {
		return err
	}
	return nil
}
func (g *Goods) UpdateGoods(goods *Goods) (*Goods, error) {
	err := mysql.FuncMysql(func(db *gorm.DB) error {
		e := db.Model(g).Where("id=?", goods.ID).Updates(goods).Error
		return e
	})
	if err != nil {
		return nil, err
	}
	return goods, nil
}
func (g Goods) SelectGoods(name string) (*Goods, error) {
	gg := new(Goods)
	err := mysql.FuncMysql(func(db *gorm.DB) error {
		r := db.Where("name=?", name).First(gg).Error
		return r
	})
	if err != nil {
		logs.Info(err, "====err")
		return nil, err
	}
	return gg, nil
}

func (g Goods) SelectsGoods(offset, limit int) ([]Goods, error) {
	var n []Goods
	logs.Info(offset, limit)
	err := mysql.FuncMysql(func(db *gorm.DB) error {
		r := db.Offset(offset).Limit(limit).Find(&n).Error
		return r
	})
	if err != nil {
		return nil, err
	}
	logs.Info(n, "======n")
	return n, nil
}

package model

import (
	"demo/lx/framework/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Mobile   string
}

func (u User) Select(username string) (*User, error) {
	uu := &User{}
	err := mysql.FuncMysql(func(db *gorm.DB) error {
		db.Where("username=?", username).First(&uu)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return uu, nil
}

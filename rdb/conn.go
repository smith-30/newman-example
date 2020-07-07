package rdb

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Mysql(username, password, host, port, dbName string, logging bool) (*gorm.DB, error) {
	address := host + ":" + port
	setting := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, address, dbName)
	fmt.Printf("setting %#v\n", setting)
	db, err := gorm.Open("mysql", setting)
	db.LogMode(logging)

	return db, err
}

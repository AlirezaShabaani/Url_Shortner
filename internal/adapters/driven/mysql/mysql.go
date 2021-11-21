package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
//===========
//type mysqlutils struct {
//	DB *gorm.DB
//}
//
//func New(DB *gorm.DB) *mysqlutils {
//	return &mysqlutils{DB: DB}
//}
//============

func ConnectMysql(host, port, username, password, dbName string) (DB *gorm.DB, err error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbName)
	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		return
	}
	return
}

package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"url_shortner/internal/adapters/driven/mysql/model"
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

func InitMysql(host, port, username, password, dbName string) (DB *gorm.DB) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbName)
	DB, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal("can't connect to mysql db ,  err : " + err.Error())
	}
	err = DB.AutoMigrate(model.Data{})
	if err != nil {
		log.Fatal("can't migrate ")
	}
	return
}

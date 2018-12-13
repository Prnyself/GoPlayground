package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type CampusApp struct {
	Id        int
	AppKey    string
	AppSecret string
}

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:1101@tcp(127.0.0.1:3306)/new_platform?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

func GetAppById(id int) (app CampusApp) {
	db.First(&app, id) // 查询id为1的product
	return
}

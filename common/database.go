package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host 	 := "localhost"
	port 	 := "3306"
	database := "west"
	username := "root"
	password := "qwe123"
	charset  := "utf8"
	args 	 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",username,password,host,port,database,charset)
	fmt.Println(args)
	db, err  := gorm.Open(mysql.Open(args),&gorm.Config{})
	if err != nil{
		panic("fail to connect database, err: "+ err.Error())
	}
	DB = db
	return db
}

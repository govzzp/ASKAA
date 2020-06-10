package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type AnswerSheet struct {
	gorm.Model
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Depart   string `json:"depart"`
	Name     string `json:"name"`
}
var db *gorm.DB
func init() {
	if err := initDB() ; err != nil {
		panic(err)
	}
}

func initDB()(err error){
	db,err =gorm.Open("mysql","root:admin123456@/go_ask?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("connect error: %v\n",err)
	}
	db.AutoMigrate(&AnswerSheet{})
	return
}
// a.com
// 443 80
// 反向代理
//server_name a.com
//location / {
//	proxy_pass http://127.0.0.1:9090
//}
// pm2
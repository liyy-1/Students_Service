package config

import (
	"fmt"
	"log"
	"my_project/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	username := "go_user"
	password := "Liyi20050830@"
	host := "127.0.0.1"
	port := "3306"
	dbname := "student_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败: ", err)
	}

	log.Println("数据库连接成功")

	// 自动建表
	err = DB.AutoMigrate(&model.Student{})
	if err != nil {
		log.Fatal("自动建表失败: ", err)
	}
}

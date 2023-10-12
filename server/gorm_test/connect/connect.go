package main

import (
	json2 "encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Price uint
	Code  string
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root123@tcp(1.94.22.229:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // enable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("数据库链接错误")
		panic(err)
	}
	// 创建表
	//db.AutoMigrate(&Product{})
	// 通过数据的指针来创建数据
	//pro := Product{
	//	Code:  "Hello",
	//	Price: 100,
	//}
	//db.Create(&pro)
	// 查询
	var product Product
	result := db.Where(&Product{Price: 10}).Find(&product)
	if result.Error != nil {
		fmt.Println("查询错误")
		return
	}
	fmt.Println(product)
	json, err := json2.Marshal(product)
	if err != nil {
		fmt.Println("json解析错误")
		return
	}
	fmt.Println(string(json))
}

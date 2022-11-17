package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var GlobalDb *gorm.DB

// 免责声明:演示用例子重庆邮电大学与重庆大学纯作虚构,现实中双非可不敢碰瓷985捏
// 免责声明:演示用例子重庆邮电大学与重庆大学纯作虚构,现实中双非可不敢碰瓷985捏
// 免责声明:演示用例子重庆邮电大学与重庆大学纯作虚构,现实中双非可不敢碰瓷985捏
// 免责声明:演示用例子重庆邮电大学与重庆大学纯作虚构,现实中双非可不敢碰瓷985捏
// 免责声明:演示用例子重庆邮电大学与重庆大学纯作虚构,现实中双非可不敢碰瓷985捏

func main() {
	db, _ := gorm.Open(mysql.New(mysql.Config{ //配置
		DSN: "root:123456@tcp(127.0.0.1:3306)/CQUPT?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gocy_",
			SingularTable: false,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(10) //数据池
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	GlobalDb = db   //全局变量GlobalDb赋值
	TestUserCreat() //若无表单自动创建表单
	//Create() //增
	//Delete() //删
	//Update() //改
	//Find() //查

}

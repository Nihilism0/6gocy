package main

import (
	"fmt"
	"time"
)

func Create() {
	dbres := GlobalDb.Model(&User{}).Create(&[]User{ //创建表单,
		{Title: "重庆邮电大学是重庆第一大学", SubmissionTime: time.Now()},
		{Title: "重庆大学是重庆第二大学", SubmissionTime: time.Now()},
		{Title: "重庆邮电大学是重庆第一大学", SubmissionTime: time.Now()},
		{Title: "重庆大学是重庆第二大学", SubmissionTime: time.Now()},
		{Title: "重庆大学是重庆第二大学", SubmissionTime: time.Now()},
	})

	if dbres.Error != nil { //错误处理
		fmt.Println("你创建了n(s)i(h)l(i)!(t)")
	} else {
		fmt.Println("创建成功了 嘻嘻~")
	}
}

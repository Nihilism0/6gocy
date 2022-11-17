package main

import (
	"fmt"
	"time"
)

type User struct {
	ID             int
	Title          string
	SubmissionTime time.Time
}

func TestUserCreat() {
	err := GlobalDb.AutoMigrate(&User{})
	if err != nil {
		fmt.Println("创建表单出错")
		return
	}
}

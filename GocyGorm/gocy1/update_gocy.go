package main

import "fmt"

func Update() {
	GlobalDb.Model(&User{}).Where("title LIKE ?", "%重庆大学%").Update("title", "重庆邮电大学是重庆第一大学")
	fmt.Println("成功更新")
}

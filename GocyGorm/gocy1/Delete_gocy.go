package main

import "fmt"

func Delete() {
	var user []User
	GlobalDb.Model(&User{}).Unscoped().Where("title LIKE ?", "%重庆大学%").Delete(&user)
	fmt.Println("删除成功!")
}

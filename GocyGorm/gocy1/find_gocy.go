package main

import "fmt"

type Info struct {
	ID    int
	Title string
}

func Find() {
	var u []Info
	GlobalDb.Model(&User{}).Where("title LIKE ?", "%第一%").Find(&u)
	fmt.Println(u)
}

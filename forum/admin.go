package main

import "fmt"

var adminList []admin

type admin struct {
	name string
	age  int
	post string
}

func main() {

	addAdmin("Victor", 15, "Administrator of shop")
}
func addAdmin(name, age, post) {
	a := admin{name, age, post}
	adminList = append(adminList, a)
}
func petTheCat() {
	fmt.Println("Кот доволен!")
}

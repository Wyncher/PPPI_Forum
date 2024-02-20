package main

import "fmt"

var moderList = []

type moder struct {
	name string
	age  int
	post string
}


func main() {

	add("Pavlik",22,"Moderator of tree 'Флудилка'")
	add("Artemon887",17,"Moderator of tree 'Компьютерные игры'")
}

func add(name, age, post) {
	a := moder{name, age, post}
	adminList = append(adminList,a)
}

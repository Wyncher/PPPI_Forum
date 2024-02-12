package main

import "fmt"

type moder struct {
	name string
	age  int
	post string
}

func main() {
	x := moder{
		name: "Victor",
		age:  15,
		post: "Moderator of shop",
	}
	fmt.Println(x)
}

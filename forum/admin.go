package main

import "fmt"

type admin struct {
	name string
	age  int
	post string
}

func main() {
	x := admin{
		name: "Victor",
		age:  15,
		post: "Administrator of shop",
	}
	fmt.Println(x)
}

func add(name, age, post) {

}

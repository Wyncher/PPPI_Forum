package main



var adminList = []
type admin struct {
	name string
	age  int
	post string
}

func main() {

	add("Victor",15,"Administrator of shop")
}

func add(name, age, post) {
	a := admin{name, age, post}
	adminList = append(adminList,a)

}

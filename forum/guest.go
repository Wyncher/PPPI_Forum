package main

import (
	"fmt"
	"log"
)

func outMessage(){
	fmt.Println("До скорой встречи! Ждем Вас снова!")
	log.Println("Пользователь покинул сайт")
}

func intoMessage(){
	fmt.Println("Спасибо за то,что зашли к нам. Посещайте нас почаще\nА лучше заругистрируйтесь")
	log.Println("Пользователь вошел на сайт")
}

func main(){
	var choice
	fmt.Scanf("%s",&choice)
	if choice == "into"{
		intoMessage()
	}
	if choice == "out"{
		outMessage()
	}

}
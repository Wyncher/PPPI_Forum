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
func buyFunc() {
	i := "0"
	fmt.Scanf("%s",&i)
	switch i {
	case "Game Account":
		fmt.Println("Покупка игрового аккаунта")
	case "Game Currency":
		fmt.Println("Покупка игровой валюты")
	case "Facebook Account":
		fmt.Println("Покупка аккаунта Facebook")
	}
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
	if choice == "cart"{
		buyFunc()
	}

}


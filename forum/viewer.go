package main

import (
	"fmt"
	"log"
	"time"
)

var balR = 0

func outMessage(){
	fmt.Println("До скорой встречи! Ждем Вас снова!")
	log.Println("Пользователь покинул сайт")
}

func intoMessage(){
	fmt.Println("Спасибо за то,что зашли к нам. Посещайте нас почаще\nА лучше заругистрируйтесь")
	log.Println("Пользователь вошел на сайт")
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	log.Println("Начислено 2 рубля за посещение сайта")
	fmt.Println("Вам начислено 2 рубля за посещение сайта\nЧтобы получить их нужно зарегистрироваться")
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
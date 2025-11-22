package main

import (
	"fmt"
	"time"
)

func saudacao(c chan string) {
	c <- "Desafio Go - Goroutines e Channels! \nPressione Enter para sair..."
}

func ping(c chan string) {
	for {
		c <- "ping"
		time.Sleep(time.Second)
	}
}

func pong(c chan string) {
	for {
		c <- "pong"
		time.Sleep(time.Second)
	}
}

func print(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func main() {
	c := make(chan string)

	go saudacao(c)

	go ping(c)
	go pong(c)
	go print(c)

	var entrada string
	fmt.Scanln(&entrada)
}

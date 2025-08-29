package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"

	}
}

func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"

	}
}
func pingPrint(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}
func pongPrint(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	var c1 chan string = make(chan string)
	var c2 chan string = make(chan string)
	go ping(c1)
	go pong(c2)
	go pingPrint(c1)
	go pongPrint(c2)

	var entrada string
	fmt.Scanln(&entrada)
}

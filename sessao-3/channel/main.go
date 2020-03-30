package main

import (
	"fmt"
	"time"
)

func pong(canal chan string) {
	for {
		canal <- "ping"
	}
}

func ping(canal chan string) {
	for {
		canal <- "pong"
	}
}

func printer(canal chan string) {
	for {
		msg := <-canal
		println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var canal chan string
	canal = make(chan string)

	go ping(canal)
	go pong(canal)
	go printer(canal)

	var input string
	fmt.Scanln(&input)
}

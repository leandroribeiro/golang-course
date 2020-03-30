package main

import (
	"fmt"
	"time"
)

var irc = make(chan string)
var sms = make(chan string)

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

func printer() {

	var msg string
	for {
		select {
		case msg = <-irc:
			println(msg)
			time.Sleep(time.Second * 1)
		case msg = <-sms:
			println("sms >>", msg)
		}
		time.Sleep(time.Second * 1)
	}

}

func main() {
	go ping(irc)
	go pong(irc)

	go ping(sms)
	go pong(sms)

	go printer()

	var input string
	fmt.Scanln(&input)
}

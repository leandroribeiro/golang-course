package main

import (
	"fmt"
	myhandlers "github.com/leandroribeiro/go-hello-world/sessao-4/webserver/handlers"
	"net/http"
)


func main(){
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Olá Mundo!")
	})

	http.HandleFunc("/function", myhandlers.MyCustomHandler)

	http.HandleFunc("/hello", myhandlers.Hello)

	fmt.Println("I'm ready!")
	http.ListenAndServe(":8181", nil)
}

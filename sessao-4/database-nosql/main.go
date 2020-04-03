package main

import (
	"fmt"
	"github.com/leandroribeiro/go-hello-world/sessao-4/database-nosql/handler"
	"github.com/leandroribeiro/go-hello-world/sessao-4/database-nosql/repository"
	"net/http"
)

func init()  {
	fmt.Println("Começando a brincadeira...")
}

func main()  {

	err2 := repository.Open()
	if err2 != nil {
		fmt.Println("O banco de dados explodiu! ", err2.Error())
	}

	http.HandleFunc("/city/", handler.City)

	fmt.Println("Estou pronto vamos?")

	http.ListenAndServe(":8181", nil)
}

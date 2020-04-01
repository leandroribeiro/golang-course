package main

import (
	"fmt"
	"github.com/leandroribeiro/go-hello-world/sessao-4/database/handler"
	"github.com/leandroribeiro/go-hello-world/sessao-4/database/repository"
	"net/http"
)

func init()  {
	fmt.Println("Começando a brincadeira...")
}

func main()  {
	err := repository.Open()
	if err != nil {
		fmt.Println("O banco de dados explodiu! ", err.Error())
	}

	http.HandleFunc("/city/", handler.City)

	fmt.Println("Estou pronto vamos?")

	http.ListenAndServe(":8181", nil)
}

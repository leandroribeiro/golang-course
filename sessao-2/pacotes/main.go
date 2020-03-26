package main

import (
	"fmt"
	"github.com/leandroribeiro/go-hello-world/sessao-2/pacotes/operadora"
	"github.com/leandroribeiro/go-hello-world/sessao-2/pacotes/prefixo"
)

//Username é o nome do usuário logado
var Username = "Palpatine"

func main() {
	fmt.Printf("Username: %s\r\n", Username)
	fmt.Printf("Capital number: %d\r\n", prefixo.Capital)
	fmt.Printf("Provider name: %s\r\n", operadora.NomeOperadora)
}

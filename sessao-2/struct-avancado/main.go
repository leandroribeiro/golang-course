package main

import (
	"encoding/json"
	"fmt"
	"github.com/leandroribeiro/go-hello-world/sessao-2/struct-avancado/model"
)

func main()  {
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(1000)

	fmt.Printf("Casa é: %+v\r\n", casa)
	fmt.Printf("O valor da casa é: %+v\r\n", casa.GetValor())

	objJson, _ := json.Marshal(casa)

	fmt.Printf("A casa em formato JSON: %v", string(objJson))
}


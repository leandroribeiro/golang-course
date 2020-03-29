package main

import (
	"encoding/json"
	"fmt"
	"github.com/leandroribeiro/go-hello-world/sessao-2/struct-avancado/model"
)

func main()  {
	casa := model.Imovel{}
	casa.Nome = "Casa de Papel"
	casa.X = 20
	casa.Y = 30
	casa.SetValor(1)

	fmt.Printf("Casa é ... %+v\r\n", casa)

	objJSON, _ := json.Marshal(casa)
	fmt.Println("Casa em JSON ... ", string(objJSON))

//GOOS=windows
//GOARCH=386

//GOOS=linux
//GOARCH=arm
}

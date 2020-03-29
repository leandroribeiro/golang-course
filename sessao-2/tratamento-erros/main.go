package main

import (
	"encoding/json"
	"fmt"
	"github.com/leandroribeiro/go-hello-world/sessao-2/struct-avancado/model"
)

func main()  {
	casa  := model.Imovel{}
	casa.Nome = "Casa do Zezinho"
	casa.X = 17
	casa.Y = 29
	if err := casa.SetValor(50000); err !=nil {
		fmt.Println("[main] Deu pa* meu fio!", err.Error())
		if err == model.ErrValorDeImovelMuitoAlto {
			fmt.Println("Está caro demais!")
		}
		return
	}

	fmt.Printf("Casa .... %+v\r\n", casa)

	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Deu pa* meu fio!", err.Error())
		return
	}
	fmt.Println("Casa JSON", string(objJSON))
}

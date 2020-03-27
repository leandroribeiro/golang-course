package main

import (
	"fmt"
	"github.com/leandroribeiro/go-hello-world/sessao-2/struct-avancado/model"
)

var cache map[string]model.Imovel

func main()  {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(60000)

	cache["Casa Amarela"] =  casa

	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, e o que achei foi... %+v\r\n", imovel)
	}

	// add more one

	fmt.Println("Quantos itens?", len(cache))

	for key, value := range cache {
		fmt.Printf("Chave[%s] %+v\n\r", key, value)
	}

	imovel, achou =  cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}

	fmt.Println("Quantos itens?", len(cache))
}

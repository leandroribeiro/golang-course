package main

import "fmt"

//Imovel é alguma coisa
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("Casa... %+v\r\n", casa)

	apartamento := Imovel{17, 56, "Meu AP", 76000}
	fmt.Printf("Casa... %+v\r\n", apartamento)

	chacara := Imovel{
		Y:     85,
		Nome:  "Chacara",
		valor: 55,
		X:     22,
	}

	fmt.Printf("Casa... %+v\r\n", chacara)

	casa.Nome = "Minha casa minha vida"
	casa.X = 1
	casa.Y = 2
	casa.valor = 5000
	fmt.Printf("Casa... %+v\r\n", casa)
}

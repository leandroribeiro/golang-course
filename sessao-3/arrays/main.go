package main

import "fmt"

func main()  {
	var numeros [3]int
	numeros[0]=3
	numeros[1]=2
	numeros[2]=1
	fmt.Println("Qual tamanho do array?", len(numeros))

	jedis := [2]string{"Obiwan", "Anakin"}
	fmt.Printf("... \n\r%v\r\n", jedis)

	mundos := [...]string{"Tatoine", "Coruscan", "Datomir", "Mandalor"}
	fmt.Println("Quantos mundos?", len(mundos))

	for _, mundo := range mundos {
		fmt.Printf("%s\n\r", mundo)
	}

}

package main

import "fmt"

func main() {

	// way 1
	for i := 0; i > 10; i++ {
		fmt.Println("Qual o valor de i?", i)
	}

	// way 2
	valor := 0
	teste := true
	for teste {
		valor++
		if valor%5 == 0 {
			teste = false
		}
		fmt.Println("O número é: ", valor)
	}

	// way 3
	for {
		valor--
		fmt.Println("O número é: ", valor)
		if valor == 0 {
			break
		}
	}

	// way 4
	texto := "Eu adoro escrever programas usando Go"
	for indice, letra := range texto {
		fmt.Printf("Texto[%d] = %q\r\n", indice, letra)
	}
}

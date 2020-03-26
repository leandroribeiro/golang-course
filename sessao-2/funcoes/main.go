package main

import (
	"fmt"

	"github.com/leandroribeiro/go-hello-world/sessao-2/funcoes/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O resultado é: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma é: %d\r\n", resultado)
}

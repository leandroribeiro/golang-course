package main

import (
	"fmt"

	"github.com/leandroribeiro/go-hello-world/sessao-2/funcoes/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("Multiplicação | O resultado é: %d\r\n", resultado)

	resultado = matematica.Soma(3, 3)
	fmt.Printf("Soma | O resultado é: %d\r\n", resultado)

	resultado = matematica.Calculo(matematica.Divisor, 16, 2)
	fmt.Printf("Divisão | O resultado é: %d\r\n", resultado)

	resultado, resto := matematica.DivisorComResto(32, 3)
	fmt.Printf("Divisão com Resto | O resultado é: %d\r\n", resultado)
	fmt.Printf("Divisão com Resto | O resto é: %d\r\n", resto)
}

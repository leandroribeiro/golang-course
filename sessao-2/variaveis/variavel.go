package main

import "fmt"

var (
	endereco                         string //endereco = ""
	telefone                         = "9876-4321"
	umacoisa, outracoisa             string
	quantidade                       int     //quantidade = 0
	comprou                          bool    //comprou =  false
	valor                            float64 //valor = 0.00
	palavras                         rune
	UmaCoisaPublica, umaCoisaPrivada string //iniciando em maiscula vira Público
	CoisaPublica                     string
)

func main() {
	coisa := "qualquer coisa"

	fmt.Printf("endereço: %s\r\n", endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)
	fmt.Printf("coisa: %v\r\n", coisa)
}

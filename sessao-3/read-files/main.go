package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	arquivo, err := os.Open("CidadesEstados.csv")
	if err != nil {
	fmt.Println("Deu pau!")
		return
	}

	//scanner := bufio.NewScanner(arquivo)
	//for scanner.Scan(){
	//	fmt.Println(scanner.Text())
	//}

	leitor := csv.NewReader(arquivo)
	leitor.Comma = ';' //custom separator
	conteudo, err := leitor.ReadAll()
	if err != nil {
		fmt.Println("Deu pau!")
		return
	}

	for _, linha := range conteudo{
		//fmt.Println(linha)
		for _, coluna := range linha {
			fmt.Println(coluna)
		}
	}

	arquivo.Close()
}

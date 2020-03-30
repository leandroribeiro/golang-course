package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/leandroribeiro/golang-course/sessao-3/write-files/model"
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

	baseJson, err := os.Create("base.json")
	if err != nil {
		fmt.Println("Deu pau!")
		return
	}

	escritor := bufio.NewWriter(baseJson)
	escritor.WriteString("[\r\n")

	for indice, linha := range conteudo{
		//fmt.Println(linha)
		cidade := model.Cidade{}
		cidade.Estado = linha[0]
		cidade.Nome = linha[1]

		//for _, coluna := range linha {
			//fmt.Println(coluna)
			//cidade := model.Cidade{}
		//}
		jsonData, err := json.Marshal(cidade)
		if err != nil {
			fmt.Println("Deu pau!")
			return
		}
		escritor.Write(jsonData)

		if indice < len(conteudo) {
			escritor.WriteString(",")
		}

	}

	escritor.WriteString("]")
	escritor.Flush()

	baseJson.Close()
	arquivo.Close()
}

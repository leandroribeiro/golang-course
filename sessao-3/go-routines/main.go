package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/leandroribeiro/golang-course/sessao-3/write-files/model"
	"os"
	"sync"
	"time"
)

var mestre sync.WaitGroup

func main() {
	mestre.Add(2)

	go convertToJSON("cidadesRJ")
	go convertToJSON("cidadesSP")

	mestre.Wait()
}

func convertToJSON(fileName string) {

	csvFileName := fileName + ".csv"
	jsonFileName := fileName + ".json"

	arquivo, err := os.Open(csvFileName)
	if err != nil {
		fmt.Println("Deu pau!")
		return
	}

	defer arquivo.Close()

	leitor := csv.NewReader(arquivo)
	leitor.Comma = ';' //custom separator
	conteudo, err := leitor.ReadAll()
	if err != nil {
		fmt.Println("Deu pau!")
		return
	}

	baseJson, err := os.Create(jsonFileName)
	if err != nil {
		fmt.Println("Deu pau!")
		return
	}

	defer baseJson.Close()

	escritor := bufio.NewWriter(baseJson)
	escritor.WriteString("[\r\n")

	for indice, linha := range conteudo {
		cidade := model.Cidade{}
		cidade.Estado = linha[0]
		cidade.Nome = linha[1]

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

	fmt.Println(time.Now(), " done!")

	mestre.Done()
}

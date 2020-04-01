package main

import (
	"encoding/json"
	"fmt"
	"github.com/leandroribeiro/go-hello-world/sessao-4/response-webservice/model"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {


	println("")
	println("")
	println("Do Request With Auth")
	println("")
	println("")

	client := &http.Client{
		Timeout: time.Second * 30,
	}

	// Request with Authentication Head
	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		println("Deu pau!", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	response, err := client.Do(request)

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			println("Deu pau!", err.Error())
			return
		}
		blogPost := model.BlogPost{}
		_ = json.Unmarshal(body, &blogPost)

		fmt.Printf("%+v", blogPost)
	}
}

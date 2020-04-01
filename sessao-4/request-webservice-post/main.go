package main

import (
	"bytes"
	"encoding/json"
	"github.com/leandroribeiro/go-hello-world/sessao-4/request-webservice-post/model"
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

	user := model.User{}
	user.ID = 1
	user.Nome = "Dart Maul"

	userJson, _ := json.Marshal(user)

	// Request with Authentication Head
	request, err := http.NewRequest("POST", "https://enjxrdc25nr3f.x.pipedream.net", bytes.NewBuffer(userJson))
	if err != nil {
		println("Deu pau!", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	request.Header.Set("content-type", "application/json;charset=utf-8")
	response, err := client.Do(request)

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			println("Deu pau!", err.Error())
			return
		}
		println(string(body))
	}
}

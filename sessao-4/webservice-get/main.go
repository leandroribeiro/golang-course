package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	println("")
	println("")
	println("Do Request WithOut Auth")
	println("")
	println("")

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	response, err := client.Get("https://www.google.com")
	if err != nil {
		println("Deu pau!", err.Error())
		return
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			println("Deu pau!", err.Error())
			return
		}
		println(string(body))
	}

	println("")
	println("")
	println("Do Request With Auth")
	println("")
	println("")

	// Request with Authentication Head
	request, err := http.NewRequest("GET", "https://www.google.com.br", nil)
	if err != nil {
		println("Deu pau!", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	response, err = client.Do(request)

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			println("Deu pau!", err.Error())
			return
		}
		println(string(body))
	}
}

package handler

import (
	"fmt"
	model3 "github.com/leandroribeiro/go-hello-world/sessao-4/database-nosql/model"
	repository2 "github.com/leandroribeiro/go-hello-world/sessao-4/database-nosql/repository"
	"net/http"
	"strconv"
)

func City(write http.ResponseWriter, r *http.Request)  {
	model := model3.City{}
	phoneCode, err := strconv.Atoi(r.URL.Path[6:])

	if err != nil {
		http.Error(write, "Fez merda!", http.StatusBadRequest)
		fmt.Println(err.Error())
	}

	model, err = repository2.Get(phoneCode)

	fmt.Println(phoneCode)

	if err != nil {
		http.Error(write, "Não existe mula!", http.StatusNotFound)
		fmt.Println(err.Error())
	}

	if model.PhoneCode == 0 {
		http.Error(write, "Não existe mula!", http.StatusNotFound)
		return
	}

	if err := LocalTemplate.ExecuteTemplate(write, "city.html", model); err != nil{
		http.Error(write, "Deu.....ruim.. :(", http.StatusInternalServerError)
		fmt.Println(err.Error())
	}

}
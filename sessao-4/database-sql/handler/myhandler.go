package handler

import (
	"fmt"
	model2 "github.com/leandroribeiro/go-hello-world/sessao-4/database-sql/model"
	"github.com/leandroribeiro/go-hello-world/sessao-4/database-sql/repository"
	"net/http"
	"strconv"
	"time"
)

func City(write http.ResponseWriter, r *http.Request)  {
	model := model2.City{}
	phoneCode, err := strconv.Atoi(r.URL.Path[6:])

	if err != nil {
		http.Error(write, "Fez merda!", http.StatusBadRequest)
		fmt.Println(err.Error())
	}

	sql := "select * from cities where phone_code = ?"
	rows, err := repository.Db.Queryx(sql, phoneCode)

	fmt.Println(phoneCode)

	if err != nil {
		http.Error(write, "Não existe mula!", http.StatusNotFound)
		fmt.Println(err.Error())
	}

	for rows.Next() {
		err = rows.StructScan(&model)

		if err != nil {
			http.Error(write, "Deu merda", http.StatusInternalServerError)
			fmt.Println(err.Error())
		}
	}

	if model.PhoneCode == 0 {
		http.Error(write, "Não existe mula!", http.StatusNotFound)
		return
	}

	if err := LocalTemplate.ExecuteTemplate(write, "city.html", model); err != nil{
		http.Error(write, "Deu.....ruim.. :(", http.StatusInternalServerError)
		fmt.Println(err.Error())
	}

	sql = "insert into logquery (daterequest) values (?)"

	result, err := repository.Db.Exec(sql, time.Now().Format("01/04/2020 23:59:59"))

	if err != nil {
		http.Error(write, "Deu merda", http.StatusInternalServerError)
		fmt.Println(err.Error())
	}

	rowsAffecteds, _ := result.RowsAffected()

	fmt.Println(rowsAffecteds, " linhas afetadas.")
}
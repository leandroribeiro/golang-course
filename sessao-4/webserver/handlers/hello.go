package myhandlers

import (
	"fmt"
	"github.com/leandroribeiro/go-hello-world/sessao-4/webserver/models"
	"net/http"
	"time"
)

func Hello(write http.ResponseWriter, r *http.Request)  {
	model := myModels.MyPage{}
	model.Now = time.Now().Format("14:31:01")
	if err := models.ExecuteTemplate(write, "hello.html", model); err != nil{
		http.Error(write, "Deu.....ruim.. :(", http.StatusInternalServerError)
		fmt.Println(err.Error())
	}
}
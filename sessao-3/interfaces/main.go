package main

import (
	"fmt"
	"github.com/leandroribeiro/golang-course/sessao-3/interfaces/model"
)

func main()  {
	obiwan := model.SerDeForca{}
	obiwan.Nome = "Obiwan Kenobi"

	Meditar(obiwan)
	SoltarRaio(obiwan)
}

func Meditar(o model.Jedi){
	fmt.Println(o.Medita())
}

func SoltarRaio(o model.Sith)  {
	fmt.Println(o.SoltarRaio())
}

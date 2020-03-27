package main

import "fmt"

func main()  {

	meses := 6
	situacao := true
	cidade := "Teste"

	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo")
	}

	if situacao {
		fmt.Println("Tudo okay!")
	}

	if cidade == "Teste" {
		fmt.Println("Isso é apenas um teste!")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println(descricao)
		return
	}

	fmt.Println("The End")
}

func haQuantoTempoEstaDevendo(meses int)(descricao string, status bool){
	if meses > 0 {
		status = true
		descricao =  "O cliente está INAdimplente"
		return
	}
	descricao = "O cliente está Adimplente"
	return
}

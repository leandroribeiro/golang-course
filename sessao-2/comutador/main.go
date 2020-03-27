package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numero := 3
	fmt.Print("O número ", numero, " se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("Um")
	case 2:
		fmt.Println("Dois")
	case 3:
		fmt.Println("Três")
	}

	fmt.Println("Você é da família do LINUX?")

	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebas":
		fallthrough
	case "linux":
		fmt.Println("SIM =)")
	default:
		fmt.Println("Provavelmente sou do RunWindows")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("FDSSSSSSS!!")
	case time.Friday:
		fmt.Println("Sextouuuu")
	default:
		fmt.Println("Vamos trabalhar?!!")
	}

	numero = 10
	fmt.Println("Este número tem 1 digito?")
	switch {
	case numero < 10:
		fmt.Println("SIM")
	case numero >= 10 && numero < 100:
		fmt.Println("NÃO")
	default:
		fmt.Println("O número é ALTO!")
	}

}

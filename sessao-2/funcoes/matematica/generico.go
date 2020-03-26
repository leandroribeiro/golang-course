package matematica

/*
Calculo executa qualquer tipo de calculo, basta enviar a função desejada
*/
func Calculo(funcao func(int, int) int, numeroX int, numeroY int) int {
	return funcao(numeroX, numeroY)
}

/*
Multiplicador advinha o que ele faz?
*/
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor advinha?
func Divisor(numeroX int, numeroY int) (resultado int) {
	resultado = numeroX / numeroY
	return
}

//DivisorComResto faz alguma coisa
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = numeroA / numeroB
	resto = numeroA % numeroB
	return
}

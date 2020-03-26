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

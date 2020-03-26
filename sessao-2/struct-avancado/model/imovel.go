package model

type Imovel struct {
	X int `json:"positionX"`
	Y int `json:"positionY"`
	Nome string `json:"name"`
	valor int
}

//SetValor defino o valor do imóvel
func (i *Imovel)  SetValor (valor int) {
	i.valor = valor
}

//GetValor retorna o valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}
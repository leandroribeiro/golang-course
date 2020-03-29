package model

import "errors"

type Imovel struct {
	X     int    `json:"positionX"`
	Y     int    `json:"positionY"`
	Nome  string `json:"name"`
	valor int
}


var ErrValorDeImovelInvalido = errors.New("O valor informado não é válido para um imovel!")

var ErrValorDeImovelMuitoAlto = errors.New("O valor informado é muito alto!")

//SetValor defino o valor do imóvel
func (i *Imovel) SetValor(valor int) (err error) {
	err = nil
	if valor < 1000 {
		err = ErrValorDeImovelInvalido
		return
	} else if valor > 10000{
		err = ErrValorDeImovelMuitoAlto
		return
	}
	i.valor = valor
	return
}

//GetValor retorna o valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}

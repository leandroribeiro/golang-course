package model

type Jedi interface {
	Medita() string
}

type Sith interface {
	SoltarRaio() string
}

type SerDeForca struct {
	Nome string
}

func (a SerDeForca) Medita() string  {
	return "Hummm...."
}

func (a SerDeForca) SoltarRaio() string{
	return "zzzzzhhhhhzz!"
}

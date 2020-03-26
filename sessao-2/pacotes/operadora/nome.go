package operadora

import (
	"github.com/leandroribeiro/go-hello-world/sessao-2/pacotes/prefixo"
	"strconv"
)

//NomeOperadora é o nome da operadora
var NomeOperadora = "Dark Telecom"

//PrefixoOperadora é o Nome + Prefixo concatenado
var PrefixoOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora

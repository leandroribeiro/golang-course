package template

import (
	"html/template"

	"github.com/leandroribeiro/go-hello-world/sessao-4/mercurius-framework/lib/contx"
)

// FuncMaps to view
func FuncMaps() []template.FuncMap {
	return []template.FuncMap{
		map[string]interface{}{
			"Tr": contx.I18n,
		}}
}

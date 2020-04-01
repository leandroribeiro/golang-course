package myhandlers

import "html/template"

var models = template.Must(template.ParseFiles("views/hello.html"))
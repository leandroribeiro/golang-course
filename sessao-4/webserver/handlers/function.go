package myhandlers

import (
	"fmt"
	"net/http"
)

func MyCustomHandler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintln(writer, "Im a function!")
}
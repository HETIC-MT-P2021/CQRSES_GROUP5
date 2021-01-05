package controllers

import (
	"fmt"
	"net/http"
)

//RenderHome renders home page
func RenderHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World !")
	return
}

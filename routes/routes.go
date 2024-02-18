package routes

import (
	"net/http"

	"github.com/tiagoc0sta/alura-web-application/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
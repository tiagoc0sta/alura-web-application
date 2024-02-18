package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/tiagoc0sta/alura-web-application/routes"
)


func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}


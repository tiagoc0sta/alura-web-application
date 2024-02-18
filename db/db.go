package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB { // Função retorna um ponteiro pro SQL
	conexao := "user=postgres dbname=alura_loja password=$Relogio10 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error)
	}
	return db
}
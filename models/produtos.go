package models

import (
	"github.com/tiagoc0sta/alura-web-application/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil { // se o erro não for nulo, ou seja, se existir o erro.
		panic(err.Error())
	}

	p := Produto{}
	produtos:= []Produto{}

	for selectDeTodosOsProdutos.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade


		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos

}
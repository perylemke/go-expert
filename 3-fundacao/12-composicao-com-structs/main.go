package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	pery := Cliente{
		Nome:  "Pery",
		Idade: 32,
		Ativo: true,
	}

	// Muda o valor de ativo na struct
	pery.Ativo = false

	// Atribui valores para struct composta
	pery.Endereco.Logradouro = "Rua dos Alfineiros"
	pery.Endereco.Numero = 42
	pery.Endereco.Cidade = "Florianópolis"
	pery.Endereco.Estado = "Santa Catarina"

	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\nLogradouro: %s\nNúmero: %d\nCidade: %s\nEstado: %s",
		pery.Nome,
		pery.Idade,
		pery.Ativo,
		pery.Endereco.Logradouro,
		pery.Endereco.Numero,
		pery.Endereco.Cidade,
		pery.Endereco.Estado,
	)
}

package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	pery := Cliente{
		Nome:  "Pery",
		Idade: 32,
		Ativo: true,
	}
	// Muda o valor de ativo na struct
	pery.Ativo = false

	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t", pery.Nome, pery.Idade, pery.Ativo)
}

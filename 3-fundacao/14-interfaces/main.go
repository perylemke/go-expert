package main

import "fmt"

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

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

// Interface só permite métodos, não permite propriedades.
type Pessoa interface {
	Desativar()
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado!\n", c.Nome)
}

func main() {
	pery := Cliente{
		Nome:  "Pery",
		Idade: 32,
		Ativo: true,
	}

	// Declarando uma variável com a struct empresa
	minhaEmpresa := Empresa{}

	// Muda o valor de ativo para false na struct usando interface
	Desativacao(pery)

	// Chama a funcao Desativacao para empresa mostrando que ela implementa a interface Pessoa também
	Desativacao(minhaEmpresa)
}

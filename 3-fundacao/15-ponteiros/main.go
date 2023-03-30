package main

func main() {
	// Memória -> Endereço -> Valor
	a := 10

	// A variável aponta para um ponteiro que é um endereço na memória
	var ponteiro *int = &a
	*ponteiro = 20

	b := &a

	// O asterisco mostra o valor do conteúdo do ponteiro
	// O valor da variável é alterado sem alterar o endereço da memória
	*b = 30
	println(a)
}

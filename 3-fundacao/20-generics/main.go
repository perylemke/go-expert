package main

type MyNumber int

type Number interface {
	// O til signifca que se por um acaso você ter um tipo customizado, ele é aceito, pois é do tipo selecionado abaixo
	~int | ~float64
}

// O T é um tipo que anuncia que é um generic
func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// O comparable somente é pra fins de comparação se é igual.
// Existem packages que podem ajudar na questão de maior ou menor.
func Compara[T comparable](a, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m1 := map[string]int{"Wesley": 1000, "Pery": 2000, "Priscila": 3000}
	m2 := map[string]float64{"Wesley": 1000.98, "Pery": 2000.42, "Priscila": 3000.12}
	m3 := map[string]MyNumber{"Wesley": 1000, "Pery": 2000, "Priscila": 3000}

	println(Soma(m1))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
}

package main

func soma(a, b *int) int {
	*a = 50
	*b = 50
	return *a + *b
}

// Usar ponteiros quando for mexer com valores mútaveis
func main() {
	minhaVar1 := 80
	minhaVar2 := 20
	soma(&minhaVar1, &minhaVar2)
	println(minhaVar1)
	println(minhaVar2)
}

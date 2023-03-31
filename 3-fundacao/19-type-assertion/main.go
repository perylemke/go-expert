package main

import "fmt"

func main() {
	var minhaVar interface{} = "Pery Lemke"
	println(minhaVar.(string))
	res, ok := minhaVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v...", res, ok)

	// O código abaixo irá dar panic, por que não está informado a variável ok para validar.
	// NÃO FAÇA ISSO NUNCA!
	res2 := minhaVar.(int)
	fmt.Printf("O valor de res é %v", res2)
}

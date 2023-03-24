package main

import (
	"fmt"
)

func main() {
	salarios := map[string]int{"Pery": 1000, "João": 2000, "Maria": 3000}
	delete(salarios, "Pery")
	salarios["Pepê"] = 5000

	// sal := make(map[string]int)
	// sal_map := map[string]int{}
	// sal_map["Pery"] = 1000

	for _, salario := range salarios {
		fmt.Printf("O salário é R$ %d\n", salario)
	}
}

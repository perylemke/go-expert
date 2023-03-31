package main

import "fmt"

func main() {
	fmt.Println("For simples...")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("For com Range...")
	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		fmt.Println(k, v)
	}

	fmt.Println("While em Go...")
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Loop infinito em Go...")
	for {
		fmt.Println("Eu sou um loop infinito!")
	}
}

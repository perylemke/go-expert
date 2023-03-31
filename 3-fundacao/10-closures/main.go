package main

import "fmt"

func main() {
	total := func() int {
		return sum(10, 20, 45, 212, 10, 70, 700, 5, 2112, 832, 90, 12, 33) * 2
	}()
	fmt.Println(total)
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

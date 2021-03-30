package main

import (
	"fmt"
)

func main() {

	// Forma padrão de utilizar
	for x := 0; x < 10; x++ {
		fmt.Printf("Valor de X é: %v\n", x)
	}

	fmt.Println("Segunda forma de utilizar o for")
	y := 0
	for y < 10 {
		fmt.Println("Valor de X é:", y)
		y++
	}

	// Infinite loop
	for {
		fmt.Println("Loop infinito")
	}
}

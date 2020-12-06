package main

import "fmt"

//Caso a váriavel não seja inicializada, o valor atribuido é 0

type tomate int

var x tomate

var y int

func main() {
	fmt.Printf("Value X: %v\nType X: %T\n", x, x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("Value Y: %v\nType Y: %T\n", y, y)

}

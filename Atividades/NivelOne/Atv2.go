package main

import "fmt"

//Caso a váriavel não seja inicializada, o valor atribuido é 0

var x int
var y string
var z bool

func main() {
	fmt.Printf("X: %T\nY: %T\nZ: %T\n", x, y, z)
}

package main

import "fmt"

//Caso a váriavel não seja inicializada, o valor atribuido é 0

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%v, %v, %v\n", x, y, z)
	fmt.Print(s)
}

package main

import (
	"fmt"
)

var y = "leandro" //Em qualquer lugar

func main() {

	x := 10 //Apenas em code blocks

	fmt.Printf("Result X: %v %T \n", x, x)

	x, z := 20, 30
	fmt.Printf("Resulte new X: %v %T \n", x, x)
	fmt.Printf("Resulte new Z: %v %T \n", z, z)

	fmt.Printf("Result var y out code block: %v, Type: %T \n", y, y)
}

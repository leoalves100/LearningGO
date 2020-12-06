package main

import "fmt"

var y = "leandro" //Em qualquer lugar

func main() {

	idade := 22 //Apenas em code blocks
	nome := "Leandro Alves"

	concat := fmt.Sprint(idade, " ", nome)

	fmt.Println(concat)

	x, z := 20, 30
	fmt.Printf("Resulte X: %v %T \n", x, x)
	fmt.Printf("Resulte Z: %v %T \n", z, z)

	fmt.Printf("Result var y out code block: %v, Type: %T \n", y, y)

}

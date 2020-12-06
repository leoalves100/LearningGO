package main

import "fmt"

type hotdog int //tipo subjacente

var b hotdog = 101

func main() {
	x := 10
	fmt.Printf("%T %v\n", x, x)

	/*	tipos diferentes, o tipo de b é hotdog
		Apesar seu tipo subjacente ser int, não cria nenhum
		tipo de correlação
	*/

	//Converter tipos T(value)
	x = int(b)
	fmt.Printf("%T %v\n", x, x)

}

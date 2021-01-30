package main

import "fmt"

func main() {
	const nome string = "Leandro Alves"

	const idade = 10

	fmt.Printf("%v %T\n", nome, nome)
	fmt.Printf("%v %T", idade, idade)

}

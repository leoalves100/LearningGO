package main

import "fmt"

func main()  {
	frutas := []string{"goiaba", "manga", "maça", "banana"}

	//Retorna goiaba, manga
	fatia := frutas[0:2]
	fmt.Println("Os dois primeiros valores: ", fatia)

	//Acessa todos os itens
	fatia = frutas[:]
	fmt.Println("Todos os valores", fatia)

	//Remover fruta especifica
	newFrutas := []string{} 
	for _, valor := range fatia {
		if valor != "manga" {
			newFrutas = append(newFrutas, valor)
		}
	}
	
	fmt.Println("Removido fruta especifica: ",newFrutas)
}
package main

import "fmt"

func main()  {
	slice := []int{10,20}
	fmt.Println("Slice:",slice)

	/*
	 * Adiciona novo valor ao slice, 
	 * necessário append para adicionar
	*/
	NewSlice := append(slice, 10)
	fmt.Println("New Slice:",NewSlice)

	//Usando range para percorrer valores
	comidas := []string{"Lasanha","Feijão Tropeiro", "Fígado"}

	for indice, valor := range comidas {
		fmt.Println(indice ,"=>", valor)
	}
}
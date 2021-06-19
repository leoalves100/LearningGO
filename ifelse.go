package main

import "fmt"

func main()  {
	
	// Podemos inicializar valores na instrução if
	if x := 11; x == 10 {
		fmt.Printf("X igual a %v\n",x)
	} else if x < 10{
		fmt.Printf("X nenor que %v\n",x)
	} else {
	    fmt.Printf("X maior que %v\n",x)
	}
}
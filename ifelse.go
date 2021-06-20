package main

import "fmt"

var tipo interface{}
func main()  {
	
	// Podemos inicializar valores na instrução if
	if x := 11; x == 10 {
		fmt.Printf("X igual a %v\n",x)
	} else if x < 10{
		fmt.Printf("X nenor que %v\n",x)
	} else {
	    fmt.Printf("X maior que %v\n",x)
	}

	//Switch case
	x := "b"

	switch x {
	//Case composto
	case "a","b":
		fmt.Print("X = a\n")
		//Continua a execução mesmo após retorna resultado
		fallthrough
	case "c":
		fmt.Print("X = b\n")
	default:
		fmt.Print("Case default\n")
	}

	//Baseado em tipos

	tipo = 23.3
	
	switch tipo.(type) {
		case int:
			fmt.Print("Xiii é inteiro\n")
		case float64:
			fmt.Print("Valor ..\n")
		case bool:
			fmt.Print("Xiii a parada é bool\n")
		case string:
			fmt.Print("String na área\n")
	}
	
}
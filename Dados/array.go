package main

import "fmt"

//Array
var y [5]int

func main()  {
	//Array duas formas de preenchimento
	y[0] = 22
	y[1] = 10
	fmt.Println("Array Y: ",y)
	
	//Ou
	x := [5]int{10,20,30,40,50}
	fmt.Println("Array X: ",x)
}

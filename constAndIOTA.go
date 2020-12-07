package main

import "fmt"

const x = 10

const (
	y = iota
	z = iota
) //Valores sucessivos não tipados y = 0, z = 1

var f float64

func main() {

	f = 10
	fmt.Println(x)
	fmt.Println(y, z)
}

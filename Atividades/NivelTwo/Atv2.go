package main

import "fmt"

func main() {

	a := (10 == 10)
	d := (10 != 0)
	b := (10 < 9)
	c := (10 > 9)
	e := (10 >= 9)
	f := (10 <= 9)

	fmt.Printf("%v\n %v\n %v\n %v\n %v\n %v\n", a, b, c, d, e, f)
}

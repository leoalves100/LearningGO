package main

import "fmt"

type hotdog int //tipo subjacente

var b hotdog = 10

func main() {
	x := 10
	fmt.Printf("%T %v\n", x, x)

	fmt.Printf("%T %v\n", b, b)
}

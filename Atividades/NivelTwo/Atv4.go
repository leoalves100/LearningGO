package main

import "fmt"

func main() {
	idade := 10

	fmt.Printf(" Bin: %b\n Hex:%#x\n Dec:%d\n", idade, idade, idade)

	d := idade << 1

	fmt.Printf("\n Bin: %b\n Hex:%#x\n Dec:%d\n", d, d, d)

}

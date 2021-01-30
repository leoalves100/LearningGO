package main

import "fmt"

func main() {
	const (
		_ = 2022 + iota
		a
		s
		d
		f
	)
	fmt.Print(a)
	fmt.Print("\n", s)
	fmt.Print("\n", d)
	fmt.Print("\n", f)
}

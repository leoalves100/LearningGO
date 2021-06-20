package main

import (
	"fmt"
	"time"
)

func main()  {
	year, _, _ := time.Now().Date()

	for ano := 1998; ano <= year; ano ++{
		fmt.Println(ano)
	}
}
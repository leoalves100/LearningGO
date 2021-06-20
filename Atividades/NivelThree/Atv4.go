package main

import (
	"fmt"
	"time"
)

func main()  {
	
	year, _, _ := time.Now().Date()
	ano := 1998

	for {
		if ano == year + 1  {
			break;
		}

		fmt.Println(ano)
		ano ++
	}
}
package main

import "fmt"

func main() {
	name := "Leandro"
	sb := []byte(name) //Slice of bytes
	fmt.Printf("%v, %T", sb, sb)
	fmt.Println("")

	//Visualizar cada byte
	for _, v := range name {
		fmt.Printf("%v- %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")

	for i := 0; i < len(name); i++ {
		fmt.Printf("%b - %T - %#U - %#x\n", name[i], name[i], name[i], name[i])
	}

	fmt.Println("")

	nameTwo := `Leandro
		alves
	`
	fmt.Println(nameTwo)

}

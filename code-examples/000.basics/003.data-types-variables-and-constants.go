package main

import (
	"fmt"
)

func main() {
	const birthDate = "01/01/1999"
	name, lastName := "John", "Doe"
	var wallet float64 = 541

	fmt.Println(name, lastName, birthDate, wallet)
}
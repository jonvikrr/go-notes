package main

import (
	"fmt"
	"math"
)

func main() {
	formattedPi := fmt.Sprintf("%.2f", math.Pi)
	fmt.Println(`The value of Pi is:
	`, formattedPi)
}
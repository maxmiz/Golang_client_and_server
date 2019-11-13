package main

import (
	"fmt"
	"math"
)

func power(a, b, p float64) float64 {

	if b == 1 {
		return a
	} else {
		return math.Mod(math.Pow(a, b), p)
	}
}

func main() {
	fmt.Println("teste")
}

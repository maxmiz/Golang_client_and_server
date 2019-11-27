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

func power2(a, b, p complex128) complex128 {
	if b == 1 {
		return a
	} else {
		return (a ^ b) % p
	}
}

func main() {
	fmt.Println(power(524, 23, 100))
}

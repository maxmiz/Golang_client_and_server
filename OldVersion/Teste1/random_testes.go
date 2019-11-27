package main

import (
	"fmt"
	"math/cmplx"
	"math/rand"
	"time"
)

type numero struct {
	primo  complex128
	raiz   complex128
	modulo complex128
}

func main() {

	rand.Seed(time.Now().UnixNano())

	var auxiliar numero
	auxiliar.primo = complex(rand.Float64(), rand.Float64())
	auxiliar.raiz = complex(rand.Float64(), rand.Float64())
	auxiliar.modulo = complex(rand.Float64(), rand.Float64())

	fmt.Println(calculo(auxiliar.primo, auxiliar.raiz, auxiliar.modulo))
}

func calculo(primo, raiz, modulo complex128) complex128 {
	aux := cmplx.Pow(raiz, primo)
	return aux - (primo * (aux / primo))
}

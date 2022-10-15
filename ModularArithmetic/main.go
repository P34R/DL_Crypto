package main

import (
	"fmt"
	"github.com/P34R/DL_Crypto/ModularArithmetic/mod"
)

func main() {
	ar := mod.NewModularArithmetic()
	ar.SetModule(132)
	fmt.Println("Linear Equation ", ar.LinearEquation(15))
	fmt.Println("Pow Equation ", ar.PowEquation(5, 15))
	fmt.Println("Mul Equation ", ar.MulEquation(5, 15))
	fmt.Println("Random prime in range ", ar.PrimeRange(5, 150))
}

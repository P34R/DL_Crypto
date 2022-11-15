package main

import (
	"fmt"
	"github.com/P34R/DL_Crypto/ECPoint/EC"
	"math/big"
)

func main() {

	a := EC.BasePointGGet()

	fmt.Println("Scalar multiplication\ni * G (15,13)")
	for i := 0; i < 20; i++ {
		k := EC.ScalarMult(a, *big.NewInt(int64(i)))
		fmt.Print(i, " ")
		EC.PrintlnECPoint(k)
	}

	b := EC.ECPointGen(big.NewInt(2), big.NewInt(10))
	a_pl_b := EC.AddECPoints(a, b)
	fmt.Print("a(15,13) + b(2,10) = ")
	EC.PrintlnECPoint(a_pl_b)

	fmt.Print("double b(2,10) = ")
	db := EC.DoubleECPoints(b)
	EC.PrintlnECPoint(db)
	fmt.Print("Is double b on curve? ", EC.IsOnCurveCheck(db))
}

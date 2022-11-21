package main

import (
	"crypto/elliptic"
	"fmt"
	"github.com/P34R/DL_Crypto/ECPoint/EC"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	//---------------------------------------------TASK 2 WORK EXAMPLE ------------------------------

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
	fmt.Println("Is double b on curve? ", EC.IsOnCurveCheck(db))

	// ---------------------------------------------TASK 2 WORK EXAMPLE END --------------------------

	// ---------------------------------------------TASK 3 WORK EXAMPLE ------------------------------

	// my curve (y^2=x^3 +7 % 17)

	rand.Seed(time.Now().UnixMilli())
	AliceD, BobD := EC.GenerateD(), EC.GenerateD()
	fmt.Println("Da, Db: ")
	fmt.Println(AliceD)
	fmt.Println(BobD)
	AliceH := EC.ComputeH(AliceD, EC.BasePointGGet())
	BobH := EC.ComputeH(BobD, EC.BasePointGGet())
	fmt.Println("Ha, Hb: ")
	EC.PrintlnECPoint(AliceH)
	EC.PrintlnECPoint(BobH)
	AliceS := EC.ScalarMult(BobH, *AliceD)
	BobS := EC.ScalarMult(AliceH, *BobD)
	fmt.Println("Sa, Sb: ")
	//EC.PrintlnECPoint(AliceS)
	//EC.PrintlnECPoint(BobS)
	if (AliceS.X.Cmp(BobS.X)) == 0 && (AliceS.Y.Cmp(BobS.Y)) == 0 {
		fmt.Println("Sa=Sb is true") // there is a chance that result (Sa and Sb) will be 0,0. That's because curve have very few points (y^2=x^3+7 % 17)
	} else {
		panic("failed")
	}

	// Let's try a p256 curve ----------------------------------------------------------------

	p := elliptic.P256()

	//parameters for y^2=x^3+ax+b % module ||| CurveA is a, Module is module, CurveB is b
	EC.Module = p.Params().P
	EC.CurveA = big.NewInt(-3)
	EC.CurveB, _ = big.NewInt(0).SetString("41058363725152142129326129780047268409114441015993725554835256314039467401291", 10)

	BasePointP256 := EC.ECPointGen(p.Params().Gx, p.Params().Gy)
	AliceD, BobD = EC.GenerateD(), EC.GenerateD()
	fmt.Println("Da, Db: ")
	fmt.Println(AliceD)
	fmt.Println(BobD)
	AliceH = EC.ComputeH(AliceD, BasePointP256)
	BobH = EC.ComputeH(BobD, BasePointP256)
	fmt.Println("Ha, Hb: ")
	EC.PrintlnECPoint(AliceH)
	EC.PrintlnECPoint(BobH)
	AliceS = EC.ScalarMult(BobH, *AliceD)
	BobS = EC.ScalarMult(AliceH, *BobD)
	fmt.Println("Sa, Sb: ")
	//EC.PrintlnECPoint(AliceS)
	//EC.PrintlnECPoint(BobS)
	if (AliceS.X.Cmp(BobS.X)) == 0 && (AliceS.Y.Cmp(BobS.Y)) == 0 {
		fmt.Println("Sa=Sb is true")
	} else {
		panic("failed")
	}
	// ---------------------------------------------TASK 3 WORK EXAMPLE END --------------------------
}

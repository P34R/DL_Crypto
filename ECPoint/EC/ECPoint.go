package EC

import (
	"fmt"
	"math/big"
)

//y^2 = x^3+7 mod 17 =======> EC
type ECPoint struct {
	X *big.Int
	Y *big.Int
}

//curve parameters
var curveA = big.NewInt(0)
var curveB = big.NewInt(7)
var module = big.NewInt(17)

func BasePointGGet() (point ECPoint) { // i have chosen a y^2 = x^3+7 (mod 17) EC, and generator for it = (15,13)
	gX := big.NewInt(15)
	gY := big.NewInt(13)
	return ECPoint{gX, gY}
}

func ECPointGen(x, y *big.Int) (point ECPoint) {
	vX := new(big.Int)
	vY := new(big.Int)
	vX.Set(x)
	vY.Set(y)
	point.X = vX // maybe simplified to point.X.Set() maybe?
	point.Y = vY
	return
}

func IsOnCurveCheck(a ECPoint) (c bool) {
	if a.X.Cmp(big.NewInt(0)) == 0 && a.Y.Cmp(big.NewInt(0)) == 0 {
		return true
	}
	vY := new(big.Int)
	vX := new(big.Int)

	temp := big.NewInt(0)

	vX.Set(a.X)
	vY.Set(a.Y)

	vY.Mul(vY, vY)     //y^2
	vY.Mod(vY, module) //y^2 mod 17

	temp.Exp(vX, big.NewInt(3), module) //x^3 %17
	vX.Mul(vX, curveA)                  // a*x
	temp.Add(temp, vX)                  //x^3+ax
	temp.Add(temp, curveB)              // x^3+ax+b (my case x^3+7)

	temp.Mod(temp, module) //x^3+ax+b % module (my case x^3+7 % 17)
	if vY.Cmp(temp) == 0 {
		return true
	}
	return false
	// P ∈ CURVE?
}

func AddECPoints(a, b ECPoint) (c ECPoint) {
	if !IsOnCurveCheck(a) || !IsOnCurveCheck(b) { // there is no sense to add points from different curves
		panic("cannot add points belonging to different curves")
	}
	if a.X.Cmp(big.NewInt(0)) == 0 && a.Y.Cmp(big.NewInt(0)) == 0 {
		return b // 0+b = b
	}
	if b.X.Cmp(big.NewInt(0)) == 0 && b.Y.Cmp(big.NewInt(0)) == 0 {
		return a //a+0 = a
	}
	if a.X.Cmp(b.X) == 0 {
		if a.Y.Cmp(b.Y) == 0 {
			c = DoubleECPoints(a) // if a==b
			return
		}
		//we checked that both points from our curve, so x1=x2, but y1!=y2, then y1=-y2
		return ECPointGen(big.NewInt(0), big.NewInt(0))

	}
	mTop := new(big.Int)
	mBot := new(big.Int)

	mTop.Sub(a.Y, b.Y) // yP-yQ
	mBot.Sub(a.X, b.X) // xP-xQ

	mBot.ModInverse(mBot, module) // (xP-xQ)^-1 % module
	mTop.Mul(mTop, mBot)          // (yP-yQ)*(xP-xQ)^-1
	//	mTop.Mod(mTop, module)        // m % 17
	m2 := new(big.Int)
	m2.Exp(mTop, big.NewInt(2), module) // m^2 % 17

	xR := new(big.Int)
	yR := new(big.Int)
	xR.Sub(m2, a.X)  //m^2-xP
	xR.Sub(xR, b.X)  //(m^2-xP) - xQ
	yR.Sub(xR, a.X)  //xR-xP
	yR.Mul(yR, mTop) //(xR-xP) * m
	yR.Add(yR, a.Y)  // m(xR-xP) + yP
	yR.Mul(yR, big.NewInt(-1))
	xR.Mod(xR, module) // xR % 17
	yR.Mod(yR, module) // yR % 17
	c.X = xR
	c.Y = yR
	// P + Q
	return
}

func DoubleECPoints(a ECPoint) (c ECPoint) {
	x3 := new(big.Int)
	y3 := new(big.Int)

	lambda := new(big.Int)
	lambda2 := new(big.Int)
	lambda.Set(a.X)                   // x1
	lambda.Mul(a.X, a.X)              //x1^2
	lambda.Mul(lambda, big.NewInt(3)) // 3x1^2
	lambda.Add(lambda, curveA)        // 3x^2 +a, my case a =0
	lambda.Mod(lambda, module)        // 3x1^2 % 17

	x3.Mul(a.Y, big.NewInt(2)) // temporary using x3 as 2*y1
	x3.ModInverse(x3, module)

	lambda.Mul(lambda, x3)     // lambda = (3*x1^2 +a ) / (2*y1)
	x3.Mul(a.X, big.NewInt(2)) // x3= 2*x1

	lambda2.Exp(lambda, big.NewInt(2), module) // lambda ^2 %17

	x3.Sub(lambda2, x3) // x3 = lambda^2 - 2*x1

	y3.Sub(a.X, x3)    // x1-x3
	y3.Mul(lambda, y3) // lambda*(x1-x3)
	y3.Sub(y3, a.Y)    // lambda*(x1-x3)-y1
	x3.Mod(x3, module)
	y3.Mod(y3, module)
	c.X, c.Y = x3, y3
	return
}

func ScalarMult(a ECPoint, k big.Int) (c ECPoint) {
	kbits := k.Text(2)            // converting k to bits
	aCopy := ECPointGen(a.X, a.Y) //copying a (so a won't be changed)
	c = ECPointGen(big.NewInt(0), big.NewInt(0))
	for i := len(kbits) - 1; i >= 0; i-- {
		if kbits[i] == '1' {
			c = AddECPoints(c, aCopy)
		}
		aCopy = DoubleECPoints(aCopy)
	}

	// k * P
	return
}
func ECPointToString(point ECPoint) (s string) {
	//Convert point to string
	return "(" + point.X.String() + ", " + point.Y.String() + ")"
}
func PrintECPoint(point ECPoint) {
	//Print point
	fmt.Print(ECPointToString(point))
}
func PrintlnECPoint(point ECPoint) {
	//Print point
	fmt.Println(ECPointToString(point))
}

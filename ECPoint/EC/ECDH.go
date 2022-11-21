package EC

import (
	"crypto/rand"
	"math/big"
)

func GenerateD() *big.Int {
	max := new(big.Int)
	max.Exp(big.NewInt(2), big.NewInt(260), nil).Sub(max, big.NewInt(1)) // just big max integer for rand
	k, _ := rand.Int(rand.Reader, max)
	return k
}
func ComputeH(D *big.Int, G ECPoint) ECPoint {
	return ScalarMult(G, *D)
}

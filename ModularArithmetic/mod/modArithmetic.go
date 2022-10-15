package mod

import (
	"math/rand"
	"strconv"
	"time"
)

type ModularArithmetic struct {
	m int64
}

func NewModularArithmetic() *ModularArithmetic {
	return &ModularArithmetic{
		m: 0,
	}
}

func (ar *ModularArithmetic) SetModule(m uint64) {
	ar.m = int64(m)
}

func (ar *ModularArithmetic) LinearEquation(a int64) int64 {
	if a < 0 {
		return a%ar.m + ar.m
	}
	return a % ar.m
}

func (ar *ModularArithmetic) MulEquation(a, b int64) int64 {
	if gcd(a, ar.m) != 1 {
		return 0
		//May be a lot of answers or zero answers
	}
	eiler := phi(ar.m)

	k := ar.PowEquation(a, eiler-1)
	return b * k % ar.m

}

func (ar *ModularArithmetic) PowEquation(a, b int64) int64 {
	if b == 0 {
		return 1
	}
	p := getBinary(b)
	var d, t int64 = 1, a
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] == '1' {
			d = d * t % ar.m
		}
		t = t * t % ar.m
	}
	return d
}

func (ar *ModularArithmetic) PrimeRange(a, b int64) int64 {
	rand.Seed(time.Now().UnixMilli())
	k := rand.Int63n(b-a) + a
	start := k + 1
	for ; start != k; start = start + 1 {
		if isPrime(start) {
			return start
		}
		if start == b {
			start = a
		}
	}
	return 0
}

func getBinary(a int64) string {
	return strconv.FormatInt(a, 2)
}

func gcd(a, b int64) int64 {
	if a == 0 {
		return b
	}
	for b != 0 {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}

func phi(n int64) int64 {
	result := n
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			result -= result / i
		}
	}
	if n > 1 {
		result -= result / n
	}
	return result
}

func isPrime(n int64) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := int64(5); i*i <= n; i = i + 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

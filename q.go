// Copyright (c) 2015 Leonid Kneller
//
// Exact computational geometry.
package pq

import "math/big"

// Q represents a rational number of arbitrary precision.
type Q struct {
	_r *big.Rat
}

// RtoQ returns a rational number equal to r.
func RtoQ(r *big.Rat) Q {
	if r == nil {
		return Q{new(big.Rat)}
	}
	return Q{new(big.Rat).Set(r)}
}

// ItoQ returns a rational number equal to n.
func ItoQ(n int64) Q {
	return Q{new(big.Rat).SetInt64(n)}
}

// FtoQ returns a rational number equal to f. If f is not finite, a run-time panic occurs.
func FtoQ(f float64) Q {
	r := new(big.Rat)
	if r.SetFloat64(f) == nil {
		panic("not finite")
	}
	return Q{r}
}

func r(x Q) *big.Rat {
	if x._r == nil {
		return new(big.Rat)
	}
	return x._r
}

// Neg returns -x.
func (x Q) Neg() Q {
	return Q{new(big.Rat).Neg(r(x))}
}

// Abs returns |x|.
func (x Q) Abs() Q {
	return Q{new(big.Rat).Abs(r(x))}
}

// Inv returns 1/x.
func (x Q) Inv() Q {
	return Q{new(big.Rat).Inv(r(x))}
}

// Add returns x+y.
func (x Q) Add(y Q) Q {
	return Q{new(big.Rat).Add(r(x), r(y))}
}

// Sub returns x-y.
func (x Q) Sub(y Q) Q {
	return Q{new(big.Rat).Sub(r(x), r(y))}
}

// Mul returns x*y.
func (x Q) Mul(y Q) Q {
	return Q{new(big.Rat).Mul(r(x), r(y))}
}

// Div returns x/y.
func (x Q) Div(y Q) Q {
	return Q{new(big.Rat).Quo(r(x), r(y))}
}

// Sgn returns:
//
//	-1 if x < 0
//	 0 if x = 0
//	+1 if x > 0
func (x Q) Sgn() int {
	return r(x).Sign()
}

// Cmp returns:
//
//	-1 if x < y
//	 0 if x = 0
//	+1 if x > y.
func (x Q) Cmp(y Q) int {
	return r(x).Cmp(r(y))
}

// Rat returns a big.Rat number equal to x.
func (x Q) Rat() *big.Rat {
	if x._r == nil {
		return new(big.Rat)
	}
	return new(big.Rat).Set(x._r)
}

// String returns a string representation of x.
func (x Q) String() string {
	return r(x).String()
}

// Copyright (c) 2015 Leonid Kneller
//
// Exact computational geometry.
package pq

import "math/big"

// Q represents a rational number of arbitrary precision.
type Q struct {
	_r *big.Rat
}

var qzer Q = ItoQ(0)
var qone Q = ItoQ(1)
var qtwo Q = ItoQ(2)

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
//	+1 if x > y
func (x Q) Cmp(y Q) int {
	return r(x).Cmp(r(y))
}

// Max returns max{x,y}.
func (x Q) Max(y Q) Q {
	if r(x).Cmp(r(y)) > 0 {
		return x
	}
	return y
}

// Min returns min{x,y}.
func (x Q) Min(y Q) Q {
	if r(x).Cmp(r(y)) < 0 {
		return x
	}
	return y
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

// Det2x2 computes the determinant of a 2-by-2 matrix.
func Det2x2(a00, a01, a10, a11 Q) Q {
	m01 := (a00.Mul(a11)).Sub(a10.Mul(a01))
	return m01
}

// Det3x3 computes the determinant of a 3-by-3 matrix.
func Det3x3(a00, a01, a02, a10, a11, a12, a20, a21, a22 Q) Q {
	m01 := (a00.Mul(a11)).Sub(a10.Mul(a01))
	m02 := (a00.Mul(a21)).Sub(a20.Mul(a01))
	m12 := (a10.Mul(a21)).Sub(a20.Mul(a11))
	m012 := (m01.Mul(a22)).Sub(m02.Mul(a12)).Add(m12.Mul(a02))
	return m012
}

// Det4x4 computes the determinant of a 4-by-4 matrix.
func Det4x4(a00, a01, a02, a03, a10, a11, a12, a13, a20, a21, a22, a23, a30, a31, a32, a33 Q) Q {
	m01 := (a10.Mul(a01)).Sub(a00.Mul(a11))
	m02 := (a20.Mul(a01)).Sub(a00.Mul(a21))
	m03 := (a30.Mul(a01)).Sub(a00.Mul(a31))
	m12 := (a20.Mul(a11)).Sub(a10.Mul(a21))
	m13 := (a30.Mul(a11)).Sub(a10.Mul(a31))
	m23 := (a30.Mul(a21)).Sub(a20.Mul(a31))
	m012 := (m12.Mul(a02)).Sub(m02.Mul(a12)).Add(m01.Mul(a22))
	m013 := (m13.Mul(a02)).Sub(m03.Mul(a12)).Add(m01.Mul(a32))
	m023 := (m23.Mul(a02)).Sub(m03.Mul(a22)).Add(m02.Mul(a32))
	m123 := (m23.Mul(a12)).Sub(m13.Mul(a22)).Add(m12.Mul(a32))
	m0123 := (m123.Mul(a03)).Sub(m023.Mul(a13)).Add(m013.Mul(a23)).Sub(m012.Mul(a33))
	return m0123
}

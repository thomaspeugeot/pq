// Copyright (c) 2015 Leonid Kneller

package pq

// Vector2q represents a vector with rational coordinates in the 2-dimensional Euclidean plane.
type Vector2q struct {
	x, y Q
}

// XYtoV returns the vector (x,y).
func XYtoV(x, y Q) Vector2q {
	return Vector2q{x, y}
}

// X returns the Cartesian x-coordinate of u.
func (u Vector2q) X() Q {
	return u.x
}

// Y returns the Cartesian y-coordinate of u.
func (u Vector2q) Y() Q {
	return u.y
}

// XY returns the Cartesian coordinates of u.
func (u Vector2q) XY() (x, y Q) {
	return u.x, u.y
}

// Neg returns -u.
func (u Vector2q) Neg() Vector2q {
	return Vector2q{u.x.Neg(), u.y.Neg()}
}

// Add returns u+v.
func (u Vector2q) Add(v Vector2q) Vector2q {
	return Vector2q{u.x.Add(v.x), u.y.Add(v.y)}
}

// Sub returns u-v.
func (u Vector2q) Sub(v Vector2q) Vector2q {
	return Vector2q{u.x.Sub(v.x), u.y.Sub(v.y)}
}

// Mul returns a*u.
func (u Vector2q) Mul(a Q) Vector2q {
	return Vector2q{a.Mul(u.x), a.Mul(u.y)}
}

// Div returns (1/a)*u.
func (u Vector2q) Div(a Q) Vector2q {
	return Vector2q{u.x.Div(a), u.y.Div(a)}
}

// Dot returns the dot (inner) product of u and v.
func (u Vector2q) Dot(v Vector2q) Q {
	return (u.x.Mul(v.x)).Add(u.y.Mul(v.y))
}

// Abs2 returns u.x²+u.y² (L₂ norm squared).
func (u Vector2q) Abs2() Q {
	return (u.x.Mul(u.x)).Add(u.y.Mul(u.y))
}

// MaxAbs returns max{|u.x|,|u.y|} (L∞ norm).
func (u Vector2q) MaxAbs() Q {
	return (u.x.Abs()).Max(u.y.Abs())
}

// SumAbs returns |u.x|+|u.y| (L₁ norm).
func (u Vector2q) SumAbs() Q {
	return (u.x.Abs()).Add(u.y.Abs())
}

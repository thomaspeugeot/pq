// Copyright (c) 2015 Leonid Kneller

package pq

// Vector3q represents a vector with rational coordinates in the 3-dimensional Euclidean space.
type Vector3q struct {
	x, y, z Q
}

// XYZtoV returns the vector (x,y,z).
func XYZtoV(x, y, z Q) Vector3q {
	return Vector3q{x, y, z}
}

// X returns the Cartesian x-coordinate of u.
func (u Vector3q) X() Q {
	return u.x
}

// Y returns the Cartesian y-coordinate of u.
func (u Vector3q) Y() Q {
	return u.y
}

// Z returns the Cartesian z-coordinate of u.
func (u Vector3q) Z() Q {
	return u.z
}

// XYZ returns the Cartesian coordinates of u.
func (u Vector3q) XYZ() (x, y, z Q) {
	return u.x, u.y, u.z
}

// Neg returns -u.
func (u Vector3q) Neg() Vector3q {
	return Vector3q{u.x.Neg(), u.y.Neg(), u.z.Neg()}
}

// Add returns u+v.
func (u Vector3q) Add(v Vector3q) Vector3q {
	return Vector3q{u.x.Add(v.x), u.y.Add(v.y), u.z.Add(v.z)}
}

// Sub returns u-v.
func (u Vector3q) Sub(v Vector3q) Vector3q {
	return Vector3q{u.x.Sub(v.x), u.y.Sub(v.y), u.z.Sub(v.z)}
}

// Mul returns a*u.
func (u Vector3q) Mul(a Q) Vector3q {
	return Vector3q{a.Mul(u.x), a.Mul(u.y), a.Mul(u.z)}
}

// Div returns (1/a)*u.
func (u Vector3q) Div(a Q) Vector3q {
	return Vector3q{u.x.Div(a), u.y.Div(a), u.z.Div(a)}
}

// Dot returns the dot (inner) product of u and v.
func (u Vector3q) Dot(v Vector3q) Q {
	return (u.x.Mul(v.x)).Add(u.y.Mul(v.y).Add(u.z.Mul(v.z)))
}

// Crs returns the cross (outer) product of u and v.
func (u Vector3q) Crs(v Vector3q) Vector3q {
	u1, u2, u3 := u.x, u.y, u.z
	v1, v2, v3 := v.x, v.y, v.z
	w1 := (u2.Mul(v3)).Sub(u3.Mul(v2))
	w2 := (u3.Mul(v1)).Sub(u1.Mul(v3))
	w3 := (u1.Mul(v2)).Sub(u2.Mul(v1))
	return Vector3q{w1, w2, w3}
}

// Abs2 returns u.x²+u.y²+u.z² (L₂ norm squared).
func (u Vector3q) Abs2() Q {
	return (u.x.Mul(u.x)).Add(u.y.Mul(u.y).Add(u.z.Mul(u.z)))
}

// MaxAbs returns max{|u.x|,|u.y|,|u.z|} (L∞ norm).
func (u Vector3q) MaxAbs() Q {
	return (u.x.Abs()).Max(u.y.Abs()).Max(u.z.Abs())
}

// SumAbs returns |u.x|+|u.y|+|u.z| (L₁ norm).
func (u Vector3q) SumAbs() Q {
	return (u.x.Abs()).Add(u.y.Abs()).Add(u.z.Abs())
}

// String returns a string representation of u in the form "(x,y,z)".
func (u Vector3q) String() string {
	return "(" + u.x.String() + "," + u.y.String() + "," + u.z.String() + ")"
}

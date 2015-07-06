// Copyright (c) 2015 Leonid Kneller

package pq

// Point3q represents a point with rational coordinates in the 3-dimensional Euclidean space.
type Point3q struct {
	x, y, z Q
}

// XYZtoP returns the point (x,y,z).
func XYZtoP(x, y, z Q) Point3q {
	return Point3q{x, y, z}
}

// Lift0 lifts a 2-dimensional point to 3 dimensions with z=0.
func (a Point2q) Lift0() Point3q {
	return Point3q{a.x, a.y, qzer}
}

// Lift1 lifts a 2-dimensional point to 3 dimensions with z=1.
func (a Point2q) Lift1() Point3q {
	return Point3q{a.x, a.y, qone}
}

// Lift2 lifts a 2-dimensional point to 3 dimensions with z=x²+y².
func (a Point2q) Lift2() Point3q {
	return Point3q{a.x, a.y, (a.x.Mul(a.x)).Add(a.y.Mul(a.y))}
}

// X returns the Cartesian x-coordinate of a.
func (a Point3q) X() Q {
	return a.x
}

// Y returns the Cartesian y-coordinate of a.
func (a Point3q) Y() Q {
	return a.y
}

// Z returns the Cartesian z-coordinate of a.
func (a Point3q) Z() Q {
	return a.z
}

// XYZ returns the Cartesian coordinates of a.
func (a Point3q) XYZ() (x, y, z Q) {
	return a.x, a.y, a.z
}

// CmpX compares the Cartesian x-coordinates of a and b.
func (a Point3q) CmpX(b Point3q) int {
	return a.x.Cmp(b.x)
}

// CmpY compares the Cartesian y-coordinates of a and b.
func (a Point3q) CmpY(b Point3q) int {
	return a.y.Cmp(b.y)
}

// CmpZ compares the Cartesian z-coordinates of a and b.
func (a Point3q) CmpZ(b Point3q) int {
	return a.z.Cmp(b.z)
}

// CmpXYZ compares the Cartesian coordinates of a and b in xyz-order.
func (a Point3q) CmpXYZ(b Point3q) int {
	if cmpx := a.x.Cmp(b.x); cmpx != 0 {
		return cmpx
	}
	if cmpy := a.y.Cmp(b.y); cmpy != 0 {
		return cmpy
	}
	return a.z.Cmp(b.z)
}

// CmpZYX compares the Cartesian coordinates of a and b in zyx-order.
func (a Point3q) CmpZYX(b Point3q) int {
	if cmpz := a.z.Cmp(b.z); cmpz != 0 {
		return cmpz
	}
	if cmpy := a.y.Cmp(b.y); cmpy != 0 {
		return cmpy
	}
	return a.x.Cmp(b.x)
}

// Dist2 returns the distance squared between a and b.
func (a Point3q) Dist2(b Point3q) Q {
	dx := a.x.Sub(b.x)
	dy := a.y.Sub(b.y)
	dz := a.z.Sub(b.z)
	return (dx.Mul(dx)).Add(dy.Mul(dy)).Add(dz.Mul(dz))
}

// Add returns the point obtained by translating a by u.
func (a Point3q) Add(u Vector3q) Point3q {
	return Point3q{a.x.Add(u.x), a.y.Add(u.y), a.z.Add(u.z)}
}

// Sub returns the point obtained by translating a by -u.
func (a Point3q) Sub(u Vector3q) Point3q {
	return Point3q{a.x.Sub(u.x), a.y.Sub(u.y), a.z.Sub(u.z)}
}

// Midpoint returns the middle of the segment [a,b].
func (a Point3q) Midpoint(b Point3q) Point3q {
	x := (a.x.Add(b.x)).Div(qtwo)
	y := (a.y.Add(b.y)).Div(qtwo)
	z := (a.z.Add(b.z)).Div(qtwo)
	return Point3q{x, y, z}
}

// String returns a string representation of a in the form "(x,y,z)".
func (a Point3q) String() string {
	return "(" + a.x.String() + "," + a.y.String() + "," + a.z.String() + ")"
}

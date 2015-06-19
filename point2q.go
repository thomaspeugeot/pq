// Copyright (c) 2015 Leonid Kneller

package pq

// Point2q represents a point with rational coordinates in the 2-dimensional Euclidean plane.
type Point2q struct {
	x, y Q
}

// XYtoP returns the point (x,y).
func XYtoP(x, y Q) Point2q {
	return Point2q{x, y}
}

// X returns the Cartesian x-coordinate of a.
func (a Point2q) X() Q {
	return a.x
}

// Y returns the Cartesian y-coordinate of a.
func (a Point2q) Y() Q {
	return a.y
}

// XY returns the Cartesian coordinates of a.
func (a Point2q) XY() (x, y Q) {
	return a.x, a.y
}

// CmpX compares the Cartesian x-coordinates of a and b.
func (a Point2q) CmpX(b Point2q) int {
	return a.x.Cmp(b.x)
}

// CmpY compares the Cartesian y-coordinates of a and b.
func (a Point2q) CmpY(b Point2q) int {
	return a.y.Cmp(b.y)
}

// CmpXY compares the Cartesian coordinates of a and b in xy-order.
func (a Point2q) CmpXY(b Point2q) int {
	if cmpx := a.x.Cmp(b.x); cmpx != 0 {
		return cmpx
	}
	return a.y.Cmp(b.y)
}

// CmpYX compares the Cartesian coordinates of a and b in yx-order.
func (a Point2q) CmpYX(b Point2q) int {
	if cmpy := a.y.Cmp(b.y); cmpy != 0 {
		return cmpy
	}
	return a.x.Cmp(b.x)
}

// Dist2 returns the distance squared between a and b.
func (a Point2q) Dist2(b Point2q) Q {
	dx := a.x.Sub(b.x)
	dy := a.y.Sub(b.y)
	return (dx.Mul(dx)).Add(dy.Mul(dy))
}

// Add returns the point obtained by translating a by u.
func (a Point2q) Add(u Vector2q) Point2q {
	return Point2q{a.x.Add(u.x), a.y.Add(u.y)}
}

// Sub returns the point obtained by translating a by -u.
func (a Point2q) Sub(u Vector2q) Point2q {
	return Point2q{a.x.Sub(u.x), a.y.Sub(u.y)}
}

// Orientation returns:
//
//	-1 if (a,b,c) are clockwise
//	 0 if (a,b,c) are collinear
//	+1 if (a,b,c) are counter-clockwise
func (a Point2q) Orientation(b, c Point2q) int {
	det := Det2x2(b.x.Sub(a.x), b.y.Sub(a.y), c.x.Sub(a.x), c.y.Sub(a.y))
	return det.Sgn()
}

// Midpoint returns the middle of the segment [a,b].
func (a Point2q) Midpoint(b Point2q) Point2q {
	x := (a.x.Add(b.x)).Div(qtwo)
	y := (a.y.Add(b.y)).Div(qtwo)
	return Point2q{x, y}
}

// Copyright (c) 2015 Leonid Kneller

package pq

// Point2q represents a point with rational coordinates in the 2-dimensional Euclidean plane.
type Point2q struct {
	x, y Q
}

// X returns the Cartesian x-coordinate of a.
func (a Point2q) X() Q {
	return a.x
}

// Y returns the Cartesian y-coordinate of a.
func (a Point2q) Y() Q {
	return a.y
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

// Add returns the point obtained by translating a by u.
func (a Point2q) Add(u Vector2q) Point2q {
	return Point2q{a.x.Add(u.x), a.y.Add(u.y)}
}

// Sub returns the point obtained by translating a by -u.
func (a Point2q) Sub(u Vector2q) Point2q {
	return Point2q{a.x.Sub(u.x), a.y.Sub(u.y)}
}

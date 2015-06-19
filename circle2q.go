// Copyright (c) 2015 Leonid Kneller

package pq

// Circle2q represents a circle in the 2-dimensional Euclidean plane.
type Circle2q struct {
	cen Point2q
	rsq Q
}

// CR2toCir returns a circle with a given center and radius squared.
func CR2toCir(center Point2q, radius2 Q) Circle2q {
	if radius2.Sgn() < 0 {
		panic("negative radius2")
	}
	return Circle2q{center, radius2}
}

// PPtoCir returns a circle having the segment [a,b] as its diameter.
func PPtoCir(a, b Point2q) Circle2q {
	cen := a.Midpoint(b)
	rsq := a.Dist2(cen)
	return Circle2q{cen, rsq}
}

// PPPtoCir returns a circle passing through given points.
func PPPtoCir(a, b, c Point2q) Circle2q {
	// Test if (a,b,c) are collinear.
	if a.Orientation(b, c) == 0 {
		if a.CmpXY(b) == 0 {
			return PPtoCir(b, c)
		}
		if b.CmpXY(c) == 0 {
			return PPtoCir(c, a)
		}
		if c.CmpXY(a) == 0 {
			return PPtoCir(a, b)
		}
		panic("collinear points")
	}
	//
	// tC2 returns the circumcenter of translated points.
	//
	tC2 := func(dqx, dqy, drx, dry Q) (dcx, dcy Q) {
		r2 := (drx.Mul(drx)).Add(dry.Mul(dry))
		q2 := (dqx.Mul(dqx)).Add(dqy.Mul(dqy))
		den := Det2x2(dqx, dqy, drx, dry).Mul(qtwo)
		dcx = Det2x2(dry, dqy, r2, q2).Div(den)
		dcy = Det2x2(drx, dqx, r2, q2).Div(den).Neg()
		return
	}
	//
	//
	//
	px, py := a.x, a.y
	qx, qy := b.x, b.y
	rx, ry := c.x, c.y
	x, y := tC2(qx.Sub(px), qy.Sub(py), rx.Sub(px), ry.Sub(py))
	x = x.Add(px)
	y = y.Add(py)
	//
	cen := XYtoP(x, y)
	rsq := cen.Dist2(a)
	return Circle2q{cen, rsq}
}

// Center returns the center of c.
func (c Circle2q) Center() Point2q {
	return c.cen
}

// Radius2 returns the radius squared of c.
func (c Circle2q) Radius2() Q {
	return c.rsq
}

// Side returns:
//
//	-1 if a is outside c
//	 0 if a is on c
//	+1 if a is inside c
func (c Circle2q) Side(a Point2q) int {
	return c.rsq.Cmp(c.cen.Dist2(a))
}

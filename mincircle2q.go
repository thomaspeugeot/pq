// Copyright (c) 2015 Leonid Kneller

package pq

import "math/rand"

// MinCircle2q computes the smallest enclosing circle of a collection of points in the plane.
// It implements Welzl's randomized algorithm applied to the convex hull of a given collection of points.
// The function modifies the input ps by reordering it.
//
// Reference: E. Welzl, Smallest enclosing disks (balls and ellipsoids),
// Lecture Notes in Computer Science Volume 555, pp 359-370 (1991).
//
// See: http://dx.doi.org/10.1007/BFb0038202
func MinCircle2q(ps []Point2q) Circle2q {
	lower, upper := ConvHull2q(ps)
	if len(lower) > 1 {
		lower[len(lower)-1] = Point2q{}
		lower = lower[:len(lower)-1]
	}
	if len(upper) > 1 {
		upper[len(upper)-1] = Point2q{}
		upper = upper[:len(upper)-1]
	}
	chull := make([]Point2q, 0)
	chull = append(chull, lower...)
	chull = append(chull, upper...)
	return mindisc0(chull)
}

func mindisc0(ps []Point2q) Circle2q {
	n := len(ps)
	if n == 0 {
		panic("empty point set")
	}
	if n == 1 {
		return PPtoCir(ps[0], ps[0])
	}
	if n == 2 {
		return PPtoCir(ps[0], ps[1])
	}
	//
	shuffle := func(ps []Point2q) {
		for k := len(ps) - 1; k >= 0; k-- {
			i := rand.Intn(k + 1)
			ps[k], ps[i] = ps[i], ps[k]
		}
	}
	//
	shuffle(ps)
	D := PPtoCir(ps[0], ps[1])
	for k := 2; k < n; k++ {
		pk := ps[k]
		if D.Side(pk) < 0 {
			D = mindisc1(ps[:k], pk)
		}
	}
	return D
}

func mindisc1(ps []Point2q, q Point2q) Circle2q {
	D := PPtoCir(ps[0], q)
	n := len(ps)
	for k := 1; k < n; k++ {
		pk := ps[k]
		if D.Side(pk) < 0 {
			D = mindisc2(ps[:k], pk, q)
		}
	}
	return D
}

func mindisc2(ps []Point2q, q1, q2 Point2q) Circle2q {
	D := PPtoCir(q1, q2)
	n := len(ps)
	for k := 0; k < n; k++ {
		pk := ps[k]
		if D.Side(pk) < 0 {
			D = PPPtoCir(q1, q2, pk)
		}
	}
	return D
}

// ParCircle2q computes the smallest enclosing circle of a collection of points in the plane.
// It implements Welzl's randomized algorithm applied to the convex hull of a given collection of points.
// The function modifies the input ps by reordering it.
// If ncpu > 0 then the convex hull computations run in parallel using ncpu goroutines;
// otherwise the convex hull computations run in parallel using runtime.NumCPU() goroutines.
//
// Reference: E. Welzl, Smallest enclosing disks (balls and ellipsoids),
// Lecture Notes in Computer Science Volume 555, pp 359-370 (1991).
//
// See: http://dx.doi.org/10.1007/BFb0038202
func ParMinCircle2q(ncpu int, ps []Point2q) Circle2q {
	lower, upper := ParConvHull2q(ncpu, ps)
	if len(lower) > 1 {
		lower[len(lower)-1] = Point2q{}
		lower = lower[:len(lower)-1]
	}
	if len(upper) > 1 {
		upper[len(upper)-1] = Point2q{}
		upper = upper[:len(upper)-1]
	}
	chull := make([]Point2q, 0)
	chull = append(chull, lower...)
	chull = append(chull, upper...)
	return mindisc0(chull)
}

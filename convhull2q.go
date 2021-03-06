// Copyright (c) 2015 Leonid Kneller

package pq

import (
	"runtime"
	"sort"
	"sync"
)

// ConvHull2q computes the convex hull of a collection of points in the plane.
// It implements Graham's scan algorithm with Andrew's modification. It computes
// both the lower hull and the upper hull. The hull vertices are listed in
// counter-clockwise order. The function modifies the input ps by reordering it.
//
// Reference: R.L. Graham, An efficient algorithm for determining the convex hull of a
// finite planar set, Inform. Process. Lett., 1:132-133 (1972).
//
// See: http://dx.doi.org/10.1016/0020-0190(72)90045-2
//
// Reference: A.M. Andrew, Another efficient algorithm for convex hulls in two dimensions,
// Inform. Process. Lett., 9:216-219 (1979).
//
// See: http://dx.doi.org/10.1016/0020-0190(79)90072-3
func ConvHull2q(ps []Point2q) (lower, upper []Point2q) {
	//
	// Two special cases: n=0 or n=1.
	//
	n := len(ps)
	if n == 0 {
		lower, upper = []Point2q{}, []Point2q{}
		return
	}
	if n == 1 {
		lower, upper = []Point2q{ps[0]}, []Point2q{ps[0]}
		return
	}
	//
	// Sort the input in (x,y)-order.
	//
	sort.Sort(p2qs(ps))
	//
	// noccw(list,p) (list[n-2],list[n-1],p) are not counter-clockwise.
	//
	noccw := func(list []Point2q, p Point2q) bool {
		n := len(list)
		return list[n-2].Orientation(list[n-1], p) <= 0
	}
	//
	// del(list) returns list without its last element.
	//
	del := func(list []Point2q) []Point2q {
		n1 := len(list) - 1
		list[n1] = Point2q{}
		return list[:n1]
	}
	//
	// Build the lower hull.
	//
	lower = make([]Point2q, 0)
	for i := 0; i < n; i++ {
		pi := ps[i]
		for len(lower) > 1 && noccw(lower, pi) {
			lower = del(lower)
		}
		lower = append(lower, pi)
	}
	//
	// Build the upper hull.
	//
	upper = make([]Point2q, 0)
	for i := n - 1; i >= 0; i-- {
		pi := ps[i]
		for len(upper) > 1 && noccw(upper, pi) {
			upper = del(upper)
		}
		upper = append(upper, pi)
	}
	//
	// Special case.
	//
	if len(lower) == 2 && lower[0].CmpXY(lower[1]) == 0 {
		lower[1] = Point2q{}
		lower = lower[:1]
	}
	if len(upper) == 2 && upper[0].CmpXY(upper[1]) == 0 {
		upper[1] = Point2q{}
		upper = upper[:1]
	}
	//
	//
	//
	return
}

// Sort interface implementation.
type p2qs []Point2q

func (a p2qs) Len() int           { return len(a) }
func (a p2qs) Less(i, j int) bool { return a[i].CmpXY(a[j]) < 0 }
func (a p2qs) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// ParConvHull2q computes the convex hull of a collection of points in the plane.
// It implements Graham's scan algorithm with Andrew's modification. It computes
// both the lower hull and the upper hull. The hull vertices are listed in
// counter-clockwise order. The function modifies the input ps by reordering it.
// If ncpu > 0 then computations run in parallel using ncpu goroutines;
// otherwise computations run in parallel using runtime.NumCPU() goroutines.
//
// Reference: R.L. Graham, An efficient algorithm for determining the convex hull of a
// finite planar set, Inform. Process. Lett., 1:132-133 (1972).
//
// See: http://dx.doi.org/10.1016/0020-0190(72)90045-2
//
// Reference: A.M. Andrew, Another efficient algorithm for convex hulls in two dimensions,
// Inform. Process. Lett., 9:216-219 (1979).
//
// See: http://dx.doi.org/10.1016/0020-0190(79)90072-3
func ParConvHull2q(ncpu int, ps []Point2q) (lower, upper []Point2q) {
	if ncpu <= 0 {
		ncpu = runtime.NumCPU()
	}
	n := len(ps)
	//
	// No need to parallelize.
	//
	if n < ncpu {
		return ConvHull2q(ps)
	}
	//
	// Use mu to synchronize appending results to coll.
	//
	var mu sync.Mutex
	var wg sync.WaitGroup
	coll := make([]Point2q, 0)
	//
	// Parallel loop.
	//
	for cpu := 0; cpu < ncpu; cpu++ {
		wg.Add(1)
		first, limit := cpu*n/ncpu, (cpu+1)*n/ncpu
		go func() {
			defer wg.Done()
			lower, upper := ConvHull2q(ps[first:limit])
			mu.Lock()
			coll = append(coll, lower...)
			coll = append(coll, upper...)
			mu.Unlock()
		}()
	}
	//
	// Wait for all goroutines to finish.
	//
	wg.Wait()
	//
	// There exists an algorithm for merging convex hulls, but for now keep it simple.
	//
	return ConvHull2q(coll)
}

package geom

import (
	"sort"
)

type Intersectable interface {
	Intersect(r Ray) []Intersection
	NormalAt(p Point) Vec
}

type Intersection struct {
	Time   float64
	Object Intersectable
}

func NewIntersection(time float64, object Intersectable) Intersection {
	return Intersection{
		time,
		object,
	}
}

func GetHit(intersections []Intersection) (Intersection, bool) {
	if len(intersections) == 0 {
		return Intersection{}, false
	}

	// copy slice
	xs := make([]Intersection, len(intersections))
	copy(xs, intersections)

	// sort copy
	sort.Slice(xs, func(i, j int) bool {
		return xs[i].Time < xs[j].Time
	})

	// grab first non negative value
	for _, i := range xs {
		if i.Time >= 0 {
			return i, true
		}
	}

	return Intersection{}, false
}

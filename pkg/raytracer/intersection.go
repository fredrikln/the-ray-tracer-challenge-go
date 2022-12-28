package raytracer

import "sort"

type Intersectable interface {
	Intersect(Ray) []Intersection
	NormalAt(Point) Vec
	SetMaterial(*Material) Intersectable
	GetMaterial() *Material
	SetTransform(*Matrix) Intersectable
	GetTransform() *Matrix
	GetParent() *Group
}

type Intersection struct {
	Time   float64
	Object *Intersectable
}

func (i *Intersection) PrepareComputations(r Ray) Computations {
	return Computations{}
}

func NewIntersection(time float64, object Intersectable) Intersection {
	return Intersection{
		time,
		&object,
	}
}

func GetHit(xs []Intersection) (Intersection, bool) {
	if len(xs) == 0 {
		return Intersection{}, false
	}

	sort.Sort(IntersectonSorter(xs))

	// grab first non negative value
	for _, i := range xs {
		if i.Time >= 0 {
			return i, true
		}
	}

	return Intersection{}, false
}

package raytracer

import "sort"

type Intersectable interface {
	Intersect(Ray) []Intersection
	LocalIntersect(Ray) []Intersection
	NormalAt(Point, Intersection) Vec
	LocalNormalAt(Point, Intersection) Vec

	SetTransform(*Matrix) Intersectable
	GetTransform() *Matrix

	GetParent() Intersectable
	SetParent(Intersectable) Intersectable

	WorldToObject(Point) Point
	NormalToWorld(Vec) Vec

	Bounds() *BoundingBox
	Divide(int)

	GetNewMaterial() Scatters
}

type Intersection struct {
	Time   float64
	Object *Intersectable
	U      *float64
	V      *float64
}

func (i *Intersection) PrepareComputations(r Ray) Computations {
	return Computations{}
}

func NewIntersection(time float64, object Intersectable) Intersection {
	return Intersection{
		time,
		&object,
		nil,
		nil,
	}
}

func NewIntersectionWithUV(time float64, object Intersectable, u, v float64) Intersection {
	return Intersection{
		time,
		&object,
		&u,
		&v,
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

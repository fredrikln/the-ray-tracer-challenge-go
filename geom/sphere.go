package geom

import "math"

type Sphere struct {
	// origin Vec
}

func NewSphere() Sphere {
	return Sphere{}
}

func (s Sphere) Intersect(ray Ray) []Intersection {
	intersections := make([]Intersection, 0)

	sphereToRay := ray.Origin.Sub(NewPoint(0, 0, 0))

	a := ray.Direction.Dot(ray.Direction)
	b := 2 * ray.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant < 0 {
		return intersections
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	i1 := NewIntersection(t1, s)
	intersections = append(intersections, i1)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	i2 := NewIntersection(t2, s)
	intersections = append(intersections, i2)

	return intersections
}

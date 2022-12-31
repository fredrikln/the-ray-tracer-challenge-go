package raytracer

import (
	"math"
)

type Sphere struct {
	Parent Intersectable
	*object
}

func NewSphere() *Sphere {
	o := newObject()

	s := Sphere{
		object: &o,
	}

	o.parentObject = &s

	return &s
}

func NewGlassSphere() *Sphere {
	o := newObject()
	o.SetNewMaterial(NewDielectric(1.5))

	s := Sphere{
		object: &o,
	}

	o.parentObject = &s

	return &s
}

func (s *Sphere) LocalIntersect(objectRay Ray) []Intersection {
	sphereToRay := objectRay.Origin.Sub(NewPoint(0, 0, 0))

	a := objectRay.Direction.Dot(objectRay.Direction)
	b := 2 * objectRay.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return []Intersection{}
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	i1 := NewIntersection(t1, s)

	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	i2 := NewIntersection(t2, s)

	return []Intersection{
		i1,
		i2,
	}
}

func (s *Sphere) LocalNormalAt(objectPoint Point, i Intersection) Vec {
	return objectPoint.Sub(NewPoint(0, 0, 0))
}

func (s *Sphere) Bounds() *BoundingBox {
	return NewBoundingBoxWithValues(NewPoint(-1, -1, -1), NewPoint(1, 1, 1)).Transform(s.GetTransform())
}

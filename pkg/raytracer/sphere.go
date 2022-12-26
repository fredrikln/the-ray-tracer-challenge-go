package raytracer

import (
	"math"
)

type Sphere struct {
	Transform *Matrix
	Material  *Material
}

func NewSphere() *Sphere {
	return &Sphere{
		NewIdentityMatrix(),
		NewMaterial(),
	}
}

func (s *Sphere) GetMaterial() *Material {
	return s.Material
}

func (s *Sphere) SetMaterial(m *Material) *Sphere {
	s.Material = m

	return s
}

func (s *Sphere) SetTransform(m *Matrix) *Sphere {
	s.Transform = m

	return s
}

func (s *Sphere) Intersect(ray Ray) []Intersection {
	ray2 := ray.Mul(s.Transform.Inverse())

	intersections := make([]Intersection, 0)

	sphereToRay := ray2.Origin.Sub(NewPoint(0, 0, 0))

	a := ray2.Direction.Dot(ray2.Direction)
	b := 2 * ray2.Direction.Dot(sphereToRay)
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

func (s *Sphere) NormalAt(point Point) Vec {
	objectPoint := point.MulMat(s.Transform.Inverse())

	objectNormal := objectPoint.Sub(NewPoint(0, 0, 0))
	worldNormal := objectNormal.MulMat(s.Transform.Inverse().Transpose())

	return worldNormal.Norm()
}

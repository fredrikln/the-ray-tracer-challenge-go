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

func NewGlassSphere() *Sphere {
	return &Sphere{
		NewIdentityMatrix(),
		NewMaterial().SetTransparency(1.0).SetRefractiveIndex(1.5),
	}
}

func (s *Sphere) GetMaterial() *Material {
	return s.Material
}

func (s *Sphere) SetMaterial(m *Material) Intersectable {
	s.Material = m

	return s
}

func (s *Sphere) GetTransform() *Matrix {
	return s.Transform
}
func (s *Sphere) SetTransform(m *Matrix) Intersectable {
	s.Transform = m

	return s
}

func (s *Sphere) Intersect(worldRay Ray) []Intersection {
	localRay := worldRay.Mul(s.Transform.Inverse())

	intersections := make([]Intersection, 0)

	sphereToRay := localRay.Origin.Sub(NewPoint(0, 0, 0))

	a := localRay.Direction.Dot(localRay.Direction)
	b := 2 * localRay.Direction.Dot(sphereToRay)
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

func (s *Sphere) NormalAt(worldPoint Point) Vec {
	objectPoint := worldPoint.MulMat(s.Transform.Inverse())

	objectNormal := objectPoint.Sub(NewPoint(0, 0, 0))
	worldNormal := objectNormal.MulMat(s.Transform.Inverse().Transpose())

	return worldNormal.Norm()
}

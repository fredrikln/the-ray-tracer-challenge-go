package raytracer

import (
	"math"
)

type Sphere struct {
	Transform *Matrix
	Material  *Material
	Parent    Intersectable
}

func NewSphere() *Sphere {
	return &Sphere{
		Transform: NewIdentityMatrix(),
		Material:  NewMaterial(),
	}
}

func NewGlassSphere() *Sphere {
	return &Sphere{
		Transform: NewIdentityMatrix(),
		Material:  NewMaterial().SetTransparency(1.0).SetRefractiveIndex(1.5),
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

func (s *Sphere) GetParent() Intersectable {
	return s.Parent
}
func (s *Sphere) SetParent(p Intersectable) Intersectable {
	s.Parent = p

	return s
}

func (s *Sphere) Intersect(worldRay Ray) []Intersection {
	localRay := worldRay.Mul(s.Transform.Inverse())

	sphereToRay := localRay.Origin.Sub(NewPoint(0, 0, 0))

	a := localRay.Direction.Dot(localRay.Direction)
	b := 2 * localRay.Direction.Dot(sphereToRay)
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

func (s *Sphere) NormalAt(worldPoint Point, i Intersection) Vec {
	objectPoint := s.WorldToObject(worldPoint)
	objectNormal := objectPoint.Sub(NewPoint(0, 0, 0))
	worldNormal := s.NormalToWorld(objectNormal)

	return worldNormal.Norm()
}

func (s *Sphere) WorldToObject(p Point) Point {
	parent := s.GetParent()

	if parent != nil {
		p = parent.WorldToObject(p)
	}

	return s.GetTransform().Inverse().MulPoint(p)
}

func (s *Sphere) NormalToWorld(n Vec) Vec {
	normal := s.GetTransform().Inverse().Transpose().MulVec(n).Norm()

	parent := s.GetParent()

	if parent != nil {
		normal = parent.NormalToWorld(normal)
	}

	return normal
}

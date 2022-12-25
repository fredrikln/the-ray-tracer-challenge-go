package surface

import (
	"math"

	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
	m "github.com/fredrikln/the-ray-tracer-challenge-go/material"
)

type Sphere struct {
	Transform *g.Matrix

	Material *m.Material
}

func NewSphere() *Sphere {
	return &Sphere{
		g.NewIdentityMatrix(),
		m.NewMaterial(),
	}
}

func (s *Sphere) SetMaterial(m *m.Material) *Sphere {
	s.Material = m

	return s
}

func (s *Sphere) SetTransform(m *g.Matrix) *Sphere {
	s.Transform = m

	return s
}

func (s *Sphere) Intersect(ray g.Ray) []g.Intersection {
	ray2 := ray.Mul(s.Transform.Inverse())

	intersections := make([]g.Intersection, 0)

	sphereToRay := ray2.Origin.Sub(g.NewPoint(0, 0, 0))

	a := ray2.Direction.Dot(ray2.Direction)
	b := 2 * ray2.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := math.Pow(b, 2) - 4*a*c

	if discriminant < 0 {
		return intersections
	}

	t1 := (-b - math.Sqrt(discriminant)) / (2 * a)
	i1 := g.NewIntersection(t1, s)
	intersections = append(intersections, i1)
	t2 := (-b + math.Sqrt(discriminant)) / (2 * a)
	i2 := g.NewIntersection(t2, s)
	intersections = append(intersections, i2)

	return intersections
}

func (s *Sphere) NormalAt(point g.Point) g.Vec {
	objectPoint := point.MulMat(s.Transform.Inverse())

	objectNormal := objectPoint.Sub(g.NewPoint(0, 0, 0))
	worldNormal := objectNormal.MulMat(s.Transform.Inverse().Transpose())

	return worldNormal.Norm()
}

package raytracer

import (
	"math"
)

type Cylinder struct {
	Transform *Matrix
	Material  *Material
	Minimum   float64
	Maximum   float64
	Closed    bool
	Parent    Intersectable
}

func NewCylinder() *Cylinder {
	return &Cylinder{
		Transform: NewIdentityMatrix(),
		Material:  NewMaterial(),
		Minimum:   math.Inf(-1),
		Maximum:   math.Inf(1),
		Closed:    false,
	}
}

func NewGlassCylinder() *Cylinder {
	return &Cylinder{
		Transform: NewIdentityMatrix(),
		Material:  NewMaterial().SetTransparency(1.0).SetRefractiveIndex(1.5),
		Minimum:   math.Inf(-1),
		Maximum:   math.Inf(1),
		Closed:    false,
	}
}

func (c *Cylinder) GetMaterial() *Material {
	return c.Material
}
func (c *Cylinder) SetMaterial(m *Material) Intersectable {
	c.Material = m

	return c
}

func (c *Cylinder) GetTransform() *Matrix {
	return c.Transform
}
func (c *Cylinder) SetTransform(m *Matrix) Intersectable {
	c.Transform = m

	return c
}

func (c *Cylinder) GetParent() Intersectable {
	return c.Parent
}
func (c *Cylinder) SetParent(p Intersectable) Intersectable {
	c.Parent = p

	return c
}

func (cy *Cylinder) Intersect(worldRay Ray) []Intersection {
	localRay := worldRay.Mul(cy.Transform.Inverse())

	a := localRay.Direction.X*localRay.Direction.X + localRay.Direction.Z*localRay.Direction.Z
	b := 2*localRay.Origin.X*localRay.Direction.X + 2*localRay.Origin.Z*localRay.Direction.Z
	c := localRay.Origin.X*localRay.Origin.X + localRay.Origin.Z*localRay.Origin.Z - 1

	disc := b*b - 4*a*c

	if disc < 0 {
		return []Intersection{}
	}

	var xs []Intersection

	t0 := (-b - math.Sqrt(disc)) / (2 * a)
	t1 := (-b + math.Sqrt(disc)) / (2 * a)

	if t0 > t1 {
		t0, t1 = t1, t0
	}

	y0 := localRay.Origin.Y + t0*localRay.Direction.Y
	if cy.Minimum < y0 && y0 < cy.Maximum {
		xs = append(xs, NewIntersection(t0, cy))
	}

	y1 := localRay.Origin.Y + t1*localRay.Direction.Y
	if cy.Minimum < y1 && y1 < cy.Maximum {
		xs = append(xs, NewIntersection(t1, cy))
	}

	xs = append(xs, intersectCaps(cy, localRay)...)

	return xs
}

func (c *Cylinder) NormalAt(worldPoint Point, i Intersection) Vec {
	objectPoint := c.WorldToObject(worldPoint)

	var objectNormal Vec

	dist := objectPoint.X*objectPoint.X + objectPoint.Z*objectPoint.Z

	if dist < 1 && objectPoint.Y >= c.Maximum-1e-5 {
		objectNormal = NewVec(0, 1, 0)
	} else if dist < 1 && objectPoint.Y <= c.Minimum+1e-5 {
		objectNormal = NewVec(0, -1, 0)
	} else {
		objectNormal = NewVec(objectPoint.X, 0, objectPoint.Z)
	}

	worldNormal := c.NormalToWorld(objectNormal)

	return worldNormal.Norm()
}

func checkCap(r Ray, t float64) bool {
	x := r.Origin.X + t*r.Direction.X
	z := r.Origin.Z + t*r.Direction.Z

	return (x*x + z*z) <= 1.0
}

func intersectCaps(cy *Cylinder, r Ray) []Intersection {
	var xs []Intersection

	if !cy.Closed || WithinTolerance(r.Direction.Y, 0, 1e-5) {
		return xs
	}

	t := (cy.Minimum - r.Origin.Y) / r.Direction.Y
	if checkCap(r, t) {
		xs = append(xs, NewIntersection(t, cy))
	}

	t1 := (cy.Maximum - r.Origin.Y) / r.Direction.Y
	if checkCap(r, t1) {
		xs = append(xs, NewIntersection(t1, cy))
	}

	return xs
}

func (cy *Cylinder) WorldToObject(p Point) Point {
	parent := cy.GetParent()

	if parent != nil {
		p = parent.WorldToObject(p)
	}

	return cy.GetTransform().Inverse().MulPoint(p)
}

func (cy *Cylinder) NormalToWorld(n Vec) Vec {
	normal := cy.GetTransform().Inverse().Transpose().MulVec(n).Norm()

	parent := cy.GetParent()

	if parent != nil {
		normal = parent.NormalToWorld(normal)
	}

	return normal
}

func (cy *Cylinder) Bounds() *BoundingBox {
	minY := cy.Minimum
	maxY := cy.Maximum

	return NewBoundingBoxWithValues(NewPoint(-1, minY, -1), NewPoint(1, maxY, 1)).Transform(cy.Transform)
}

package raytracer

import (
	"math"
	"sort"

	c "github.com/fredrikln/the-ray-tracer-challenge-go/common"
)

type Cylinder struct {
	Transform *Matrix
	Material  *Material
	Minimum   float64
	Maximum   float64
	Closed    bool
}

func NewCylinder() *Cylinder {
	return &Cylinder{
		NewIdentityMatrix(),
		NewMaterial(),
		math.Inf(-1),
		math.Inf(1),
		false,
	}
}

func NewGlassCylinder() *Cylinder {
	return &Cylinder{
		NewIdentityMatrix(),
		NewMaterial().SetTransparency(1.0).SetRefractiveIndex(1.5),
		math.Inf(-1),
		math.Inf(1),
		false,
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

	sort.Slice(xs, func(i, j int) bool {
		return xs[i].Time < xs[j].Time
	})

	return xs
}

func (c *Cylinder) NormalAt(worldPoint Point) Vec {
	objectPoint := worldPoint.MulMat(c.Transform.Inverse())

	var objectNormal Vec

	dist := objectPoint.X*objectPoint.X + objectPoint.Z*objectPoint.Z

	if dist < 1 && objectPoint.Y >= c.Maximum-1e-5 {
		objectNormal = NewVec(0, 1, 0)
	} else if dist < 1 && objectPoint.Y <= c.Minimum+1e-5 {
		objectNormal = NewVec(0, -1, 0)
	} else {
		objectNormal = NewVec(objectPoint.X, 0, objectPoint.Z)
	}

	worldNormal := objectNormal.MulMat(c.Transform.Inverse().Transpose())

	return worldNormal.Norm()
}

func checkCap(r Ray, t float64) bool {
	x := r.Origin.X + t*r.Direction.X
	z := r.Origin.Z + t*r.Direction.Z

	return (x*x + z*z) <= 1.0
}

func intersectCaps(cy *Cylinder, r Ray) []Intersection {
	var xs []Intersection

	if !cy.Closed || c.WithinTolerance(r.Direction.Y, 0, 1e-5) {
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

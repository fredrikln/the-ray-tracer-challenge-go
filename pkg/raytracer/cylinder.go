package raytracer

import (
	"math"
)

type Cylinder struct {
	Minimum float64
	Maximum float64
	Closed  bool
	*object
}

func NewCylinder() *Cylinder {
	o := newObject()

	c := Cylinder{
		Minimum: math.Inf(-1),
		Maximum: math.Inf(1),
		Closed:  false,
		object:  &o,
	}

	o.parentObject = &c

	return &c
}

func NewGlassCylinder() *Cylinder {
	o := newObject()
	o.SetNewMaterial(NewDielectric(1.5))

	c := Cylinder{
		Minimum: math.Inf(-1),
		Maximum: math.Inf(1),
		Closed:  false,
		object:  &o,
	}

	o.parentObject = &c

	return &c
}

func (cy *Cylinder) LocalIntersect(objectRay Ray) []Intersection {
	a := objectRay.Direction.X*objectRay.Direction.X + objectRay.Direction.Z*objectRay.Direction.Z
	b := 2*objectRay.Origin.X*objectRay.Direction.X + 2*objectRay.Origin.Z*objectRay.Direction.Z
	c := objectRay.Origin.X*objectRay.Origin.X + objectRay.Origin.Z*objectRay.Origin.Z - 1

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

	y0 := objectRay.Origin.Y + t0*objectRay.Direction.Y
	if cy.Minimum < y0 && y0 < cy.Maximum {
		xs = append(xs, NewIntersection(t0, cy))
	}

	y1 := objectRay.Origin.Y + t1*objectRay.Direction.Y
	if cy.Minimum < y1 && y1 < cy.Maximum {
		xs = append(xs, NewIntersection(t1, cy))
	}

	xs = append(xs, intersectCaps(cy, objectRay)...)

	return xs
}

func (c *Cylinder) LocalNormalAt(objectPoint Point, i Intersection) Vec {
	var objectNormal Vec

	dist := objectPoint.X*objectPoint.X + objectPoint.Z*objectPoint.Z

	if dist < 1 && objectPoint.Y >= c.Maximum-1e-5 {
		objectNormal = NewVec(0, 1, 0)
	} else if dist < 1 && objectPoint.Y <= c.Minimum+1e-5 {
		objectNormal = NewVec(0, -1, 0)
	} else {
		objectNormal = NewVec(objectPoint.X, 0, objectPoint.Z)
	}

	return objectNormal.Norm()
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

func (cy *Cylinder) Bounds() *BoundingBox {
	minY := cy.Minimum
	maxY := cy.Maximum

	return NewBoundingBoxWithValues(NewPoint(-1, minY, -1), NewPoint(1, maxY, 1)).Transform(cy.GetTransform())
}

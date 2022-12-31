package raytracer

import (
	"math"
)

type Cone struct {
	Minimum float64
	Maximum float64
	Closed  bool
	*object
}

func NewCone() *Cone {
	o := newObject()

	c := Cone{
		Minimum: math.Inf(-1),
		Maximum: math.Inf(1),
		Closed:  false,
		object:  &o,
	}

	o.parentObject = &c

	return &c
}

func NewGlassCone() *Cone {
	o := newObject()
	o.SetNewMaterial(NewDielectric(1.5))

	c := Cone{
		Minimum: math.Inf(-1),
		Maximum: math.Inf(1),
		Closed:  false,
		object:  &o,
	}

	o.parentObject = &c

	return &c
}

func (co *Cone) LocalIntersect(objectRay Ray) []Intersection {
	a := objectRay.Direction.X*objectRay.Direction.X - objectRay.Direction.Y*objectRay.Direction.Y + objectRay.Direction.Z*objectRay.Direction.Z
	b := 2*objectRay.Origin.X*objectRay.Direction.X - 2*objectRay.Origin.Y*objectRay.Direction.Y + 2*objectRay.Origin.Z*objectRay.Direction.Z
	c := objectRay.Origin.X*objectRay.Origin.X - objectRay.Origin.Y*objectRay.Origin.Y + objectRay.Origin.Z*objectRay.Origin.Z

	disc := b*b - 4*a*c

	if disc < 0 {
		return []Intersection{}
	}

	var xs []Intersection

	if WithinTolerance(a, 0, 1e-5) && !WithinTolerance(b, 0, 1e-5) {
		t := -c / (2 * b)

		xs = append(xs, NewIntersection(t, co))
	}

	t0 := (-b - math.Sqrt(disc)) / (2 * a)
	t1 := (-b + math.Sqrt(disc)) / (2 * a)

	if t0 > t1 {
		t0, t1 = t1, t0
	}

	y0 := objectRay.Origin.Y + t0*objectRay.Direction.Y

	if co.Minimum < y0 && y0 < co.Maximum {
		xs = append(xs, NewIntersection(t0, co))
	}

	y1 := objectRay.Origin.Y + t1*objectRay.Direction.Y
	if co.Minimum < y1 && y1 < co.Maximum {
		xs = append(xs, NewIntersection(t1, co))
	}

	xs = append(xs, intersectCaps2(co, objectRay)...)

	return xs
}

func (co *Cone) LocalNormalAt(objectPoint Point, i Intersection) Vec {
	var objectNormal Vec

	dist := objectPoint.X*objectPoint.X + objectPoint.Z*objectPoint.Z

	if dist < 1 && objectPoint.Y >= co.Maximum-1e-5 {
		objectNormal = NewVec(0, 1, 0)
	} else if dist < 1 && objectPoint.Y <= co.Minimum+1e-5 {
		objectNormal = NewVec(0, -1, 0)
	} else {
		y := math.Sqrt(objectPoint.X*objectPoint.X + objectPoint.Z*objectPoint.Z)

		if objectPoint.Y > 0 {
			y = -y
		}

		objectNormal = NewVec(objectPoint.X, y, objectPoint.Z)
	}

	return objectNormal
}

func checkCap2(r Ray, t float64, radius float64) bool {
	x := r.Origin.X + t*r.Direction.X
	z := r.Origin.Z + t*r.Direction.Z

	return (x*x + z*z) <= math.Abs(radius)
}

func intersectCaps2(co *Cone, r Ray) []Intersection {
	var xs []Intersection

	if !co.Closed || WithinTolerance(r.Direction.Y, 0, 1e-5) {
		return []Intersection{}
	}

	t0 := (co.Minimum - r.Origin.Y) / r.Direction.Y
	if checkCap2(r, t0, co.Minimum) {
		xs = append(xs, NewIntersection(t0, co))
	}

	t1 := (co.Maximum - r.Origin.Y) / r.Direction.Y
	if checkCap2(r, t1, co.Maximum) {
		xs = append(xs, NewIntersection(t1, co))
	}

	return xs
}

func (co *Cone) Bounds() *BoundingBox {
	a := math.Abs(co.Minimum)
	b := math.Abs(co.Maximum)

	limit := math.Max(a, b)

	return NewBoundingBoxWithValues(NewPoint(-limit, co.Minimum, -limit), NewPoint(limit, co.Maximum, limit)).Transform(co.GetTransform())
}

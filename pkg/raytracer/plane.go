package raytracer

import (
	"math"
)

type Plane struct {
	*object
}

func NewPlane() *Plane {
	o := newObject()

	p := Plane{
		object: &o,
	}

	o.parentObject = &p

	return &p
}

func NewGlassPlane() *Plane {
	o := newObject()
	o.SetNewMaterial(NewDielectric(1.5))

	p := Plane{
		object: &o,
	}

	o.parentObject = &p

	return &p
}

func (p *Plane) LocalIntersect(objectRay Ray) []Intersection {
	if WithinTolerance(0, math.Abs(objectRay.Direction.Y), 1e-5) {
		return []Intersection{}
	}

	t := (-objectRay.Origin.Y) / objectRay.Direction.Y

	return []Intersection{NewIntersection(t, p)}
}

func (p *Plane) LocalNormalAt(Point, Intersection) Vec {
	return NewVec(0, 1, 0)
}

func (pl *Plane) WorldToObject(p Point) Point {
	parent := pl.GetParent()

	if parent != nil {
		p = parent.WorldToObject(p)
	}

	return pl.GetTransform().Inverse().MulPoint(p)
}

func (pl *Plane) NormalToWorld(n Vec) Vec {
	normal := pl.GetTransform().Inverse().Transpose().MulVec(n).Norm()

	parent := pl.GetParent()

	if parent != nil {
		normal = parent.NormalToWorld(normal)
	}

	return normal
}

func (pl *Plane) Bounds() *BoundingBox {
	return NewBoundingBoxWithValues(NewPoint(math.Inf(-1), 0, math.Inf(-1)), NewPoint(math.Inf(1), 0, math.Inf(1))).Transform(pl.GetTransform())
}

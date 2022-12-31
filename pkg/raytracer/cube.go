package raytracer

import (
	"math"
)

type Cube struct {
	*object
}

func NewCube() *Cube {
	o := newObject()

	c := Cube{
		object: &o,
	}

	o.parentObject = &c

	return &c
}

func NewGlassCube() *Cube {
	o := newObject()
	o.SetNewMaterial(NewDielectric(1.5))

	c := Cube{
		object: &o,
	}

	o.parentObject = &c

	return &c
}

func (c *Cube) LocalIntersect(objectRay Ray) []Intersection {
	xtmin, xtmax := checkAxis(objectRay.Origin.X, objectRay.Direction.X)
	ytmin, ytmax := checkAxis(objectRay.Origin.Y, objectRay.Direction.Y)
	ztmin, ztmax := checkAxis(objectRay.Origin.Z, objectRay.Direction.Z)

	tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	tmax := math.Min(math.Min(xtmax, ytmax), ztmax)

	if tmin > tmax {
		return []Intersection{}
	}

	return []Intersection{
		NewIntersection(tmin, c),
		NewIntersection(tmax, c),
	}
}

func (c *Cube) LocalNormalAt(objectPoint Point, i Intersection) Vec {
	maxC := math.Max(math.Max(math.Abs(objectPoint.X), math.Abs(objectPoint.Y)), math.Abs(objectPoint.Z))

	var objectNormal Vec

	if maxC == math.Abs(objectPoint.X) {
		objectNormal = NewVec(objectPoint.X, 0, 0)
	} else if maxC == math.Abs(objectPoint.Y) {
		objectNormal = NewVec(0, objectPoint.Y, 0)
	} else {
		objectNormal = NewVec(0, 0, objectPoint.Z)
	}

	return objectNormal
}

func checkAxis(origin, direction float64) (float64, float64) {
	tmin_numerator := -1 - origin
	tmax_numerator := 1 - origin

	var tmin, tmax float64

	if math.Abs(direction) > 1e-5 {
		tmin = tmin_numerator / direction
		tmax = tmax_numerator / direction
	} else {
		tmin = tmin_numerator * math.Inf(1)
		tmax = tmax_numerator * math.Inf(1)
	}

	if tmin > tmax {
		tmin, tmax = tmax, tmin
	}

	return tmin, tmax
}

func (cu *Cube) WorldToObject(p Point) Point {
	parent := cu.GetParent()

	if parent != nil {
		p = parent.WorldToObject(p)
	}

	return cu.GetTransform().Inverse().MulPoint(p)
}

func (cu *Cube) NormalToWorld(n Vec) Vec {
	normal := cu.GetTransform().Inverse().Transpose().MulVec(n).Norm()

	parent := cu.GetParent()

	if parent != nil {
		normal = parent.NormalToWorld(normal)
	}

	return normal
}

func (cu *Cube) Bounds() *BoundingBox {
	return NewBoundingBoxWithValues(NewPoint(-1, -1, -1), NewPoint(1, 1, 1)).Transform(cu.GetTransform())
}

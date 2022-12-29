package raytracer

import (
	"math"
)

type BoundingBox struct {
	Minimum Point
	Maximum Point
}

func NewBoundingBox() *BoundingBox {
	positiveInfinity := math.Inf(1)
	negativeInfinity := math.Inf(-1)

	return &BoundingBox{
		Minimum: NewPoint(positiveInfinity, positiveInfinity, positiveInfinity),
		Maximum: NewPoint(negativeInfinity, negativeInfinity, negativeInfinity),
	}
}

func NewBoundingBoxWithValues(min, max Point) *BoundingBox {
	return &BoundingBox{
		Minimum: min,
		Maximum: max,
	}
}

func (bb *BoundingBox) Add(p Point) {
	if p.X < bb.Minimum.X {
		bb.Minimum.X = p.X
	}
	if p.Y < bb.Minimum.Y {
		bb.Minimum.Y = p.Y
	}
	if p.Z < bb.Minimum.Z {
		bb.Minimum.Z = p.Z
	}

	if p.X > bb.Maximum.X {
		bb.Maximum.X = p.X
	}
	if p.Y > bb.Maximum.Y {
		bb.Maximum.Y = p.Y
	}
	if p.Z > bb.Maximum.Z {
		bb.Maximum.Z = p.Z
	}
}

func (bb *BoundingBox) AddBoundingBox(b BoundingBox) {
	bb.Add(b.Minimum)
	bb.Add(b.Maximum)
}

func (bb *BoundingBox) Contains(p Point) bool {
	return (p.X >= bb.Minimum.X && p.X <= bb.Maximum.X) && (p.Y >= bb.Minimum.Y && p.Y <= bb.Maximum.Y) && p.Z >= bb.Minimum.Z && p.Z <= bb.Maximum.Z
}

func (bb *BoundingBox) ContainsBox(b *BoundingBox) bool {
	return bb.Contains(b.Minimum) && bb.Contains(b.Maximum)
}

func (bb *BoundingBox) Transform(m *Matrix) *BoundingBox {
	p1 := bb.Minimum
	p2 := NewPoint(bb.Minimum.X, bb.Minimum.Y, bb.Maximum.Z)
	p3 := NewPoint(bb.Minimum.X, bb.Maximum.Y, bb.Minimum.Z)
	p4 := NewPoint(bb.Minimum.X, bb.Maximum.Y, bb.Maximum.Z)
	p5 := NewPoint(bb.Maximum.X, bb.Minimum.Y, bb.Minimum.Z)
	p6 := NewPoint(bb.Maximum.X, bb.Maximum.Y, bb.Maximum.Z)
	p7 := NewPoint(bb.Maximum.X, bb.Maximum.Y, bb.Minimum.Z)
	p8 := bb.Maximum

	result := NewBoundingBox()
	result.Add(p1.MulMat(m))
	result.Add(p2.MulMat(m))
	result.Add(p3.MulMat(m))
	result.Add(p4.MulMat(m))
	result.Add(p5.MulMat(m))
	result.Add(p6.MulMat(m))
	result.Add(p7.MulMat(m))
	result.Add(p8.MulMat(m))

	return result
}

func (bb *BoundingBox) Intersect(localRay Ray) bool {
	xtmin, xtmax := bbCheckAxis(localRay.Origin.X, localRay.Direction.X, bb.Minimum.X, bb.Maximum.X)
	ytmin, ytmax := bbCheckAxis(localRay.Origin.Y, localRay.Direction.Y, bb.Minimum.Y, bb.Maximum.Y)
	ztmin, ztmax := bbCheckAxis(localRay.Origin.Z, localRay.Direction.Z, bb.Minimum.Z, bb.Maximum.Z)

	tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	tmax := math.Min(math.Min(xtmax, ytmax), ztmax)

	return tmin <= tmax
}

func bbCheckAxis(origin, direction, min, max float64) (float64, float64) {
	tmin_numerator := min - origin
	tmax_numerator := max - origin

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

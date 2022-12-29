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

func SplitBoundingBox(bb *BoundingBox) (*BoundingBox, *BoundingBox) {
	dx := bb.Maximum.X - bb.Minimum.X
	dy := bb.Maximum.Y - bb.Minimum.Y
	dz := bb.Maximum.Z - bb.Minimum.Z

	greatest := math.Max(dz, math.Max(dx, dy))

	x0, y0, z0 := bb.Minimum.X, bb.Minimum.Y, bb.Minimum.Z
	x1, y1, z1 := bb.Maximum.X, bb.Maximum.Y, bb.Maximum.Z

	if WithinTolerance(dx, greatest, 1e-5) {
		val := x0 + dx/2
		x0, x1 = val, val
	} else if WithinTolerance(dy, greatest, 1e-5) {
		val := y0 + dy/2
		y0, y1 = val, val
	} else {
		val := z0 + dz/2
		z0, z1 = val, val
	}

	midMin := NewPoint(x0, y0, z0)
	midMax := NewPoint(x1, y1, z1)

	left := NewBoundingBoxWithValues(bb.Minimum, midMax)
	right := NewBoundingBoxWithValues(midMin, bb.Maximum)

	return left, right
}

func PartitionChildren(g *Group) ([]Intersectable, []Intersectable) {
	var left, right []Intersectable

	leftBounds, rightBounds := SplitBoundingBox(g.Bounds())

	for i := len(g.Items) - 1; i >= 0; i-- {
		item := g.Items[i]

		if leftBounds.ContainsBox(item.Bounds()) {
			left = append(left, item)
			g.Items = append(g.Items[:i], g.Items[i+1:]...)
		} else if rightBounds.ContainsBox(item.Bounds()) {
			right = append(right, item)
			g.Items = append(g.Items[:i], g.Items[i+1:]...)
		}
	}

	return left, right
}

func MakeSubGroup(g *Group, items []Intersectable) {
	g2 := NewGroup()

	for _, i := range items {
		g2.AddChild(i)
	}

	g.AddChild(g2)
}

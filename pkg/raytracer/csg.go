package raytracer

import (
	"sort"
)

type Operation int64

const (
	Unknown Operation = iota
	Difference
	Union
	Intersect
)

type CSG struct {
	Transform   *Matrix
	Parent      Intersectable
	SavedBounds *BoundingBox

	Operand Operation
	Left    Intersectable
	Right   Intersectable
}

func NewCSG(operation Operation, left Intersectable, right Intersectable) *CSG {
	csg := CSG{
		Operand:   operation,
		Left:      left,
		Right:     right,
		Transform: NewIdentityMatrix(),
		Parent:    nil,
	}

	left.SetParent(&csg)
	right.SetParent(&csg)

	return &csg
}

func (csg *CSG) Intersect(worldRay Ray) []Intersection {
	if !csg.Bounds().Intersect(worldRay) {
		return []Intersection{}
	}

	var result []Intersection

	localRay := worldRay.Mul(csg.Transform.Inverse())

	leftxs := csg.Left.Intersect(localRay)
	result = append(result, leftxs...)
	rightxs := csg.Right.Intersect(localRay)
	result = append(result, rightxs...)

	sort.Sort(IntersectonSorter(result))

	result = csg.FilterIntersections(result)

	return result
}
func (csg *CSG) NormalAt(p Point, i Intersection) Vec {
	panic("should not happen")
}

func (csg *CSG) SetMaterial(*Material) Intersectable {
	panic("SetMaterial should not happen")
}
func (csg *CSG) GetMaterial() *Material {
	panic("GetMaterial should not happen")
}

func (csg *CSG) SetTransform(t *Matrix) Intersectable {
	csg.Transform = t

	return csg
}
func (csg *CSG) GetTransform() *Matrix {
	return csg.Transform
}

func (csg *CSG) GetParent() Intersectable {
	return csg.Parent
}
func (csg *CSG) SetParent(p Intersectable) Intersectable {
	csg.Parent = p

	return csg
}

func (csg *CSG) WorldToObject(p Point) Point {
	parent := csg.GetParent()

	if parent != nil {
		p = parent.WorldToObject(p)
	}

	return csg.GetTransform().Inverse().MulPoint(p)
}

func (csg *CSG) NormalToWorld(n Vec) Vec {
	inv := csg.GetTransform().Inverse()
	trans := inv.Transpose()
	normal := trans.MulVec(n).Norm()

	parent := csg.GetParent()

	if parent != nil {
		normal = parent.NormalToWorld(normal)
	}

	return normal
}

func IntersectionAllowed(op Operation, lhit, inl, inr bool) bool {
	if op == Union {
		return (lhit && !inr) || (!lhit && !inl)
	} else if op == Intersect {
		return (lhit && inr) || (!lhit && inl)
	} else if op == Difference {
		return (lhit && !inr) || (!lhit && inl)
	}

	return false
}

func (csg *CSG) FilterIntersections(xs []Intersection) []Intersection {
	var result []Intersection

	inl := false
	inr := false

	for _, intersection := range xs {
		lhit := false

		switch csg.Left.(type) {
		case *Group:
			lhit = (csg.Left.(*Group)).Includes(*intersection.Object)
		case *CSG:
			lhit = (csg.Left.(*CSG)).Includes(*intersection.Object)
		default:
			lhit = csg.Left == *intersection.Object
		}

		if IntersectionAllowed(csg.Operand, lhit, inl, inr) {
			result = append(result, intersection)
		}

		if lhit {
			inl = !inl
		} else {
			inr = !inr
		}
	}

	return result
}

func (csg *CSG) Includes(object Intersectable) bool {
	switch csg.Left.(type) {
	case *CSG:
		return (*csg.Left.(*CSG)).Includes(object)
	default:
		if csg.Left == object {
			return true
		}
	}

	switch csg.Right.(type) {
	case *CSG:
		return (*csg.Right.(*CSG)).Includes(object)
	default:
		if csg.Right == object {
			return true
		}
	}

	return false
}

func (csg *CSG) Bounds() *BoundingBox {
	if csg.SavedBounds != nil {
		return csg.SavedBounds
	}

	bb := NewBoundingBox()

	bb.AddBoundingBox(*csg.Left.Bounds())
	bb.AddBoundingBox(*csg.Right.Bounds())

	csg.SavedBounds = bb.Transform(csg.Transform)

	return csg.SavedBounds
}

func (csg *CSG) Divide(threshold int) {
	csg.Left.Divide(threshold)
	csg.Right.Divide(threshold)
}

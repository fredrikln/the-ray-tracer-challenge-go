package raytracer

import (
	"fmt"
	"math"
	"testing"
)

type MockShape struct{}

func (MockShape) Intersect(ray Ray) []Intersection {
	return []Intersection{}
}
func (MockShape) NormalAt(point Point, i Intersection) Vec {
	return Vec{}
}
func (MockShape) GetMaterial() *Material {
	return NewMaterial()
}
func (MockShape) SetMaterial(*Material) Intersectable {
	return &MockShape{}
}
func (MockShape) GetTransform() *Matrix {
	return NewIdentityMatrix()
}
func (ms MockShape) SetTransform(*Matrix) Intersectable {
	return &MockShape{}
}
func (ms MockShape) GetParent() Intersectable {
	return nil
}
func (ms MockShape) SetParent(g Intersectable) Intersectable {
	return nil
}
func (ms MockShape) WorldToObject(p Point) Point {
	return p
}
func (ms MockShape) NormalToWorld(n Vec) Vec {
	return n
}
func (ms MockShape) Bounds() *BoundingBox {
	return NewBoundingBoxWithValues(NewPoint(-1, -1, -1), NewPoint(1, 1, 1))
}
func (ms MockShape) Divide(int) {
	return
}

func (ms MockShape) GetNewMaterial() Scatters {
	return NewDiffuse(NewColor(0.8, 0.8, 0.8))
}

func TestIntersection(t *testing.T) {
	element := MockShape{}
	time := 3.5

	intersection := NewIntersection(time, element)

	if intersection.Time != 3.5 || *intersection.Object != element {
		t.Error("Invalid intersection created")
	}
}

func TestGetHitAllPositive(t *testing.T) {
	m := MockShape{}
	i1 := NewIntersection(1, m)
	i2 := NewIntersection(2, m)
	xs := []Intersection{i1, i2}

	i, hit := GetHit(xs)

	if i != i1 || !hit {
		t.Error("Received wrong hit")
	}
}

func TestGetHitSomeNegative(t *testing.T) {
	m := MockShape{}
	i1 := NewIntersection(-1, m)
	i2 := NewIntersection(1, m)
	xs := []Intersection{i1, i2}

	i, hit := GetHit(xs)

	if i != i2 || !hit {
		t.Error("Received wrong hit")
	}
}

func TestGetHitAllNegative(t *testing.T) {
	m := MockShape{}
	i1 := NewIntersection(-2, m)
	i2 := NewIntersection(-1, m)
	xs := []Intersection{i1, i2}

	_, hit := GetHit(xs)

	if hit != false {
		t.Error("Received hit when should not")
	}
}

func TestGetHitGetLowestNonNegative(t *testing.T) {
	s := MockShape{}
	i1 := NewIntersection(5, s)
	i2 := NewIntersection(7, s)
	i3 := NewIntersection(-3, s)
	i4 := NewIntersection(2, s)
	xs := []Intersection{i1, i2, i3, i4}

	i, hit := GetHit(xs)

	if i != i4 || !hit {
		t.Error("Received wrong hit")
	}
}

func TestHitShouldOffsetPoint(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVec(0, 0, 1))

	s := NewSphere().SetTransform(NewTranslation(0, 0, 1))

	i := NewIntersection(5, s)

	comps := PrepareComputations(i, r)

	if !(comps.OverPoint.Z < (1e-5)/2 && comps.Point.Z > comps.OverPoint.Z) {
		t.Error("OverPoint not correct value", comps.OverPoint)
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		point  Point
		result bool
	}{
		{
			NewPoint(5, -2, 0),
			true,
		},
		{
			NewPoint(11, 4, 7),
			true,
		},
		{
			NewPoint(8, 1, 3),
			true,
		},
		{
			NewPoint(3, 0, 3),
			false,
		},
		{
			NewPoint(8, -4, 3),
			false,
		},
		{
			NewPoint(8, 1, -1),
			false,
		},
		{
			NewPoint(13, 1, 3),
			false,
		},
		{
			NewPoint(8, 5, 3),
			false,
		},
		{
			NewPoint(8, 1, 8),
			false,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			bb := NewBoundingBoxWithValues(NewPoint(5, -2, 0), NewPoint(11, 4, 7))

			result := bb.Contains(tC.point)

			if result != tC.result {
				t.Errorf("Got %v, want %v", result, tC.result)
			}
		})
	}
}

func TestBoxContainsBox(t *testing.T) {
	testCases := []struct {
		min    Point
		max    Point
		result bool
	}{
		{
			NewPoint(5, -2, 0),
			NewPoint(11, 4, 7),
			true,
		},
		{
			NewPoint(6, -1, 1),
			NewPoint(10, 3, 6),
			true,
		},
		{
			NewPoint(4, -1, -1),
			NewPoint(10, 3, 6),
			false,
		},
		{
			NewPoint(6, -1, 1),
			NewPoint(12, 5, 8),
			false,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			bb := NewBoundingBoxWithValues(NewPoint(5, -2, 0), NewPoint(11, 4, 7))

			result := bb.ContainsBox(NewBoundingBoxWithValues(tC.min, tC.max))

			if result != tC.result {
				t.Errorf("Got %v, want %v", result, tC.result)
			}
		})
	}
}

func TestTransformingBoundingBox(t *testing.T) {
	bb := NewBoundingBoxWithValues(NewPoint(-1, -1, -1), NewPoint(1, 1, 1))

	matrix := NewRotationX(math.Pi / 4).RotateY(math.Pi / 4)

	result := bb.Transform(matrix)

	if !result.Minimum.Eq(NewPoint(-1.4142, -1.7071, -1.7071)) {
		t.Errorf("Invalid minimum, got %v, want %v", bb.Minimum, NewPoint(-1.4142, -1.7071, -1.7071))
	}
	if !result.Maximum.Eq(NewPoint(1.4142, 1.7071, 1.7071)) {
		t.Errorf("Invalid maximum, got %v, want %v", bb.Maximum, NewPoint(1.4142, 1.7071, 1.7071))
	}
}

func TestSphereBoundingBoxParentSpace(t *testing.T) {
	s := NewSphere()
	s.SetTransform(NewTranslation(1, -3, 5).Scale(0.5, 2, 4))

	bb := s.Bounds()

	if !bb.Minimum.Eq(NewPoint(0.5, -5, 1)) {
		t.Errorf("invalid minimum, got %v, want %v", bb.Minimum, NewPoint(0.5, -5, 1))
	}

	if !bb.Maximum.Eq(NewPoint(1.5, -1, 9)) {
		t.Errorf("invalid minimum, got %v, want %v", bb.Minimum, NewPoint(1.5, -1, 9))
	}
}

func TestGroupHasChildrenBoundingBox(t *testing.T) {
	s := NewSphere()
	s.SetTransform(NewTranslation(2, 5, -3).Scale(2, 2, 2))

	c := NewCylinder()
	c.Minimum = -2
	c.Maximum = 2
	c.SetTransform(NewTranslation(-4, -1, 4).Scale(0.5, 1, 0.5))

	g := NewGroup()
	g.AddChild(s)
	g.AddChild(c)

	bb := g.Bounds()

	if !bb.Minimum.Eq(NewPoint(-4.5, -3, -5)) {
		t.Errorf("Invalid Minimum, got %v, want %v", bb.Minimum, NewPoint(-4.5, -3, -5))
	}
	if !bb.Maximum.Eq(NewPoint(4, 7, 4.5)) {
		t.Errorf("Invalid Maximum, got %v, want %v", bb.Maximum, NewPoint(4, 7, 4.5))
	}
}

func TestCSGHasChildrenBound(t *testing.T) {
	left := NewSphere()
	right := NewSphere()
	right.SetTransform(NewTranslation(2, 3, 4))

	s := NewCSG(Difference, left, right)

	bb := s.Bounds()

	if !bb.Minimum.Eq(NewPoint(-1, -1, -1)) {
		t.Errorf("Invalid Minimum, got %v, want %v", bb.Minimum, NewPoint(-1, -1, -1))
	}
	if !bb.Maximum.Eq(NewPoint(3, 4, 5)) {
		t.Errorf("Invalid Maximum, got %v, want %v", bb.Maximum, NewPoint(3, 4, 5))
	}
}

func TestIntersectingABoundingBox(t *testing.T) {
	testCases := []struct {
		origin    Point
		direction Vec
		result    bool
	}{
		{
			NewPoint(5, 0.5, 0),
			NewVec(-1, 0, 0),
			true,
		},
		{
			NewPoint(-5, 0.5, 0),
			NewVec(1, 0, 0),
			true,
		},
		{
			NewPoint(0.5, 5, 0),
			NewVec(0, -1, 0),
			true,
		},
		{
			NewPoint(0.5, -5, 0),
			NewVec(0, 1, 0),
			true,
		},
		{
			NewPoint(0.5, 0, 5),
			NewVec(0, 0, -1),
			true,
		},
		{
			NewPoint(0.5, 0, -5),
			NewVec(0, 0, 1),
			true,
		},
		{
			NewPoint(0, 0.5, 0),
			NewVec(0, 0, 1),
			true,
		},
		{
			NewPoint(-2, 0, 0),
			NewVec(2, 4, 6),
			false,
		},
		{
			NewPoint(0, -2, 0),
			NewVec(6, 2, 4),
			false,
		},
		{
			NewPoint(0, 0, -2),
			NewVec(4, 6, 2),
			false,
		},
		{
			NewPoint(2, 0, 2),
			NewVec(0, 0, -1),
			false,
		},
		{
			NewPoint(0, 2, 2),
			NewVec(0, -1, 0),
			false,
		},
		{
			NewPoint(2, 2, 0),
			NewVec(-1, 0, 0),
			false,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			bb := NewBoundingBoxWithValues(NewPoint(-1, -1, -1), NewPoint(1, 1, 1))

			r := NewRay(tC.origin, tC.direction)

			result := bb.Intersect(r)

			if result != tC.result {
				t.Errorf("Got %v, want %v", result, tC.result)
			}
		})
	}
}

func TestIntersectingNonCubicBoundingBox(t *testing.T) {
	testCases := []struct {
		point     Point
		direction Vec
		result    bool
	}{
		{
			NewPoint(15, 1, 2),
			NewVec(-1, 0, 0),
			true,
		},
		{
			NewPoint(-5, -1, 4),
			NewVec(1, 0, 0),
			true,
		},
		{
			NewPoint(7, 6, 5),
			NewVec(0, -1, 0),
			true,
		},
		{
			NewPoint(9, -5, 6),
			NewVec(0, 1, 0),
			true,
		},
		{
			NewPoint(8, 2, 12),
			NewVec(0, 0, -1),
			true,
		},
		{
			NewPoint(6, 0, -5),
			NewVec(0, 0, 1),
			true,
		},
		{
			NewPoint(8, 1, 3.5),
			NewVec(0, 0, 1),
			true,
		},
		{
			NewPoint(9, -1, -8),
			NewVec(2, 4, 6),
			false,
		},
		{
			NewPoint(8, 3, -4),
			NewVec(6, 2, 4),
			false,
		},
		{
			NewPoint(9, -1, -2),
			NewVec(4, 6, 2),
			false,
		},
		{
			NewPoint(4, 0, 9),
			NewVec(0, 0, -1),
			false,
		},
		{
			NewPoint(8, 6, -1),
			NewVec(0, -1, 0),
			false,
		},
		{
			NewPoint(12, 5, 4),
			NewVec(-1, 0, 0),
			false,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			bb := NewBoundingBoxWithValues(NewPoint(5, -2, 0), NewPoint(11, 4, 7))

			r := NewRay(tC.point, tC.direction)

			result := bb.Intersect(r)

			if result != tC.result {
				t.Errorf("Got %v, want %v", result, tC.result)
			}
		})
	}
}

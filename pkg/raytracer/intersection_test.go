package raytracer

import (
	"testing"
)

type MockShape struct{}

func (MockShape) Intersect(ray Ray) []Intersection {
	return []Intersection{}
}
func (MockShape) NormalAt(point Point) Vec {
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

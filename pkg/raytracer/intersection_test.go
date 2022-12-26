package raytracer

import (
	"testing"
)

type MockSurface struct{}

func (MockSurface) Intersect(ray Ray) []Intersection {
	return []Intersection{}
}
func (MockSurface) NormalAt(point Point) Vec {
	return Vec{}
}
func (MockSurface) GetMaterial() *Material {
	return NewMaterial()
}

func TestIntersection(t *testing.T) {
	element := MockSurface{}
	time := 3.5

	intersection := NewIntersection(time, element)

	if intersection.Time != 3.5 || *intersection.Object != element {
		t.Error("Invalid intersection created")
	}
}

func TestGetHitAllPositive(t *testing.T) {
	m := MockSurface{}
	i1 := NewIntersection(1, m)
	i2 := NewIntersection(2, m)
	xs := []Intersection{i1, i2}

	i, hit := GetHit(xs)

	if i != i1 || !hit {
		t.Error("Received wrong hit")
	}
}

func TestGetHitSomeNegative(t *testing.T) {
	m := MockSurface{}
	i1 := NewIntersection(-1, m)
	i2 := NewIntersection(1, m)
	xs := []Intersection{i1, i2}

	i, hit := GetHit(xs)

	if i != i2 || !hit {
		t.Error("Received wrong hit")
	}
}

func TestGetHitAllNegative(t *testing.T) {
	m := MockSurface{}
	i1 := NewIntersection(-2, m)
	i2 := NewIntersection(-1, m)
	xs := []Intersection{i1, i2}

	_, hit := GetHit(xs)

	if hit != false {
		t.Error("Received hit when should not")
	}
}

func TestGetHitGetLowestNonNegative(t *testing.T) {
	s := MockSurface{}
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

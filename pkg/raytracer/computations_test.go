package raytracer

import (
	"math"
	"testing"
)

func TestPrepareComputations(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVec(0, 0, 1))
	s := NewSphere()

	i := NewIntersection(4, s)

	comps := PrepareComputations(i, r)

	if (*comps.Object).(*Sphere) != s {
		t.Error("comps wrong object")
	}
	if comps.Time != i.Time {
		t.Errorf("Invalid time, got %v, want %v", comps.Time, i.Time)
	}
	if !comps.Point.Eq(NewPoint(0, 0, -1)) {
		t.Errorf("Invalid point, got %v, want %v", comps.Point, NewPoint(0, 0, -1))
	}
	if !comps.Normalv.Eq(NewVec(0, 0, -1)) {
		t.Errorf("Invalid normal, got %v, want %v", comps.Normalv, NewVec(0, 0, -1))
	}
	if !comps.Eyev.Eq(NewVec(0, 0, -1)) {
		t.Errorf("Invalid normal, got %v, want %v", comps.Normalv, NewVec(0, 0, -1))
	}
}

func TestIntersectionOnOutside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVec(0, 0, 1))
	s := NewSphere()

	i := NewIntersection(4, s)

	comps := PrepareComputations(i, r)

	if comps.Inside != false {
		t.Errorf("Got %v, want %v", comps.Inside, false)
	}
}

func TestIntersectiononInside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 0), NewVec(0, 0, 1))
	s := NewSphere()

	i := NewIntersection(4, s)

	comps := PrepareComputations(i, r)

	if comps.Inside != true {
		t.Errorf("Got %v, want %v", comps.Inside, true)
	}
}

func TestPrecomputeReflectionVector(t *testing.T) {
	s := NewPlane()
	r := NewRay(NewPoint(0, 1, -1), NewVec(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i := NewIntersection(math.Sqrt(2), s)

	comps := PrepareComputations(i, r)

	got := comps.Reflectv
	want := NewVec(0, math.Sqrt(2)/2, math.Sqrt(2)/2)

	if !got.Eq(want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

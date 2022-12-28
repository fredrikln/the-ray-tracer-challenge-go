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

func TestCalculateUnderPoint(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -0.5), NewVec(0, 0, 1))

	s := NewGlassSphere().SetTransform(NewTranslation(0, 0, 1))

	i := NewIntersection(5, s)
	xs := []Intersection{i}

	comps := PrepareComputationsWithHit(i, r, xs)

	if comps.UnderPoint.Z < float64(1e-5)/2.0 {
		t.Error("UnderPoint is not correct")
	}
	if comps.Point.Z > comps.UnderPoint.Z {
		t.Error("Point larger than Underpoint")
	}
}

func TestDetermineReflectanceUnderTotalInternalReflection(t *testing.T) {
	s := NewGlassSphere()
	r := NewRay(NewPoint(0, 0, math.Sqrt(2)/2), NewVec(0, 1, 0))

	xs := []Intersection{
		NewIntersection(-math.Sqrt(2)/2, s),
		NewIntersection(math.Sqrt(2)/2, s),
	}

	comps := PrepareComputationsWithHit(xs[1], r, xs)

	reflectance := Schlick(comps)

	if reflectance != 1.0 {
		t.Errorf("Got %v, want %v", reflectance, 1.0)
	}
}

func TestDetermineReflectanceOfAPerpendicularRay(t *testing.T) {
	s := NewGlassSphere()
	r := NewRay(NewPoint(0, 0, 0), NewVec(0, 1, 0))
	xs := []Intersection{
		NewIntersection(-1, s),
		NewIntersection(1, s),
	}

	comps := PrepareComputationsWithHit(xs[1], r, xs)

	reflectance := Schlick(comps)

	if !WithinTolerance(reflectance, 0.04, 1e-5) {
		t.Errorf("Got %v, want %v", reflectance, 0.04)
	}
}

func TestSchlickWithSmallAngleAndN2LargerThanN1(t *testing.T) {
	s := NewGlassSphere()
	r := NewRay(NewPoint(0, 0.99, -2), NewVec(0, 0, 1))
	xs := []Intersection{
		NewIntersection(1.8589, s),
	}

	comps := PrepareComputationsWithHit(xs[0], r, xs)

	reflectance := Schlick(comps)

	if !WithinTolerance(reflectance, 0.4887308, 1e-5) {
		t.Errorf("Got %v, want %v", reflectance, 0.48873)
	}
}

package raytracer

import (
	"testing"
)

func TestCreatingTriangle(t *testing.T) {
	tr := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))

	if !tr.P1.Eq(NewPoint(0, 1, 0)) {
		t.Error("Invalid p1")
	}
	if !tr.P2.Eq(NewPoint(-1, 0, 0)) {
		t.Error("Invalid p2")
	}
	if !tr.P3.Eq(NewPoint(1, 0, 0)) {
		t.Error("Invalid p3")
	}
	if !tr.E1.Eq(NewVec(-1, -1, 0)) {
		t.Error("Invalid e1")
	}
	if !tr.E2.Eq(NewVec(1, -1, 0)) {
		t.Error("Invalid e2")
	}
	if !tr.Normal.Eq(NewVec(0, 0, -1)) {
		t.Errorf("Invalid normal, got %v, want, %v", tr.Normal, NewVec(0, 0, -1))
	}
}

func TestNormalVectorForTriangle(t *testing.T) {
	tr := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))

	n1 := tr.NormalAt(NewPoint(0, 0.5, 0))
	n2 := tr.NormalAt(NewPoint(-0.5, 0.75, 0))
	n3 := tr.NormalAt(NewPoint(0.5, 0.25, 0))

	if !n1.Eq(tr.Normal) || !n2.Eq(tr.Normal) || !n3.Eq(tr.Normal) {
		t.Error("Invalid normal")
	}
}

func TestIntersectingRayWithTriangle(t *testing.T) {
	tr := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))

	r := NewRay(NewPoint(0, -1, -2), NewVec(0, 1, 0))

	xs := tr.Intersect(r)

	if len(xs) != 0 {
		t.Error("Invalid intersections")
	}
}

func TestRayMissesP1P3Edge(t *testing.T) {
	tr := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))

	r := NewRay(NewPoint(1, 1, -2), NewVec(0, 0, 1))

	xs := tr.Intersect(r)

	if len(xs) != 0 {
		t.Error("Invalid intersections")
	}
}

func TestRayMissesP1P2Edge(t *testing.T) {
	tr := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))

	r := NewRay(NewPoint(-1, 1, -2), NewVec(0, 0, 1))

	xs := tr.Intersect(r)

	if len(xs) != 0 {
		t.Error("Invalid intersections")
	}
}

func TestRayMissesP2P3Edge(t *testing.T) {
	tr := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))

	r := NewRay(NewPoint(0, -1, -2), NewVec(0, 0, 1))

	xs := tr.Intersect(r)

	if len(xs) != 0 {
		t.Error("Invalid intersections")
	}
}

func TestRayStrikesTriangle(t *testing.T) {
	tr := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))

	r := NewRay(NewPoint(0, 0.5, -2), NewVec(0, 0, 1))

	xs := tr.Intersect(r)

	if len(xs) != 1 {
		t.Error("Invalid intersections")
	}

	if xs[0].Time != 2 {
		t.Error("Invalid time")
	}
}

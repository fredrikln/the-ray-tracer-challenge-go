package raytracer

import "testing"

func getSmoothTriangle() *SmoothTriangle {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)

	n1 := NewVec(0, 1, 0)
	n2 := NewVec(-1, 0, 0)
	n3 := NewVec(1, 0, 0)

	t := NewSmoothTriangle(p1, p2, p3, n1, n2, n3)

	return t
}

func TestNewSmoothTriangle(t *testing.T) {
	tr := getSmoothTriangle()

	if !tr.P1.Eq(NewPoint(0, 1, 0)) {
		t.Error("Invalid p1")
	}
	if !tr.P2.Eq(NewPoint(-1, 0, 0)) {
		t.Error("Invalid p2")
	}
	if !tr.P3.Eq(NewPoint(1, 0, 0)) {
		t.Error("Invalid p3")
	}

	if !tr.N1.Eq(NewVec(0, 1, 0)) {
		t.Error("Invalid n1")
	}
	if !tr.N2.Eq(NewVec(-1, 0, 0)) {
		t.Error("Invalid n2")
	}
	if !tr.N3.Eq(NewVec(1, 0, 0)) {
		t.Error("Invalid n3")
	}
}

func TestIntersectionEncapsulatesUV(t *testing.T) {
	tr := NewTriangle(NewPoint(0, 1, 0), NewPoint(-1, 0, 0), NewPoint(1, 0, 0))

	i := NewIntersectionWithUV(3.5, tr, 0.2, 0.4)

	if *i.U != 0.2 {
		t.Error("Invalid U")
	}
	if *i.V != 0.4 {
		t.Error("Invalid V")
	}
}

func TestIntersectionWithSmoothTriangleStoresUV(t *testing.T) {
	tr := getSmoothTriangle()

	r := NewRay(NewPoint(-0.2, 0.3, -2), NewVec(0, 0, 1))

	xs := tr.Intersect(r)

	if !WithinTolerance(*xs[0].U, 0.45, 1e-5) {
		t.Error("Invalid U")
	}

	if !WithinTolerance(*xs[0].V, 0.25, 1e-5) {
		t.Error("Invalid V")
	}
}

func TestSmoothTriangleInterpolatesUsingUV(t *testing.T) {
	tr := getSmoothTriangle()

	i := NewIntersectionWithUV(1, tr, 0.45, 0.25)

	n := tr.NormalAt(NewPoint(0, 0, 0), i)

	if !n.Eq(NewVec(-0.5547, 0.83205, 0)) {
		t.Errorf("Invalid normal, got %v, want %v", n, NewVec(-0.5547, 0.83205, 0))
	}
}

func TestPreparingNormalOnSmoothTriangle(t *testing.T) {
	tr := getSmoothTriangle()

	i := NewIntersectionWithUV(1, tr, 0.45, 0.25)

	r := NewRay(NewPoint(-0.2, 0.3, -2), NewVec(0, 0, 1))

	xs := []Intersection{i}

	comps := PrepareComputationsWithHit(i, r, xs)

	if !comps.Normalv.Eq(NewVec(-0.5547, 0.83205, 0)) {
		t.Errorf("Got %v, want %v", comps.Normalv, NewVec(-0.5547, 0.83205, 0))
	}
}

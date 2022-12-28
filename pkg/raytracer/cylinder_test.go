package raytracer

import (
	"fmt"
	"math"
	"testing"
)

func TestARayMissesACylinder(t *testing.T) {
	testCases := []struct {
		origin    Point
		direction Vec
	}{
		{
			origin:    NewPoint(1, 0, 0),
			direction: NewVec(0, 1, 0),
		},
		{
			origin:    NewPoint(0, 0, 0),
			direction: NewVec(0, 1, 0),
		},
		{
			origin:    NewPoint(0, 0, -5),
			direction: NewVec(1, 1, 1),
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			c := NewCylinder()

			r := NewRay(tC.origin, tC.direction.Norm())

			xs := c.Intersect(r)

			if len(xs) != 0 {
				t.Errorf("Got %v, want %v", len(xs), 0)
			}
		})
	}
}

func TestRayHitsCylinder(t *testing.T) {
	testCases := []struct {
		origin    Point
		direction Vec
		t0        float64
		t1        float64
	}{
		{
			origin:    NewPoint(1, 0, -5),
			direction: NewVec(0, 0, 1),
			t0:        5,
			t1:        5,
		},
		{
			origin:    NewPoint(0, 0, -5),
			direction: NewVec(0, 0, 1),
			t0:        4,
			t1:        6,
		},
		{
			origin:    NewPoint(0.5, 0, -5),
			direction: NewVec(0.1, 1, 1),
			t0:        6.80798,
			t1:        7.08872,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			c := NewCylinder()
			r := NewRay(tC.origin, tC.direction.Norm())

			xs := c.Intersect(r)

			if len(xs) != 2 {
				t.Errorf("count got %v, want %v", len(xs), 2)
			}

			if !WithinTolerance(xs[0].Time, tC.t0, 1e-5) {
				t.Errorf("t0 got %v, want %v", xs[0].Time, tC.t0)
			}

			if !WithinTolerance(xs[1].Time, tC.t1, 1e-5) {
				t.Errorf("t1 got %v, want %v", xs[1].Time, tC.t1)
			}
		})
	}
}

func TestFindingNormalOnCylinder(t *testing.T) {
	testCases := []struct {
		point  Point
		normal Vec
	}{
		{
			point:  NewPoint(1, 0, 0),
			normal: NewVec(1, 0, 0),
		},
		{
			point:  NewPoint(0, 5, -1),
			normal: NewVec(0, 0, -1),
		},
		{
			point:  NewPoint(0, -2, 1),
			normal: NewVec(0, 0, 1),
		},
		{
			point:  NewPoint(-1, 1, 0),
			normal: NewVec(-1, 0, 0),
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			cy := NewCylinder()

			got := cy.NormalAt(tC.point, NewIntersection(1, cy))

			if !got.Eq(tC.normal) {
				t.Errorf("Got %v, want %v", got, tC.normal)
			}
		})
	}
}

func TestDefaultMinMaxForCylinder(t *testing.T) {
	cy := NewCylinder()

	if cy.Minimum != math.Inf(-1) {
		t.Errorf("Invalid minimum, got %v, want %v", cy.Minimum, math.Inf(-1))
	}

	if cy.Maximum != math.Inf(1) {
		t.Errorf("Invalid maximum, got %v, want %v", cy.Maximum, math.Inf(1))
	}
}

func TestIntersectingAConstrainedCylinder(t *testing.T) {
	testCases := []struct {
		point     Point
		direction Vec
		count     int
	}{
		{
			point:     NewPoint(0, 1.5, 0),
			direction: NewVec(0.1, 1, 0),
			count:     0,
		},
		{
			point:     NewPoint(0, 3, -5),
			direction: NewVec(0, 0, 1),
			count:     0,
		},
		{
			point:     NewPoint(0, 0, -5),
			direction: NewVec(0, 0, 1),
			count:     0,
		},
		{
			point:     NewPoint(0, 2, -5),
			direction: NewVec(0, 0, 1),
			count:     0,
		},
		{
			point:     NewPoint(0, 1, -5),
			direction: NewVec(0, 0, 1),
			count:     0,
		},
		{
			point:     NewPoint(0, 1.5, -2),
			direction: NewVec(0, 0, 1),
			count:     2,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			cy := NewCylinder()
			cy.Minimum = 1
			cy.Maximum = 2

			r := NewRay(tC.point, tC.direction.Norm())

			xs := cy.Intersect(r)

			if len(xs) != tC.count {
				t.Errorf("Invalid intersect count, got %v, want %v", len(xs), tC.count)
			}
		})
	}
}

func TestDefaultClosedForCylinder(t *testing.T) {
	cy := NewCylinder()

	if cy.Closed != false {
		t.Error("Cylinder not closed by default")
	}
}

func TestIntersectingCylinderEndCaps(t *testing.T) {
	testCases := []struct {
		point     Point
		direction Vec
		count     int
	}{
		{
			point:     NewPoint(0, 3, 0),
			direction: NewVec(0, -1, 0),
			count:     2,
		},
		{
			point:     NewPoint(0, 3, -2),
			direction: NewVec(0, -1, 2),
			count:     2,
		},
		{
			point:     NewPoint(0, 4, -2),
			direction: NewVec(0, -1, 1),
			count:     2,
		},
		{
			point:     NewPoint(0, 0, -2),
			direction: NewVec(0, 1, 2),
			count:     2,
		},
		{
			point:     NewPoint(0, -1, -2),
			direction: NewVec(0, 1, 1),
			count:     2,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			cy := NewCylinder()
			cy.Minimum = 1
			cy.Maximum = 2
			cy.Closed = true

			r := NewRay(tC.point, tC.direction.Norm())

			xs := cy.Intersect(r)

			if len(xs) != tC.count {
				t.Errorf("Got %v, want %v", len(xs), tC.count)
			}
		})
	}
}

func TestNormalVectorOnEndcaps(t *testing.T) {
	testCases := []struct {
		point  Point
		normal Vec
	}{
		{
			point:  NewPoint(0, 1, 0),
			normal: NewVec(0, -1, 0),
		},
		{
			point:  NewPoint(0.5, 1, 0),
			normal: NewVec(0, -1, 0),
		},
		{
			point:  NewPoint(0, 1, 0.5),
			normal: NewVec(0, -1, 0),
		},
		{
			point:  NewPoint(0, 2, 0),
			normal: NewVec(0, 1, 0),
		},
		{
			point:  NewPoint(0.5, 2, 0),
			normal: NewVec(0, 1, 0),
		},
		{
			point:  NewPoint(0, 2, 0.5),
			normal: NewVec(0, 1, 0),
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			cy := NewCylinder()
			cy.Minimum = 1
			cy.Maximum = 2
			cy.Closed = true

			normal := cy.NormalAt(tC.point, NewIntersection(1, cy))

			if !normal.Eq(tC.normal) {
				t.Errorf("Got %v, want %v", normal, tC.normal)
			}
		})
	}
}

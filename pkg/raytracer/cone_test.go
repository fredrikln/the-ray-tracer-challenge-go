package raytracer

import (
	"fmt"
	"math"
	"testing"

	co "github.com/fredrikln/the-ray-tracer-challenge-go/common"
)

func TestIntersecingAConeWithARay(t *testing.T) {
	testCases := []struct {
		point     Point
		direction Vec
		t0        float64
		t1        float64
	}{
		{
			point:     NewPoint(0, 0, -5),
			direction: NewVec(0, 0, 1),
			t0:        5,
			t1:        5,
		},
		{
			point:     NewPoint(0, 0, -5),
			direction: NewVec(1, 1, 1),
			t0:        8.66025,
			t1:        8.66025,
		},
		{
			point:     NewPoint(1, 1, -5),
			direction: NewVec(-0.5, -1, 1),
			t0:        4.55006,
			t1:        49.44994,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			c := NewCone()

			r := NewRay(tC.point, tC.direction.Norm())

			xs := c.Intersect(r)

			if len(xs) != 2 {
				t.Errorf("Invalid intersection count, got %v, want %v", len(xs), 2)
			}

			if !co.WithinTolerance(xs[0].Time, tC.t0, 1e-5) {
				t.Errorf("t0 got %v, want %v", xs[0].Time, tC.t0)
			}

			if !co.WithinTolerance(xs[1].Time, tC.t1, 1e-5) {
				t.Errorf("t1 got %v, want %v", xs[1].Time, tC.t1)
			}
		})
	}
}

func TestIntersectingAConeWithRayParalellToOneHalf(t *testing.T) {
	c := NewCone()

	r := NewRay(NewPoint(0, 0, -1), NewVec(0, 1, 1).Norm())

	xs := c.Intersect(r)

	if !co.WithinTolerance(xs[0].Time, 0.35355, 1e-5) {
		t.Errorf("Invalid intersection got %v, want %v", xs[0].Time, 0.35355)
	}
}

func TestIntersectingConeEndcap(t *testing.T) {
	testCases := []struct {
		point     Point
		direction Vec
		count     int
	}{
		{
			point:     NewPoint(0, 0, -5),
			direction: NewVec(0, 1, 0),
			count:     0,
		},
		{
			point:     NewPoint(0, 0, -0.25),
			direction: NewVec(0, 1, 1),
			count:     2,
		},
		{
			point:     NewPoint(0, 0, -0.25),
			direction: NewVec(0, 1, 0),
			count:     4,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			c := NewCone()
			c.Minimum = -0.5
			c.Maximum = 0.5
			c.Closed = true

			r := NewRay(tC.point, tC.direction.Norm())

			xs := c.Intersect(r)

			if len(xs) != tC.count {
				t.Errorf("got %v, want %v", len(xs), tC.count)
			}
		})
	}
}

func TestNormalVectorOnCone(t *testing.T) {
	testCases := []struct {
		point  Point
		normal Vec
	}{
		{
			point:  NewPoint(0, 0, 0),
			normal: NewVec(0, 0, 0),
		},
		{
			point:  NewPoint(1, 1, 1),
			normal: NewVec(1, -math.Sqrt(2), 1),
		},
		{
			point:  NewPoint(-1, -1, 0),
			normal: NewVec(-1, 1, 0),
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			co := NewCone()

			normal := co.NormalAt(tC.point)

			if !normal.Eq(tC.normal) {
				t.Errorf("got %v, want %v", normal, tC.normal)
			}
		})
	}
}

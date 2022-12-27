package raytracer

import (
	"fmt"
	"testing"
)

func TestRayIntersectsCube(t *testing.T) {
	testCases := []struct {
		desc      string
		origin    Point
		direction Vec
		t1        float64
		t2        float64
	}{
		{
			desc:      "+x",
			origin:    NewPoint(5, 0.5, 0),
			direction: NewVec(-1, 0, 0),
			t1:        4,
			t2:        6,
		},
		{
			desc:      "-x",
			origin:    NewPoint(-5, 0.5, 0),
			direction: NewVec(1, 0, 0),
			t1:        4,
			t2:        6,
		},
		{
			desc:      "+y",
			origin:    NewPoint(0.5, 5, 0),
			direction: NewVec(0, -1, 0),
			t1:        4,
			t2:        6,
		},
		{
			desc:      "-y",
			origin:    NewPoint(0.5, -5, 0),
			direction: NewVec(0, 1, 0),
			t1:        4,
			t2:        6,
		},
		{
			desc:      "+z",
			origin:    NewPoint(0.5, 0, 5),
			direction: NewVec(0, 0, -1),
			t1:        4,
			t2:        6,
		},
		{
			desc:      "-z",
			origin:    NewPoint(0.5, 0, -5),
			direction: NewVec(0, 0, 1),
			t1:        4,
			t2:        6,
		},
		{
			desc:      "inside",
			origin:    NewPoint(0.5, 0, 0),
			direction: NewVec(0, 0, 1),
			t1:        -1,
			t2:        1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			c := NewCube()

			r := NewRay(tC.origin, tC.direction)

			xs := c.Intersect(r)

			if xs[0].Time != tC.t1 {
				t.Errorf("Invalid t1, got %v, want %v", xs[0].Time, tC.t1)
			}
			if xs[1].Time != tC.t2 {
				t.Errorf("Invalid t2, got %v, want %v", xs[1].Time, tC.t2)
			}
		})
	}
}

func TestRayMissesCube(t *testing.T) {
	testCases := []struct {
		desc      string
		origin    Point
		direction Vec
	}{
		{
			desc:      "1",
			origin:    NewPoint(-2, 0, 0),
			direction: NewVec(0.2673, 0.5345, 0.8018),
		},
		{
			desc:      "2",
			origin:    NewPoint(0, -2, 0),
			direction: NewVec(0.8018, 0.2673, 0.5345),
		},
		{
			desc:      "3",
			origin:    NewPoint(0, 0, -2),
			direction: NewVec(0.5345, 0.8018, 0.2673),
		},
		{
			desc:      "4",
			origin:    NewPoint(2, 0, 2),
			direction: NewVec(0, 0, -1),
		},
		{
			desc:      "5",
			origin:    NewPoint(0, 2, 2),
			direction: NewVec(0, -1, 0),
		},
		{
			desc:      "6",
			origin:    NewPoint(2, 2, 0),
			direction: NewVec(-1, 0, 0),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			c := NewCube()
			r := NewRay(tC.origin, tC.direction)

			xs := c.Intersect(r)

			if len(xs) != 0 {
				t.Errorf("Got %v, want %v", len(xs), 0)
			}
		})
	}
}

func TestFindingNormalOnCube(t *testing.T) {
	testCases := []struct {
		point  Point
		normal Vec
	}{
		{
			point:  NewPoint(1, 0.5, -0.8),
			normal: NewVec(1, 0, 0),
		},
		{
			point:  NewPoint(-1, -0.2, 0.9),
			normal: NewVec(-1, 0, 0),
		},
		{
			point:  NewPoint(-0.4, 1, -0.1),
			normal: NewVec(0, 1, 0),
		},
		{
			point:  NewPoint(0.3, -1, -0.7),
			normal: NewVec(0, -1, 0),
		},
		{
			point:  NewPoint(-0.6, 0.3, 1),
			normal: NewVec(0, 0, 1),
		},
		{
			point:  NewPoint(0.4, 0.4, -1),
			normal: NewVec(0, 0, -1),
		},
		{
			point:  NewPoint(1, 1, 1),
			normal: NewVec(1, 0, 0),
		},
		{
			point:  NewPoint(-1, -1, -1),
			normal: NewVec(-1, 0, 0),
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			c := NewCube()
			p := tC.point

			normal := c.NormalAt(p)

			if !normal.Eq(tC.normal) {
				t.Errorf("Got %v, want %v", normal, tC.normal)
			}
		})
	}
}

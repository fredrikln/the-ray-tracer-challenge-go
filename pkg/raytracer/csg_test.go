package raytracer

import (
	"fmt"
	"testing"
)

func TestNewCSG(t *testing.T) {
	s := NewSphere()
	c := NewCube()

	csg := NewCSG(Union, s, c)

	if csg.Left != s {
		t.Error("Invalid left")
	}

	if csg.Right != c {
		t.Error("Invalid right")
	}

	if s.Parent != csg {
		t.Error("Invalid sphere parent")
	}

	if c.Parent != csg {
		t.Error("Invalid cube parent")
	}
}

func TestEvaluatingRuleForOperation(t *testing.T) {
	testCases := []struct {
		op     Operation
		lhit   bool
		inl    bool
		inr    bool
		result bool
	}{
		{
			Union,
			true,
			true,
			true,
			false,
		},
		{
			Union,
			true,
			true,
			false,
			true,
		},
		{
			Union,
			true,
			false,
			true,
			false,
		},
		{
			Union,
			true,
			false,
			false,
			true,
		},
		{
			Union,
			false,
			true,
			true,
			false,
		},
		{
			Union,
			false,
			true,
			false,
			false,
		},
		{
			Union,
			false,
			false,
			false,
			true,
		},

		{
			Intersect,
			true,
			true,
			true,
			true,
		},
		{
			Intersect,
			true,
			true,
			false,
			false,
		},
		{
			Intersect,
			true,
			false,
			true,
			true,
		},
		{
			Intersect,
			true,
			false,
			false,
			false,
		},
		{
			Intersect,
			false,
			true,
			true,
			true,
		},
		{
			Intersect,
			false,
			true,
			false,
			true,
		},
		{
			Intersect,
			false,
			false,
			true,
			false,
		},
		{
			Intersect,
			false,
			false,
			false,
			false,
		},
		{
			Difference,
			true,
			true,
			true,
			false,
		},
		{
			Difference,
			true,
			true,
			false,
			true,
		},
		{
			Difference,
			true,
			false,
			true,
			false,
		},
		{
			Difference,
			true,
			false,
			false,
			true,
		},
		{
			Difference,
			false,
			true,
			true,
			true,
		},
		{
			Difference,
			false,
			true,
			false,
			true,
		},
		{
			Difference,
			false,
			false,
			true,
			false,
		},
		{
			Difference,
			false,
			false,
			false,
			false,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			result := IntersectionAllowed(tC.op, tC.lhit, tC.inl, tC.inr)

			if result != tC.result {
				t.Errorf("Got %v, want %v", result, tC.result)
			}
		})
	}
}

func TestFilterListOfIntersections(t *testing.T) {
	testCases := []struct {
		op Operation
		x0 int
		x1 int
	}{
		{
			Union,
			0,
			3,
		},
		{
			Intersect,
			1,
			2,
		},
		{
			Difference,
			0,
			1,
		},
	}
	for i, tC := range testCases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			s1 := NewSphere()
			s2 := NewCube()

			csg := NewCSG(tC.op, s1, s2)

			xs := []Intersection{
				NewIntersection(1, s1),
				NewIntersection(2, s2),
				NewIntersection(3, s1),
				NewIntersection(4, s2),
			}

			result := csg.FilterIntersections(xs)

			if result[0] != xs[tC.x0] {
				t.Errorf("Invalid x0 got %v, want %v", result[0], xs[tC.x0])
			}
			if result[0] != xs[tC.x0] {
				t.Errorf("Invalid x1 got %v, want %v", result[1], xs[tC.x0])
			}
		})
	}
}

func TestRayMissesACSGObject(t *testing.T) {
	csg := NewCSG(Union, NewSphere(), NewCube())

	r := NewRay(NewPoint(0, 2, -5), NewVec(0, 0, 1))

	xs := csg.Intersect(r)

	if len(xs) != 0 {
		t.Error("Non empty intersection list")
	}
}

func TestRayHitsCSGObject(t *testing.T) {
	s1 := NewSphere()
	s2 := NewSphere()
	s2.SetTransform(NewTranslation(0, 0, 0.5))
	csg := NewCSG(Union, s1, s2)

	r := NewRay(NewPoint(0, 0, -5), NewVec(0, 0, 1))

	xs := csg.Intersect(r)

	if xs[0].Time != 4 {
		t.Errorf("Got %v, want %v", xs[0].Time, 4)
	}
	if *xs[0].Object != s1 {
		t.Errorf("Got %v, want %v", *xs[0].Object, s1)
	}

	if xs[1].Time != 6.5 {
		t.Errorf("Got %v, want %v", xs[1].Time, 6.5)
	}
	if *xs[1].Object != s2 {
		t.Errorf("Got %v, want %v", *xs[1].Object, s2)
	}
}

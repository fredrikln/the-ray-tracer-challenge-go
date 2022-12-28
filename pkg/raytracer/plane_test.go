package raytracer

import (
	"testing"
)

func TestNormalIsConstant(t *testing.T) {
	testCases := []struct {
		desc  string
		point Point
		want  Vec
	}{
		{
			desc:  "Test 1",
			point: NewPoint(0, 0, 0),
			want:  NewVec(0, 1, 0),
		},
		{
			desc:  "Test 2",
			point: NewPoint(10, 0, -10),
			want:  NewVec(0, 1, 0),
		},
		{
			desc:  "Test 2",
			point: NewPoint(-5, 0, 150),
			want:  NewVec(0, 1, 0),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			p := NewPlane()

			got := p.NormalAt(tC.point, NewIntersection(1, p))
			want := tC.want

			if !got.Eq(want) {
				t.Errorf("Got %v, want %v", got, want)
			}
		})
	}
}

func TestIntersectionsRayMisses(t *testing.T) {
	testCases := []struct {
		desc string
		ray  Ray
		want int
	}{
		{
			desc: "Parallel to plane",
			ray:  NewRay(NewPoint(0, 10, 0), NewVec(0, 0, 1)),
			want: 0,
		},
		{
			desc: "Coplanar to plane",
			ray:  NewRay(NewPoint(0, 0, 0), NewVec(0, 0, 1)),
			want: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			p := NewPlane()

			xs := p.Intersect(tC.ray)
			want := tC.want

			if len(xs) != want {
				t.Errorf("Got %v, want %v", len(xs), want)
			}
		})
	}
}

func TestIntersectFromAbove(t *testing.T) {
	p := NewPlane()
	r := NewRay(NewPoint(0, 1, 0), NewVec(0, -1, 0))

	xs := p.Intersect(r)

	if len(xs) != 1 {
		t.Error("Wrong amount of intersections")
	}

	if xs[0].Time != 1 {
		t.Errorf("Intersection at wrong time, got %v, want %v", xs[0].Time, 0)
	}

	if (*xs[0].Object).(*Plane) != p {
		t.Error("Wrong object intersected")
	}
}

func TestIntersectFroBelow(t *testing.T) {
	p := NewPlane()
	r := NewRay(NewPoint(0, -1, 0), NewVec(0, 1, 0))

	xs := p.Intersect(r)

	if len(xs) != 1 {
		t.Error("Wrong amount of intersections")
	}

	if xs[0].Time != 1 {
		t.Errorf("Intersection at wrong time, got %v, want %v", xs[0].Time, 0)
	}

	if (*xs[0].Object).(*Plane) != p {
		t.Error("Wrong object intersected")
	}
}

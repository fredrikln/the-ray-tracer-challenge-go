package surface

import (
	"math"
	"testing"

	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
	m "github.com/fredrikln/the-ray-tracer-challenge-go/material"
)

func TestSphereIntersectTwoPoints(t *testing.T) {
	ray := g.NewRay(g.NewPoint(0, 0, -5), g.NewVec(0, 0, 1))
	sphere := NewSphere()

	xs := sphere.Intersect(ray)

	if len(xs) != 2 {
		t.Error("Not enough intersections")
		return
	}

	if xs[0].Time != 4.0 || xs[1].Time != 6.0 {
		t.Error("Intersection not att correct times", xs[0].Time, xs[1].Time)
	}
}

func TestSphereIntersectTangent(t *testing.T) {
	ray := g.NewRay(g.NewPoint(0, 1, -5), g.NewVec(0, 0, 1))
	sphere := NewSphere()

	xs := sphere.Intersect(ray)

	if len(xs) != 2 {
		t.Error("Not enough intersections")
		return
	}

	if xs[0].Time != 5.0 || xs[1].Time != 5.0 {
		t.Error("Intersection not att correct times", xs[0].Time, xs[1].Time)
	}
}

func TestSphereIntersectMisses(t *testing.T) {
	ray := g.NewRay(g.NewPoint(0, 2, -5), g.NewVec(0, 0, 1))
	sphere := NewSphere()

	xs := sphere.Intersect(ray)

	if len(xs) != 0 {
		t.Error("Too many intersections")
		return
	}
}

func TestSphereIntersectOriginInside(t *testing.T) {
	ray := g.NewRay(g.NewPoint(0, 0, 0), g.NewVec(0, 0, 1))
	sphere := NewSphere()

	xs := sphere.Intersect(ray)

	if len(xs) != 2 {
		t.Error("Not enough intersections")
		return
	}

	if xs[0].Time != -1.0 || xs[1].Time != 1.0 {
		t.Error("Intersection not att correct times", xs[0].Time, xs[1].Time)
	}
}

func TestSphereIntersectBehind(t *testing.T) {
	ray := g.NewRay(g.NewPoint(0, 0, 5), g.NewVec(0, 0, 1))
	sphere := NewSphere()

	xs := sphere.Intersect(ray)

	if len(xs) != 2 {
		t.Error("Not enough intersections")
		return
	}

	if xs[0].Time != -6.0 || xs[1].Time != -4.0 {
		t.Error("Intersection not att correct times", xs[0].Time, xs[1].Time)
	}
}

func TestSphereDefaultTransform(t *testing.T) {
	s := NewSphere()

	if !s.Transform.Eq(g.NewIdentityMatrix()) {
		t.Errorf("Sphere default transform is wrong, got %v", s.Transform)
	}
}

func TestSetTransform(t *testing.T) {
	s := NewSphere()
	tf := g.NewTranslation(2, 3, 4)
	s.SetTransform(tf)

	if !s.Transform.Eq(tf) {
		t.Errorf("Sphere set transform got wrong, got %v", s.Transform)
	}
}

func TestIntersectScaled(t *testing.T) {
	r := g.NewRay(g.NewPoint(0, 0, -5), g.NewVec(0, 0, 1))
	s := NewSphere()
	s.SetTransform(g.NewScaling(2, 2, 2))

	xs := s.Intersect(r)

	if len(xs) != 2 {
		t.Error("Not enough intersections")
		return
	}

	if xs[0].Time != 3 || xs[1].Time != 7 {
		t.Errorf("Invalid time for intersect, got %f %f, want %f %f", xs[0].Time, xs[1].Time, 3.0, 7.0)
	}
}

func TestIntersectTranslated(t *testing.T) {
	r := g.NewRay(g.NewPoint(0, 0, -5), g.NewVec(0, 0, 1))
	s := NewSphere()
	s.SetTransform(g.NewTranslation(5, 0, 0))

	xs := s.Intersect(r)

	if len(xs) != 0 {
		t.Error("Got too many intersections", xs)
		return
	}
}

func TestSphereNormalAt(t *testing.T) {
	sphere := NewSphere()

	tests := []struct {
		name string
		p    g.Point
		want g.Vec
	}{
		{
			"Test 1",
			g.NewPoint(1, 0, 0),
			g.NewVec(1, 0, 0),
		},
		{
			"Test 2",
			g.NewPoint(0, 1, 0),
			g.NewVec(0, 1, 0),
		},
		{
			"Test 3",
			g.NewPoint(0, 0, 1),
			g.NewVec(0, 0, 1),
		},
		{
			"Test 4",
			g.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
			g.NewVec(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
		},
		{
			"Test 5",
			g.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
			g.NewVec(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3).Norm(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sphere.NormalAt(tt.p); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSphereNormalAtTransformed(t *testing.T) {
	tests := []struct {
		name string
		m    *g.Matrix
		p    g.Point
		want g.Vec
	}{
		{
			"Test 1",
			g.NewTranslation(0, 1, 0),
			g.NewPoint(0, 1.70711, -0.70711),
			g.NewVec(0, 0.70711, -0.70711),
		},
		{
			"Test 2",
			g.NewScaling(1, 0.5, 1).Mul(g.NewRotationZ(math.Pi / 5)),
			g.NewPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			g.NewVec(0, 0.970142, -0.242535),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSphere()
			s.SetTransform(tt.m)

			if got := s.NormalAt(tt.p); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSphereDefaultMaterial(t *testing.T) {
	s := NewSphere()

	if *s.Material != *m.NewMaterial() {
		t.Error("Invalid sphere default material")
	}
}

func TestSphereSetMaterial(t *testing.T) {
	s := NewSphere()

	mat := m.NewMaterial()
	mat.Ambient = 1

	s.Material = mat

	if s.Material != mat {
		t.Error("Invalid sphere material")
	}
}

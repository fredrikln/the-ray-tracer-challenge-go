package tests

import (
	"testing"

	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
	s "github.com/fredrikln/the-ray-tracer-challenge-go/surface"
)

func TestIntersection(t *testing.T) {
	sphere := s.Sphere{}
	time := 3.5

	intersection := g.NewIntersection(time, sphere)

	if intersection.Time != 3.5 || intersection.Object != sphere {
		t.Error("Invalid intersection created")
	}
}

func TestGetHitAllPositive(t *testing.T) {
	s := s.NewSphere()
	i1 := g.NewIntersection(1, s)
	i2 := g.NewIntersection(2, s)
	xs := []g.Intersection{i1, i2}

	i, hit := g.GetHit(xs)

	if i != i1 || !hit {
		t.Error("Received wrong hit")
	}
}

func TestGetHitSomeNegative(t *testing.T) {
	s := s.NewSphere()
	i1 := g.NewIntersection(-1, s)
	i2 := g.NewIntersection(1, s)
	xs := []g.Intersection{i1, i2}

	i, hit := g.GetHit(xs)

	if i != i2 || !hit {
		t.Error("Received wrong hit")
	}
}

func TestGetHitAllNegative(t *testing.T) {
	s := s.NewSphere()
	i1 := g.NewIntersection(-2, s)
	i2 := g.NewIntersection(-1, s)
	xs := []g.Intersection{i1, i2}

	_, hit := g.GetHit(xs)

	if hit != false {
		t.Error("Received hit when should not")
	}
}

func TestGetHitGetLowestNonNegative(t *testing.T) {
	s := s.NewSphere()
	i1 := g.NewIntersection(5, s)
	i2 := g.NewIntersection(7, s)
	i3 := g.NewIntersection(-3, s)
	i4 := g.NewIntersection(2, s)
	xs := []g.Intersection{i1, i2, i3, i4}

	i, hit := g.GetHit(xs)

	if i != i4 || !hit {
		t.Error("Received wrong hit")
	}
}

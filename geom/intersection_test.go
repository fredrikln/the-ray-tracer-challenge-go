package geom

import "testing"

func TestIntersection(t *testing.T) {
	sphere := Sphere{}
	time := 3.5

	intersection := NewIntersection(time, sphere)

	if intersection.Time != 3.5 || intersection.Object != sphere {
		t.Error("Invalid intersection created")
	}
}

func TestGetHitAllPositive(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	xs := []Intersection{i1, i2}

	i, hit := GetHit(xs)

	if i != i1 || !hit {
		t.Error("Received wrong hit")
	}
}

func TestGetHitSomeNegative(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(-1, s)
	i2 := NewIntersection(1, s)
	xs := []Intersection{i1, i2}

	i, hit := GetHit(xs)

	if i != i2 || !hit {
		t.Error("Received wrong hit")
	}
}

func TestGetHitAllNegative(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(-2, s)
	i2 := NewIntersection(-1, s)
	xs := []Intersection{i1, i2}

	_, hit := GetHit(xs)

	if hit != false {
		t.Error("Received hit when should not")
	}
}

func TestGetHitGetLowestNonNegative(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(5, s)
	i2 := NewIntersection(7, s)
	i3 := NewIntersection(-3, s)
	i4 := NewIntersection(2, s)
	xs := []Intersection{i1, i2, i3, i4}

	i, hit := GetHit(xs)

	if i != i4 || !hit {
		t.Error("Received wrong hit")
	}
}

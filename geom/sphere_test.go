package geom

import "testing"

func TestSphereIntersectTwoPoints(t *testing.T) {
	ray := NewRay(NewVec(0, 0, -5), NewVec(0, 0, 1))
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
	ray := NewRay(NewVec(0, 1, -5), NewVec(0, 0, 1))
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
	ray := NewRay(NewVec(0, 2, -5), NewVec(0, 0, 1))
	sphere := NewSphere()

	xs := sphere.Intersect(ray)

	if len(xs) != 0 {
		t.Error("Too many intersections")
		return
	}
}

func TestSphereIntersectOriginInside(t *testing.T) {
	ray := NewRay(NewVec(0, 0, 0), NewVec(0, 0, 1))
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
	ray := NewRay(NewVec(0, 0, 5), NewVec(0, 0, 1))
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

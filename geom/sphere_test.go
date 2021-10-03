package geom

import "testing"

func TestSphereIntersectTwoPoints(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, -5), NewVec(0, 0, 1))
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
	ray := NewRay(NewPoint(0, 1, -5), NewVec(0, 0, 1))
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
	ray := NewRay(NewPoint(0, 2, -5), NewVec(0, 0, 1))
	sphere := NewSphere()

	xs := sphere.Intersect(ray)

	if len(xs) != 0 {
		t.Error("Too many intersections")
		return
	}
}

func TestSphereIntersectOriginInside(t *testing.T) {
	ray := NewRay(NewPoint(0, 0, 0), NewVec(0, 0, 1))
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
	ray := NewRay(NewPoint(0, 0, 5), NewVec(0, 0, 1))
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

	if !s.transform.Eq(NewIdentityMatrix()) {
		t.Errorf("Sphere default transform is wrong, got %v", s.transform)
	}
}

func TestSetTransform(t *testing.T) {
	s := NewSphere()
	tf := NewTranslation(2, 3, 4)
	s.SetTransform(tf)

	if !s.transform.Eq(tf) {
		t.Errorf("Sphere set transform got wrong, got %v", s.transform)
	}
}

func TestIntersectScaled(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVec(0, 0, 1))
	s := NewSphere()
	s.SetTransform(NewScaling(2, 2, 2))

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
	r := NewRay(NewPoint(0, 0, -5), NewVec(0, 0, 1))
	s := NewSphere()
	s.SetTransform(NewTranslation(5, 0, 0))

	xs := s.Intersect(r)

	if len(xs) != 0 {
		t.Error("Got too many intersections", xs)
		return
	}
}

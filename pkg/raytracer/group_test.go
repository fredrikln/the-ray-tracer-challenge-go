package raytracer

import (
	"math"
	"testing"
)

func TestNewGroup(t *testing.T) {
	g := NewGroup()

	if !g.Transform.Eq(NewIdentityMatrix()) {
		t.Errorf("got %v, want %v", g.Transform, NewIdentityMatrix())
	}

	if len(g.Items) != 0 {
		t.Errorf("Got %v, want %v", len(g.Items), 0)
	}
}

func TestShapeHasParent(t *testing.T) {
	s := NewSphere()

	if s.GetParent() != nil {
		t.Error("Parent is not nil")
	}
}

func TestAddingChildToGroup(t *testing.T) {
	g := NewGroup()

	s := NewSphere()

	g.AddChild(s)

	if len(g.Items) != 1 {
		t.Errorf("Got %v, want %v", len(g.Items), 1)
	}

	if g.Items[0] != s {
		t.Error("Invalid item")
	}

	if s.Parent != g {
		t.Error("Invalid parent")
	}
}

func TestIntersectingRayWithEmptyGroup(t *testing.T) {
	g := NewGroup()

	r := NewRay(NewPoint(0, 0, 0), NewVec(0, 0, 1))

	xs := g.Intersect(r)

	if len(xs) != 0 {
		t.Error("Empty group should have no intersections")
	}
}

func TestIntersectingRayWithNonEmptyGroup(t *testing.T) {
	g := NewGroup()

	s1 := NewSphere()
	s2 := NewSphere().SetTransform(NewTranslation(0, 0, -3))
	s3 := NewSphere().SetTransform(NewTranslation(5, 0, 0))
	g.AddChild(s1)
	g.AddChild(s2)
	g.AddChild(s3)

	r := NewRay(NewPoint(0, 0, -5), NewVec(0, 0, 1))

	xs := g.Intersect(r)

	if len(xs) != 4 {
		t.Error("Not enough intersections")
	}

	if *xs[0].Object != s2 {
		t.Error("Wrong first intersection")
	}
	if *xs[1].Object != s2 {
		t.Error("Wrong second intersection")
	}
	if *xs[2].Object != s1 {
		t.Error("Wrong third intersection")
	}
	if *xs[3].Object != s1 {
		t.Error("Wrong fourth intersection")
	}
}

func TestGroupTransformation(t *testing.T) {
	g := NewGroup()
	g.SetTransform(NewScaling(2, 2, 2))

	s := NewSphere().SetTransform(NewTranslation(5, 0, 0))

	g.AddChild(s)

	r := NewRay(NewPoint(10, 0, -10), NewVec(0, 0, 1))

	xs := g.Intersect(r)

	if len(xs) != 2 {
		t.Errorf("Got %v, want %v", len(xs), 2)
	}
}

func TestFindingNormalOnChildObject(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(NewRotationY(math.Pi / 2))

	g2 := NewGroup()
	g2.SetTransform(NewScaling(2, 2, 2))

	g1.AddChild(g2)

	s := NewSphere()
	s.SetTransform(NewTranslation(5, 0, 0))
	g2.AddChild(s)

	p := NewPoint(-2, 0, -10)

	p2 := s.WorldToObject(p)

	if !p2.Eq(NewPoint(0, 0, -1)) {
		t.Errorf("Got %v, want %v", p2, NewPoint(0, 0, -1))
	}
}

func TestConvertNormalFromObjectToWorldSpace(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(NewRotationY(math.Pi / 2))

	g2 := NewGroup()
	g2.SetTransform(NewScaling(1, 2, 3))

	g1.AddChild(g2)

	s := NewSphere()
	s.SetTransform(NewTranslation(5, 0, 0))

	g2.AddChild(s)

	n := s.NormalToWorld(NewVec(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	want := NewVec(0.285714, 0.428571, -0.857142)

	if !n.Eq(want) {
		t.Errorf("Got %v, want %v", n, want)
	}
}

func TestFindNormalOnObjectInGroup(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(NewRotationY(math.Pi / 2))

	g2 := NewGroup()
	g2.SetTransform(NewScaling(1, 2, 3))

	g1.AddChild(g2)

	s := NewSphere()
	s.SetTransform(NewTranslation(5, 0, 0))

	g2.AddChild(s)

	n := s.NormalAt(NewPoint(1.7321, 1.1547, -5.5774), NewIntersection(1, s))
	want := NewVec(0.285703, 0.428543, -0.857160)

	if !n.Eq(want) {
		t.Errorf("Got %v, want %v", n, want)
	}
}

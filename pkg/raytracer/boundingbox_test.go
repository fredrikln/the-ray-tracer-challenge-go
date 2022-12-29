package raytracer

import (
	"math"
	"testing"
)

func TestNewBoundingBox(t *testing.T) {
	bb := NewBoundingBox()

	if !bb.Minimum.Eq(NewPoint(math.Inf(1), math.Inf(1), math.Inf(1))) {
		t.Error("Invalid minimum")
	}
	if !bb.Maximum.Eq(NewPoint(math.Inf(-1), math.Inf(-1), math.Inf(-1))) {
		t.Error("Invalid maximum")
	}
}

func TestNewBoundingBoxWithValues(t *testing.T) {
	bb := NewBoundingBoxWithValues(NewPoint(-1, -2, -3), NewPoint(3, 2, 1))

	if !bb.Minimum.Eq(NewPoint(-1, -2, -3)) {
		t.Error("Invalid minimum")
	}
	if !bb.Maximum.Eq(NewPoint(3, 2, 1)) {
		t.Error("Invalid maximum")
	}
}

func TestAddPointsToBoundingBox(t *testing.T) {
	bb := NewBoundingBox()
	p1 := NewPoint(-5, 2, 0)
	p2 := NewPoint(7, 0, -3)

	bb.Add(p1)
	bb.Add(p2)

	if !bb.Minimum.Eq(NewPoint(-5, 0, -3)) {
		t.Error("Invalid minimum")
	}
	if !bb.Maximum.Eq(NewPoint(7, 2, 0)) {
		t.Error("Invalid maximum")
	}
}

func TestSplitBoundingBox(t *testing.T) {
	box := NewBoundingBoxWithValues(NewPoint(-1, -4, -5), NewPoint(9, 6, 5))

	left, right := SplitBoundingBox(box)

	if !left.Minimum.Eq(NewPoint(-1, -4, -5)) {
		t.Errorf("Got %v, want %v", left.Minimum, NewPoint(-1, -4, -5))
	}
	if !left.Maximum.Eq(NewPoint(4, 6, 5)) {
		t.Errorf("Got %v, want %v", left.Maximum, NewPoint(4, 6, 5))
	}
	if !right.Minimum.Eq(NewPoint(4, -4, -5)) {
		t.Errorf("Got %v, want %v", right.Minimum, NewPoint(4, -4, -5))
	}
	if !right.Maximum.Eq(NewPoint(9, 6, 5)) {
		t.Errorf("Got %v, want %v", right.Maximum, NewPoint(9, 6, 5))
	}
}

func TestSplitXWideBox(t *testing.T) {
	box := NewBoundingBoxWithValues(NewPoint(-1, -2, -3), NewPoint(9, 5.5, 3))

	left, right := SplitBoundingBox(box)

	if !left.Minimum.Eq(NewPoint(-1, -2, -3)) {
		t.Errorf("Got %v, want %v", left.Minimum, NewPoint(-1, -2, -3))
	}
	if !left.Maximum.Eq(NewPoint(4, 5.5, 3)) {
		t.Errorf("Got %v, want %v", left.Maximum, NewPoint(4, 5.5, 3))
	}
	if !right.Minimum.Eq(NewPoint(4, -2, -3)) {
		t.Errorf("Got %v, want %v", right.Minimum, NewPoint(4, -2, -3))
	}
	if !right.Maximum.Eq(NewPoint(9, 5.5, 3)) {
		t.Errorf("Got %v, want %v", right.Maximum, NewPoint(9, 5.5, 3))
	}
}

func TestSplitYWideBox(t *testing.T) {
	box := NewBoundingBoxWithValues(NewPoint(-1, -2, -3), NewPoint(5, 8, 3))

	left, right := SplitBoundingBox(box)

	if !left.Minimum.Eq(NewPoint(-1, -2, -3)) {
		t.Errorf("Got %v, want %v", left.Minimum, NewPoint(-1, -2, -3))
	}
	if !left.Maximum.Eq(NewPoint(5, 3, 3)) {
		t.Errorf("Got %v, want %v", left.Maximum, NewPoint(5, 3, 3))
	}
	if !right.Minimum.Eq(NewPoint(-1, 3, -3)) {
		t.Errorf("Got %v, want %v", right.Minimum, NewPoint(-1, 3, -3))
	}
	if !right.Maximum.Eq(NewPoint(5, 8, 3)) {
		t.Errorf("Got %v, want %v", right.Maximum, NewPoint(5, 8, 3))
	}
}

func TestSplitZWideBox(t *testing.T) {
	box := NewBoundingBoxWithValues(NewPoint(-1, -2, -3), NewPoint(5, 3, 7))

	left, right := SplitBoundingBox(box)

	if !left.Minimum.Eq(NewPoint(-1, -2, -3)) {
		t.Errorf("Got %v, want %v", left.Minimum, NewPoint(-1, -2, -3))
	}
	if !left.Maximum.Eq(NewPoint(5, 3, 2)) {
		t.Errorf("Got %v, want %v", left.Maximum, NewPoint(5, 3, 2))
	}
	if !right.Minimum.Eq(NewPoint(-1, -2, 2)) {
		t.Errorf("Got %v, want %v", right.Minimum, NewPoint(-1, -2, 2))
	}
	if !right.Maximum.Eq(NewPoint(5, 3, 7)) {
		t.Errorf("Got %v, want %v", right.Maximum, NewPoint(5, 3, 7))
	}
}

func TestPartitioningAGroupsChildren(t *testing.T) {
	s1 := NewSphere()
	s1.SetTransform(NewTranslation(-2, 0, 0))

	s2 := NewSphere()
	s2.SetTransform(NewTranslation(2, 0, 0))

	s3 := NewSphere()

	g := NewGroup()
	g.AddChild(s1)
	g.AddChild(s2)
	g.AddChild(s3)

	left, right := PartitionChildren(g)

	if len(g.Items) != 1 && g.Items[0] != s3 {
		t.Error("Invalid children left in g")
	}

	if len(left) != 1 && left[0] != s1 {
		t.Error("Invalid left")
	}

	if len(right) != 1 && right[0] != s2 {
		t.Error("Invalid right")
	}
}

func TestGreatingSubgroupOfChildren(t *testing.T) {
	s1 := NewSphere()
	s2 := NewSphere()

	g := NewGroup()

	MakeSubGroup(g, []Intersectable{s1, s2})

	if len(g.Items) != 1 {
		t.Error("Invalid child length")
	}

	if g.Items[0].(*Group).Items[0] != s1 && g.Items[0].(*Group).Items[1] != s2 {
		t.Error("Invalid subgroup")
	}
}

func TestDivideGroup(t *testing.T) {
	s1 := NewSphere()
	s1.SetTransform(NewTranslation(-2, -2, 0))

	s2 := NewSphere()
	s2.SetTransform(NewTranslation(-2, 2, 0))

	s3 := NewSphere()
	s3.SetTransform(NewScaling(4, 4, 4))

	g := NewGroup()
	g.AddChild(s1)
	g.AddChild(s2)
	g.AddChild(s3)

	g.Divide(1)

	if g.Items[0] != s3 {
		t.Error("Invalid first item")
	}

	subgroup := g.Items[1].(*Group)

	if len(subgroup.Items) != 2 {
		t.Error("First subgroup wrong count")
	}

	subgroup0 := subgroup.Items[0].(*Group)
	subgroup1 := subgroup.Items[1].(*Group)

	if len(subgroup0.Items) != 1 && subgroup0.Items[0] != s1 {
		t.Error("Invalid subgroup 0")
	}
	if len(subgroup1.Items) != 1 && subgroup0.Items[0] != s2 {
		t.Error("Invalid subgroup 1")
	}
}

func TestDivideCSGShapes(t *testing.T) {
	s1 := NewSphere()
	s1.SetTransform(NewTranslation(-1.5, 0, 0))

	s2 := NewSphere()
	s2.SetTransform(NewTranslation(1.5, 0, 0))

	left := NewGroup()
	left.AddChild(s1)
	left.AddChild(s2)

	s3 := NewSphere()
	s3.SetTransform(NewTranslation(0, 0, -1.5))

	s4 := NewSphere()
	s4.SetTransform(NewTranslation(0, 0, 1.5))

	right := NewGroup()
	right.AddChild(s3)
	right.AddChild(s4)

	shape := NewCSG(Difference, left, right)

	shape.Divide(1)

	if left.Items[0].(*Group).Items[0] != s1 {
		t.Error("Invalid left subgroup 0")
	}
	if left.Items[1].(*Group).Items[0] != s2 {
		t.Error("Invalid left subgroup 1")
	}
	if right.Items[0].(*Group).Items[0] != s3 {
		t.Error("Invalid right subgroup 0")
	}
	if right.Items[1].(*Group).Items[0] != s4 {
		t.Error("Invalid left subgroup 1")
	}
}

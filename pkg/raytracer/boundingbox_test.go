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

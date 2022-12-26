package raytracer

import (
	"testing"
)

func TestNewPointLight(t *testing.T) {
	pl := NewPointLight(NewPoint(0, 0, 0), NewColor(1, 1, 1))

	if !pl.Intensity.Eq(NewColor(1, 1, 1)) || !pl.Position.Eq(NewPoint(0, 0, 0)) {
		t.Error("PointLight initialized incorrectly")
	}
}

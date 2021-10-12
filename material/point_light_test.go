package material_test

import (
	"testing"

	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
	m "github.com/fredrikln/the-ray-tracer-challenge-go/material"
)

func TestNewPointLight(t *testing.T) {
	pl := m.NewPointLight(g.NewPoint(0, 0, 0), m.NewColor(1, 1, 1))

	if !pl.Intensity.Eq(m.NewColor(1, 1, 1)) || !pl.Position.Eq(g.NewPoint(0, 0, 0)) {
		t.Error("PointLight initialized incorrectly")
	}
}

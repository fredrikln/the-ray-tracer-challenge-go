package material

import (
	"testing"

	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
)

func TestNewPointLight(t *testing.T) {
	pl := NewPointLight(g.NewPoint(0, 0, 0), NewColor(1, 1, 1))

	if !pl.Intensity.Eq(NewColor(1, 1, 1)) || !pl.Position.Eq(g.NewPoint(0, 0, 0)) {
		t.Error("PointLight initialized incorrectly")
	}
}

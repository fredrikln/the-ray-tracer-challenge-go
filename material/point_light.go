package material

import (
	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
)

type PointLight struct {
	Intensity Color
	Position  g.Point
}

func NewPointLight(position g.Point, intensity Color) *PointLight {
	return &PointLight{
		intensity,
		position,
	}
}

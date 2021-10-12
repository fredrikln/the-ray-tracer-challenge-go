package material_test

import (
	"math"
	"testing"

	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
	m "github.com/fredrikln/the-ray-tracer-challenge-go/material"
)

func TestMaterial(t *testing.T) {
	ma := m.NewMaterial()

	if !ma.Color.Eq(m.NewColor(1, 1, 1)) || ma.Ambient != 0.1 || ma.Diffuse != 0.9 || ma.Specular != 0.9 || ma.Shininess != 200 {
		t.Error("Material not initialized correctly")
	}
}

func TestLighting(t *testing.T) {
	mat := m.NewMaterial()
	p := g.NewPoint(0, 0, 0)

	tests := []struct {
		name    string
		eyev    g.Vec
		normalv g.Vec
		light   m.PointLight
		want    m.Color
	}{
		{
			"Test 1",
			g.NewVec(0, 0, -1),
			g.NewVec(0, 0, -1),
			m.NewPointLight(g.NewPoint(0, 0, -10), m.NewColor(1, 1, 1)),
			m.NewColor(1.9, 1.9, 1.9),
		},
		{
			"Test 2",
			g.NewVec(0, math.Sqrt(2)/2, math.Sqrt(2)/2),
			g.NewVec(0, 0, -1),
			m.NewPointLight(g.NewPoint(0, 0, -10), m.NewColor(1, 1, 1)),
			m.NewColor(1.0, 1.0, 1.0),
		},
		{
			"Test 3",
			g.NewVec(0, 0, -1),
			g.NewVec(0, 0, -1),
			m.NewPointLight(g.NewPoint(0, 10, -10), m.NewColor(1, 1, 1)),
			m.NewColor(0.7364, 0.7364, 0.7364),
		},
		{
			"Test 4",
			g.NewVec(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2),
			g.NewVec(0, 0, -1),
			m.NewPointLight(g.NewPoint(0, 10, -10), m.NewColor(1, 1, 1)),
			m.NewColor(1.6364, 1.6364, 1.6364),
		},
		{
			"Test 5",
			g.NewVec(0, 0, -1),
			g.NewVec(0, 0, -1),
			m.NewPointLight(g.NewPoint(0, 0, 10), m.NewColor(1, 1, 1)),
			m.NewColor(0.1, 0.1, 0.1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mat.Lighting(tt.light, p, tt.eyev, tt.normalv)

			if !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

package raytracer

import (
	"math"
	"testing"
)

func TestMaterial(t *testing.T) {
	ma := NewMaterial()

	if !ma.Color.Eq(NewColor(1, 1, 1)) || ma.Ambient != 0.1 || ma.Diffuse != 0.9 || ma.Specular != 0.9 || ma.Shininess != 200 {
		t.Error("Material not initialized correctly")
	}
}

func TestLighting(t *testing.T) {
	mat := NewMaterial()
	p := NewPoint(0, 0, 0)

	tests := []struct {
		name    string
		eyev    Vec
		normalv Vec
		light   *PointLight
		want    Color
	}{
		{
			"Test 1",
			NewVec(0, 0, -1),
			NewVec(0, 0, -1),
			NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1)),
			NewColor(1.9, 1.9, 1.9),
		},
		{
			"Test 2",
			NewVec(0, math.Sqrt(2)/2, math.Sqrt(2)/2),
			NewVec(0, 0, -1),
			NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1)),
			NewColor(1.0, 1.0, 1.0),
		},
		{
			"Test 3",
			NewVec(0, 0, -1),
			NewVec(0, 0, -1),
			NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1)),
			NewColor(0.7364, 0.7364, 0.7364),
		},
		{
			"Test 4",
			NewVec(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2),
			NewVec(0, 0, -1),
			NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1)),
			NewColor(1.6364, 1.6364, 1.6364),
		},
		{
			"Test 5",
			NewVec(0, 0, -1),
			NewVec(0, 0, -1),
			NewPointLight(NewPoint(0, 0, 10), NewColor(1, 1, 1)),
			NewColor(0.1, 0.1, 0.1),
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

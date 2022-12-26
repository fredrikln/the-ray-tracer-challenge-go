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
			got := mat.Lighting(NewSphere(), tt.light, p, tt.eyev, tt.normalv, false)

			if !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLightingInShadow(t *testing.T) {
	eyev := NewVec(0, 0, -1)
	normalv := NewVec(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	inShadow := true

	m := NewMaterial()
	position := NewPoint(0, 0, 0)
	got := m.Lighting(NewSphere(), light, position, eyev, normalv, inShadow)
	want := NewColor(0.1, 0.1, 0.1)

	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestIsShadowed(t *testing.T) {
	testCases := []struct {
		desc  string
		world *World
		point Point
		want  bool
	}{
		{
			desc:  "Nothing collinear",
			world: NewDefaultWorld(),
			point: NewPoint(0, 10, 0),
			want:  false,
		},
		{
			desc:  "Object between",
			world: NewDefaultWorld(),
			point: NewPoint(10, -10, 10),
			want:  true,
		},
		{
			desc:  "Object behind light",
			world: NewDefaultWorld(),
			point: NewPoint(-20, 20, -20),
			want:  false,
		},
		{
			desc:  "Point between light and object",
			world: NewDefaultWorld(),
			point: NewPoint(-2, 2, -2),
			want:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := tC.world.IsShadowed(*tC.world.Lights[0], tC.point)
			want := tC.want

			if got != want {
				t.Errorf("Got %v, want %v", got, want)
			}
		})
	}
}

func TestLightingWithPatternApplied(t *testing.T) {
	p := NewStripePattern(white, black)
	m := NewMaterial().SetAmbient(1).SetDiffuse(0).SetSpecular(0).SetPattern(p)
	eyev := NewVec(0, 0, -1)
	normalv := NewVec(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))

	c1 := m.Lighting(NewSphere(), light, NewPoint(0.9, 0, 0), eyev, normalv, false)
	c2 := m.Lighting(NewSphere(), light, NewPoint(1.1, 0, 0), eyev, normalv, false)

	if !c1.Eq(NewColor(1, 1, 1)) {
		t.Errorf("Invalid c1, got %v, want %v", c1, NewColor(1, 1, 1))
	}

	if !c2.Eq(NewColor(0, 0, 0)) {
		t.Errorf("Invalid c1 got %v, want %v", c2, NewColor(0, 0, 0))
	}
}

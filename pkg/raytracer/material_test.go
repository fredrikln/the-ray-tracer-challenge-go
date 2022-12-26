package raytracer

import (
	"math"
	"strconv"
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

func TestDefaultMaterialReflectivity(t *testing.T) {
	m := NewMaterial()

	if m.Reflectivity != 0.0 {
		t.Error("Invalid reflectivity in default material")
	}
}

func TestTransparencyAndRefractiveIndex(t *testing.T) {
	m := NewMaterial()

	if m.Transparency != 0.0 {
		t.Error("Invalid transparency")
	}

	if m.RefractiveIndex != 1 {
		t.Error("Invalid refractive index")
	}
}

func TestNewGlassSphere(t *testing.T) {
	s := NewGlassSphere()

	if !s.Transform.Eq(NewIdentityMatrix()) {
		t.Error("Invalid default transform")
	}

	if s.Material.Transparency != 1.0 {
		t.Error("Invalid transparency")
	}

	if s.Material.RefractiveIndex != 1.5 {
		t.Error("Invalid refractive index")
	}
}

func TestFindingN1AndN2(t *testing.T) {
	a := NewGlassSphere().SetTransform(NewScaling(2, 2, 2))
	a.GetMaterial().SetRefractiveIndex(1.5)

	b := NewGlassSphere().SetTransform(NewTranslation(0, 0, -0.25))
	b.GetMaterial().SetRefractiveIndex(2.0)

	c := NewGlassSphere().SetTransform(NewTranslation(0, 0, 0.25))
	c.GetMaterial().SetRefractiveIndex(2.5)

	r := NewRay(NewPoint(0, 0, -4), NewVec(0, 0, 1))

	testCases := []struct {
		desc         string
		intersection Intersection
		n1           float64
		n2           float64
	}{
		{
			desc: "0",
			n1:   1.0,
			n2:   1.5,
		},
		{
			desc: "1",
			n1:   1.5,
			n2:   2.0,
		},
		{
			desc: "2",
			n1:   2.0,
			n2:   2.5,
		},
		{
			desc: "3",
			n1:   2.5,
			n2:   2.5,
		},
		{
			desc: "4",
			n1:   2.5,
			n2:   1.5,
		},
		{
			desc: "5",
			n1:   1.5,
			n2:   1.0,
		},
	}

	xs := []Intersection{
		NewIntersection(2, a),
		NewIntersection(2.75, b),
		NewIntersection(3.25, c),
		NewIntersection(4.75, b),
		NewIntersection(5.25, c),
		NewIntersection(6, a),
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			index, _ := strconv.Atoi(tC.desc)
			comps := PrepareComputationsWithHit(xs[index], r, xs)

			if comps.N1 != tC.n1 {
				t.Errorf("Invalid n1, got %v, want %v", comps.N1, tC.n1)
			}
			if comps.N2 != tC.n2 {
				t.Errorf("Invalid n2, got %v, want %v", comps.N2, tC.n2)
			}
		})
	}
}

package material

import (
	"math"

	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
)

type Material struct {
	Color     Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func NewMaterial() Material {
	return Material{
		NewColor(1, 1, 1),
		0.1,
		0.9,
		0.9,
		200,
	}
}

func (m Material) Lighting(light PointLight, point g.Point, eyev g.Vec, normalv g.Vec) Color {
	effectiveColor := m.Color.Mul(light.Intensity)

	lightv := light.Position.Sub(point).Norm()

	ambient := effectiveColor.MulFloat(m.Ambient)

	lightDotNormal := lightv.Dot(normalv)

	diffuse := NewColor(0, 0, 0)
	specular := NewColor(0, 0, 0)

	if lightDotNormal >= 0 {
		diffuse = effectiveColor.MulFloat(m.Diffuse).MulFloat(lightDotNormal)

		reflectv := lightv.Mul(-1).Reflect(normalv)
		reflectDotEye := reflectv.Dot(eyev)

		if reflectDotEye <= 0 {
			specular = NewColor(0, 0, 0)
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.Intensity.MulFloat(m.Specular).MulFloat(factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}

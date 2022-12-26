package raytracer

import (
	"math"
)

type Material struct {
	Color           Color
	Ambient         float64
	Diffuse         float64
	Specular        float64
	Shininess       float64
	Reflectivity    float64
	Transparency    float64
	RefractiveIndex float64
	Pattern         *Pattern
}

func NewMaterial() *Material {
	return &Material{
		Color:           NewColor(1, 1, 1),
		Ambient:         0.1,
		Diffuse:         0.9,
		Specular:        0.9,
		Shininess:       200,
		Reflectivity:    0.0,
		Transparency:    0.0,
		RefractiveIndex: 1.0,
		Pattern:         nil,
	}
}

func (m *Material) SetColor(c Color) *Material {
	m.Color = c

	return m
}
func (m *Material) SetDiffuse(d float64) *Material {
	m.Diffuse = d

	return m
}
func (m *Material) SetSpecular(s float64) *Material {
	m.Specular = s

	return m
}
func (m *Material) SetShininess(s float64) *Material {
	m.Shininess = s

	return m
}
func (m *Material) SetAmbient(a float64) *Material {
	m.Ambient = a

	return m
}
func (m *Material) SetReflective(r float64) *Material {
	m.Reflectivity = r

	return m
}
func (m *Material) SetTransparency(t float64) *Material {
	m.Transparency = t

	return m
}
func (m *Material) SetRefractiveIndex(ri float64) *Material {
	m.RefractiveIndex = ri

	return m
}
func (m *Material) SetPattern(p Pattern) *Material {
	m.Pattern = &p

	return m
}

func (m *Material) Lighting(object Intersectable, light Light, point Point, eyev Vec, normalv Vec, inShadow bool) Color {
	var color Color

	if m.Pattern != nil {
		color = (*m.Pattern).ColorAtObject(object, point)
	} else {
		color = m.Color
	}

	effectiveColor := color.Mul(light.GetIntensity())

	lightv := light.GetPosition().Sub(point).Norm()

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
			specular = light.GetIntensity().MulFloat(m.Specular).MulFloat(factor)
		}
	}

	if inShadow {
		diffuse = diffuse.MulFloat(0)
		specular = specular.MulFloat(0)
	}

	return ambient.Add(diffuse).Add(specular)
}

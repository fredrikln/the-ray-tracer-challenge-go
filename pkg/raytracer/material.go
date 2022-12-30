package raytracer

import (
	"math"
	"math/rand"
)

type Scatters interface {
	Emit() Color
	Scatter(rayIn *Ray, comps *Computations, attenuation *Color, scattered *Ray, source *rand.Rand) bool
}

type Diffuse struct {
	Albedo Color
}

func NewDiffuse(color Color) *Diffuse {
	return &Diffuse{
		Albedo: color,
	}
}

func (d *Diffuse) Emit() Color {
	return NewColor(0, 0, 0)
}

func (d *Diffuse) Scatter(rayIn *Ray, comps *Computations, attenuation *Color, scattered *Ray, source *rand.Rand) bool {
	scatterDirection := comps.Normalv.Add(RandomUnitVector(source))

	if scatterDirection.NearZero() {
		scatterDirection = comps.Normalv
	}

	*scattered = NewRay(comps.OverPoint, scatterDirection)
	*attenuation = d.Albedo

	return true
}

type Metal struct {
	Albedo    Color
	Fuzziness float64
}

func NewMetal(color Color, fuzziness float64) *Metal {
	return &Metal{
		Albedo:    color,
		Fuzziness: fuzziness,
	}
}

func (m *Metal) Emit() Color {
	return NewColor(0, 0, 0)
}

func (m *Metal) Scatter(rayIn *Ray, comps *Computations, attenuation *Color, scattered *Ray, source *rand.Rand) bool {
	scatterDirection := comps.Reflectv.Add(RandomInUnitSphere(source).Mul(m.Fuzziness))

	*scattered = NewRay(comps.OverPoint, scatterDirection)
	*attenuation = m.Albedo

	return true
}

type Dielectric struct {
	IndexOfRefraction float64
}

func NewDielectric(indexOfRefraction float64) *Dielectric {
	return &Dielectric{
		IndexOfRefraction: indexOfRefraction,
	}
}

func (d *Dielectric) Emit() Color {
	return NewColor(0, 0, 0)
}

func (d *Dielectric) Scatter(rayIn *Ray, comps *Computations, attenuation *Color, scattered *Ray, source *rand.Rand) bool {
	*attenuation = NewColor(1, 1, 1)

	var refractionRatio float64
	if !comps.Inside {
		refractionRatio = 1.0 / d.IndexOfRefraction
	} else {
		refractionRatio = d.IndexOfRefraction
	}

	unitDirection := rayIn.Direction.Norm()

	cosTheta := math.Min(unitDirection.Mul(-1).Dot(comps.Normalv), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)

	cannotRefract := refractionRatio*sinTheta > 1.0

	var direction Vec

	if cannotRefract || Reflectance(cosTheta, refractionRatio) > source.Float64() {
		direction = Reflect(unitDirection, comps.Normalv)
	} else {
		direction = Refract(unitDirection, comps.Normalv, refractionRatio)
	}

	*scattered = NewRay(comps.UnderPoint, direction)

	return true
}

type Emissive struct {
	Emission Color
}

func NewEmissive(color Color) *Emissive {
	return &Emissive{
		Emission: color,
	}
}

func (e *Emissive) Emit() Color {
	return e.Emission
}

func (e *Emissive) Scatter(rayIn *Ray, comps *Computations, attenuation *Color, scattered *Ray, source *rand.Rand) bool {
	return false
}

func Reflectance(cosine, refIdx float64) float64 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	oneMinCosine := (1 - cosine)
	oneMinCosine5 := oneMinCosine * oneMinCosine * oneMinCosine * oneMinCosine * oneMinCosine

	return r0 + (1-r0)*oneMinCosine5

}

func Reflect(v Vec, n Vec) Vec {
	return v.Sub(n.Mul(v.Dot(n)).Mul(2))
}

func Refract(uv Vec, n Vec, etaiOverEtaT float64) Vec {
	cosTheta := math.Min(uv.Mul(-1).Dot(n), 1.0)
	rOutPerp := (uv.Add(n.Mul(cosTheta))).Mul(etaiOverEtaT)
	rOutParallel := n.Mul(-math.Sqrt(math.Abs(1.0 - (rOutPerp.Mag() * rOutPerp.Mag()))))

	return rOutPerp.Add(rOutParallel)
}

func RandFloat(source *rand.Rand, min, max float64) float64 {
	return min + (max-min)*source.Float64()
}

func RandomUnitVector(source *rand.Rand) Vec {
	for {
		v := NewVec(RandFloat(source, -1, 1), RandFloat(source, -1, 1), RandFloat(source, -1, 1))

		if v.LengthSquared() >= 1 {
			continue
		}

		return v
	}
}

func RandomInUnitSphere(source *rand.Rand) Vec {
	return RandomUnitVector(source).Norm()
}

func RandomInHemisphere(source *rand.Rand, n Vec) Vec {
	inUnitSphere := RandomInUnitSphere(source)

	if inUnitSphere.Dot(n) > 0.0 {
		return inUnitSphere
	}

	return inUnitSphere.Mul(-1)

}

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

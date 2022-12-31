package raytracer

import (
	"math"
	"math/rand"
)

type Scatters interface {
	Emit() Color
	Scatter(rayIn *Ray, comps *Computations, attenuation *Color, scattered *Ray, source *rand.Rand) bool
}

// Diffuse
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

// Dielectric
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

// Emissive
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

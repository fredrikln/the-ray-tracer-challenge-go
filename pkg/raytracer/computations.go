package raytracer

import "math"

type Computations struct {
	Time       float64
	Object     *Intersectable
	Point      Point
	Eyev       Vec
	Normalv    Vec
	Reflectv   Vec
	Inside     bool
	OverPoint  Point
	UnderPoint Point
	N1         float64
	N2         float64
}

func PrepareComputationsWithHit(i Intersection, r Ray, xs []Intersection) Computations {
	comps := PrepareComputations(i, r)

	var n1, n2 float64 = 1.0, 1.0

	containers := make([]*Intersectable, 0)
	for _, item := range xs {
		if i == item {
			if len(containers) == 0 {
				n1 = 1.0
			} else {
				n1 = (*containers[len(containers)-1]).GetMaterial().RefractiveIndex
			}
		}

		var itemIndex int = -1
		for index := 0; index < len(containers); index++ {
			if *containers[index] == *item.Object {
				itemIndex = index
			}
		}

		if itemIndex != -1 {
			containers = append(containers[:itemIndex], containers[itemIndex+1:]...)
		} else {
			containers = append(containers, item.Object)
		}
		if i == item {
			if len(containers) == 0 {
				n2 = 1.0
			} else {
				n2 = (*containers[len(containers)-1]).GetMaterial().RefractiveIndex
			}

			break
		}

	}

	comps.N1 = n1
	comps.N2 = n2

	return comps
}

func PrepareComputations(i Intersection, r Ray) Computations {
	p := r.Position(i.Time)

	eyev := r.Direction.Neg()
	normalv := (*i.Object).NormalAt(p).Norm()
	inside := normalv.Dot(eyev) < 0

	if inside {
		normalv = normalv.Neg()
	}

	reflectv := r.Direction.Reflect(normalv)

	return Computations{
		Time:       i.Time,
		Object:     i.Object,
		Point:      p,
		Eyev:       eyev,
		Normalv:    normalv,
		Reflectv:   reflectv,
		Inside:     inside,
		OverPoint:  p.AddVec(normalv.Mul((1e-5) / 2)),
		UnderPoint: p.AddVec(normalv.Mul(-(1e-5) / 2)),
		N1:         1.0,
		N2:         1.0,
	}
}

func Schlick(comps Computations) float64 {
	cos := comps.Eyev.Dot(comps.Normalv)

	if comps.N1 > comps.N2 {
		n := comps.N1 / comps.N2

		sin2T := n * n * (1 - cos*cos)

		if sin2T > 1.0 {
			return 1.0
		}

		cosT := math.Sqrt(1 - sin2T)

		cos = cosT
	}

	d := (comps.N1 - comps.N2) / (comps.N1 + comps.N2)
	d2 := d * d

	r0 := d2

	e := 1 - cos
	e5 := e * e * e * e * e

	return r0 + (1-r0)*e5
}

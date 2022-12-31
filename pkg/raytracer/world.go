package raytracer

import (
	"math/rand"
	"sort"
	"time"
)

type IntersectonSorter []Intersection

func (is IntersectonSorter) Len() int {
	return len(is)
}
func (is IntersectonSorter) Less(i, j int) bool {
	return is[i].Time < is[j].Time
}
func (is IntersectonSorter) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

type World struct {
	Objects    []*Intersectable
	Lights     []*Light
	Background *Color
	Source     *rand.Rand
}

func NewWorld() *World {
	return &World{
		Source: rand.New(rand.NewSource(time.Now().Unix() + rand.Int63())),
	}
}

func NewDefaultWorld() *World {
	s1 := NewSphere()
	m1 := NewDiffuse(NewColor(0.8, 1.0, 0.6))
	s1.SetNewMaterial(m1)

	s2 := NewSphere().SetTransform(NewScaling(0.5, 0.5, 0.5))

	w := NewWorld()
	w.AddObject(s1)
	w.AddObject(s2)

	c := NewColor(0.1, 0.1, 0.1)
	w.Background = &c

	return w
}

func (w *World) AddLight(l Light) *World {
	w.Lights = append(w.Lights, &l)

	return w
}

func (w *World) AddObject(i Intersectable) *World {
	w.Objects = append(w.Objects, &i)

	return w
}

func (w *World) Intersect(r Ray) []Intersection {
	xs := make([]Intersection, 0)

	for _, object := range w.Objects {
		objectXs := (*object).Intersect(r)
		xs = append(xs, objectXs...)
	}

	sort.Sort(IntersectonSorter(xs))

	return xs
}

// func (w *World) ShadeHit(comps Computations, remaining int) Color {
// 	color := NewColor(0, 0, 0)

// 	for _, light := range w.Lights {
// 		inShadow := w.IsShadowed(*light, comps.OverPoint)

// 		c2 := (*comps.Object).GetMaterial().Lighting(*comps.Object, *light, comps.OverPoint, comps.Eyev, comps.Normalv, inShadow)
// 		color = color.Add(c2)
// 	}

// 	reflected := w.ReflectedColor(comps, remaining)
// 	refracted := w.RefractedColor(comps, remaining)

// 	material := (*comps.Object).GetMaterial()

// 	if material.Reflectivity > 0 && material.Transparency > 0 {
// 		reflectance := Schlick(comps)

// 		return color.Add(reflected.MulFloat(reflectance)).Add(refracted.MulFloat(1 - reflectance))
// 	}

// 	return color.Add(reflected).Add(refracted)
// }

var colorBlack Color = Color{0, 0, 0}

func (w *World) ColorAt(r *Ray, remaining int) Color {
	if remaining <= 0 {
		return colorBlack
	}

	xs := w.Intersect(*r)
	hit, didHit := GetHit(xs)

	if !didHit {
		if w.Background != nil {
			return *w.Background
		}

		return colorBlack
	}

	var scattered Ray
	var attenuation Color

	comps := PrepareComputationsWithHit(hit, *r, xs)
	object := *comps.Object

	emit := object.GetNewMaterial().Emit()

	if !object.GetNewMaterial().Scatter(r, comps, &attenuation, &scattered, w.Source) {
		return emit
	}

	return emit.Add(attenuation.Mul(w.ColorAt(&scattered, remaining-1)))

	// unitDirection := r.Direction.Norm()
	// t := 0.5 * (unitDirection.Y + 1.0)

	// return (NewColor(1, 1, 1).MulFloat(1 - t)).Add(NewColor(0.5, 0.7, 1.0).MulFloat(t))

	// if !didHit {
	// 	return NewColor(0, 0, 0)
	// }

	// comps := PrepareComputationsWithHit(hit, r, xs)

	// return w.ShadeHit(comps, remaining)
}

func (w *World) IsShadowed(l Light, p Point) bool {
	v := l.GetPosition().Sub(p)
	distance := v.Mag()
	direction := v.Norm()

	r := NewRay(p, direction)
	xs := w.Intersect(r)

	hit, didHit := GetHit(xs)

	if didHit && hit.Time < distance {
		return true
	}

	return false
}

// func (w *World) ReflectedColor(comps Computations, remaining int) Color {
// 	if remaining <= 0 {
// 		return NewColor(0, 0, 0)
// 	}

// 	if (*comps.Object).GetMaterial().Reflectivity == 0 {
// 		return Color{0, 0, 0}
// 	}

// 	reflectRay := NewRay(comps.OverPoint, comps.Reflectv)
// 	color := w.ColorAt(reflectRay, remaining-1)

// 	return color.MulFloat((*comps.Object).GetMaterial().Reflectivity)
// }

// func (w *World) RefractedColor(comps Computations, remaining int) Color {
// 	if remaining <= 0 {
// 		return NewColor(0, 0, 0)
// 	}

// 	if (*comps.Object).GetMaterial().Transparency == 0 {
// 		return NewColor(0, 0, 0)
// 	}

// 	nRatio := comps.N1 / comps.N2
// 	cosI := comps.Eyev.Dot(comps.Normalv)
// 	sin2T := nRatio * nRatio * (1 - cosI*cosI)

// 	if sin2T > 1 {
// 		return NewColor(0, 0, 0)
// 	}

// 	cosT := math.Sqrt(1.0 - sin2T)

// 	direction := comps.Normalv.Mul(nRatio*cosI - cosT).Sub(comps.Eyev.Mul(nRatio))

// 	refractRay := NewRay(comps.UnderPoint, direction)

// 	color := w.ColorAt(refractRay, remaining-1).MulFloat((*comps.Object).GetMaterial().Transparency)

// 	return color
// }

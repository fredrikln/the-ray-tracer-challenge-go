package raytracer

import (
	"sort"
)

type World struct {
	Objects []*Intersectable
	Lights  []*Light
}

func NewWorld() *World {
	return &World{}
}

func NewDefaultWorld() *World {
	s1 := NewSphere()
	m1 := NewMaterial().SetColor(NewColor(0.8, 1.0, 0.6)).SetDiffuse(0.7).SetSpecular(0.2)
	s1.SetMaterial(m1)

	s2 := NewSphere().SetTransform(NewScaling(0.5, 0.5, 0.5))

	l1 := NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1))

	w := NewWorld()
	w.AddObject(s1)
	w.AddObject(s2)

	w.AddLight(l1)

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
	xs := make([]Intersection, 0, len(w.Objects)*2)

	for _, object := range w.Objects {
		objectXs := (*object).Intersect(r)
		xs = append(xs, objectXs...)
	}

	sort.Slice(xs, func(i, j int) bool {
		return xs[i].Time < xs[j].Time
	})

	return xs
}

func (w *World) ShadeHit(comps Computations, remaining int) Color {
	color := NewColor(0, 0, 0)

	for _, light := range w.Lights {
		inShadow := w.IsShadowed(*light, comps.OverPoint)

		c2 := (*comps.Object).GetMaterial().Lighting(*comps.Object, *light, comps.OverPoint, comps.Eyev, comps.Normalv, inShadow)
		color = color.Add(c2)
	}

	reflected := w.ReflectedColor(comps, remaining)

	return color.Add(reflected)
}

func (w *World) ColorAt(r Ray, remaining int) Color {
	xs := w.Intersect(r)
	hit, didHit := GetHit(xs)

	if !didHit {
		return NewColor(0, 0, 0)
	}

	comps := PrepareComputations(hit, r)

	return w.ShadeHit(comps, remaining)
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

func (w *World) ReflectedColor(comps Computations, remaining int) Color {
	if remaining <= 0 {
		return NewColor(0, 0, 0)
	}

	if (*comps.Object).GetMaterial().Reflectivity == 0 {
		return Color{0, 0, 0}
	}

	reflectRay := NewRay(comps.OverPoint, comps.Reflectv)
	color := w.ColorAt(reflectRay, remaining-1)

	return color.MulFloat((*comps.Object).GetMaterial().Reflectivity)
}

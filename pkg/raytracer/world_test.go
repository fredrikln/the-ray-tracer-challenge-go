package raytracer

import (
	"math"
	"testing"
)

func TestNewWorld(t *testing.T) {
	world := NewWorld()

	if len(world.Lights) != 0 {
		t.Errorf("Got %v lights, want %v", len(world.Lights), 0)
	}

	if len(world.Objects) != 0 {
		t.Errorf("Got %v objects, want %v", len(world.Objects), 0)
	}
}

func TestNewDefaultWorld(t *testing.T) {
	dw := NewDefaultWorld()

	if len(dw.Lights) != 1 {
		t.Error("Not right amount of lights")
	}

	if len(dw.Objects) != 2 {
		t.Error("Not right amount of objects")
	}
}

func TestIntersectWorld(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVec(0, 0, 1))

	xs := w.Intersect(r)

	if len(xs) != 4 {
		t.Errorf("Got %v, want %v", len(xs), 4)
	}

	if xs[0].Time != 4 {
		t.Errorf("xs 0, got %v, want %v", xs[0].Time, 4)
	}
	if xs[1].Time != 4.5 {
		t.Errorf("xs 0, got %v, want %v", xs[1].Time, 4.5)
	}
	if xs[2].Time != 5.5 {
		t.Errorf("xs 0, got %v, want %v", xs[2].Time, 5.5)
	}
	if xs[3].Time != 6 {
		t.Errorf("xs 0, got %v, want %v", xs[3].Time, 6)
	}
}

func TestShadeIntersection(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVec(0, 0, 1))

	s := w.Objects[0]
	i := NewIntersection(4, *s)

	comps := PrepareComputations(i, r)

	c := w.ShadeHit(comps, 4)

	want := NewColor(0.380661, 0.475826, 0.285495)

	if !c.Eq(want) {
		t.Errorf("Got %v, want %v", c, want)
	}
}

func TestShadeIntersectionFromInside(t *testing.T) {
	w := NewDefaultWorld()
	pl := NewPointLight(NewPoint(0, 0.25, 0), NewColor(1, 1, 1))
	w.Lights = []*Light{}
	w.AddLight(pl)

	r := NewRay(NewPoint(0, 0, 0), NewVec(0, 0, 1))

	s := w.Objects[1]
	i := NewIntersection(0.5, *s)

	comps := PrepareComputations(i, r)

	c := w.ShadeHit(comps, 4)

	want := NewColor(0.90498, 0.90498, 0.90498)

	if !c.Eq(want) {
		t.Errorf("Got %v, want %v", c, want)
	}
}

func TestWhenRayMisses(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVec(0, 1, 0))

	want := NewColor(0, 0, 0)
	got := w.ColorAt(r, 4)

	if !got.Eq(want) {
		t.Errorf("Did not get black, got: %v", got)
	}
}

func TestWhenRayHits(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVec(0, 0, 1))

	want := NewColor(0.380661, 0.475826, 0.285495)
	got := w.ColorAt(r, 4)

	if !got.Eq(want) {
		t.Errorf("Did not get correct color, got: %v, want %v", got, want)
	}
}

func TestUsesHitToComputeColor(t *testing.T) {
	w := NewDefaultWorld()
	outer := *w.Objects[0]
	outer.GetMaterial().SetAmbient(1)
	inner := *w.Objects[1]
	inner.GetMaterial().SetAmbient(1)

	r := NewRay(NewPoint(0, 0, 0.75), NewVec(0, 0, -1))

	want := inner.GetMaterial().Color
	got := w.ColorAt(r, 4)

	if !got.Eq(want) {
		t.Errorf("Did not get correct color, got: %v, want %v", got, want)
	}
}

func TestShadeInShadow(t *testing.T) {
	w := NewWorld()
	w.AddLight(NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1)))

	s1 := NewSphere()
	w.AddObject(s1)

	s2 := NewSphere().SetTransform(NewTranslation(0, 0, 10))
	w.AddObject(s1)

	r := NewRay(NewPoint(0, 0, 5), NewVec(0, 0, 1))
	i := NewIntersection(4, s2)

	comps := PrepareComputations(i, r)

	got := w.ShadeHit(comps, 4)
	want := NewColor(0.1, 0.1, 0.1)

	if !got.Eq(want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

func TestReflectedColor(t *testing.T) {
	w := NewDefaultWorld()
	r := NewRay(NewPoint(0, 0, 0), NewVec(0, 0, 1))
	s := *w.Objects[1]
	s.GetMaterial().SetAmbient(1)
	i := NewIntersection(1, s)

	comps := PrepareComputations(i, r)
	color := w.ReflectedColor(comps, 4)

	if !color.Eq(NewColor(0, 0, 0)) {
		t.Errorf("Invalid color, got %v, want %v", color, NewColor(0, 0, 0))
	}
}

func TestReflectedColorForAReflectiveMaterial(t *testing.T) {
	w := NewDefaultWorld()
	s := NewPlane().SetMaterial(NewMaterial().SetReflective(0.5)).SetTransform(NewTranslation(0, -1, 0))
	w.AddObject(s)

	r := NewRay(NewPoint(0, 0, -3), NewVec(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i := NewIntersection(math.Sqrt(2), s)

	comps := PrepareComputations(i, r)
	color := w.ShadeHit(comps, 4)

	want := NewColor(0.876756, 0.924339, 0.829173)

	if !color.Eq(want) {
		t.Errorf("Invalid color, got %v, want %v", color, want)
	}
}

func TestMutuallyReflectiveSurfaces(t *testing.T) {
	w := NewWorld()
	w.AddLight(NewPointLight(NewPoint(0, 0, 0), NewColor(1, 1, 1)))

	lower := NewPlane().SetMaterial(NewMaterial().SetReflective(1)).SetTransform(NewTranslation(0, -1, 0))
	upper := NewPlane().SetMaterial(NewMaterial().SetReflective(1)).SetTransform(NewTranslation(0, 1, 0))
	w.AddObject(lower)
	w.AddObject(upper)

	r := NewRay(NewPoint(0, 0, 0), NewVec(0, 1, 0))

	color := w.ColorAt(r, 4)

	if !color.Eq(NewColor(9.5, 9.5, 9.5)) {
		t.Errorf("Got %v, want %v", color, NewColor(9.5, 9.5, 9.5))
	}
}

func TestReflectedColorAtMaximumRecursiveDepth(t *testing.T) {
	w := NewDefaultWorld()

	s := NewPlane().SetMaterial(NewMaterial().SetReflective(0.5)).SetTransform(NewTranslation(0, -1, 0))
	w.AddObject(s)

	r := NewRay(NewPoint(0, 0, -3), NewVec(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	i := NewIntersection(math.Sqrt(2), s)

	comps := PrepareComputations(i, r)

	color := w.ReflectedColor(comps, 0)

	if !color.Eq(NewColor(0, 0, 0)) {
		t.Errorf("Got %v, want %v", color, NewColor(0, 0, 0))
	}
}

func TestFindRefractedColorOfOpaqueObject(t *testing.T) {
	w := NewDefaultWorld()

	s := w.Objects[0]

	r := NewRay(NewPoint(0, 0, -5), NewVec(0, 0, 1))

	xs := []Intersection{
		NewIntersection(4, *s),
		NewIntersection(6, *s),
	}

	comps := PrepareComputationsWithHit(xs[0], r, xs)

	c := w.RefractedColor(comps, 5)

	if !c.Eq(NewColor(0, 0, 0)) {
		t.Errorf("Got %v, want %v", c, NewColor(0, 0, 0))
	}
}

func TestFindRefractedColorAtMaxRecursiveDepth(t *testing.T) {
	w := NewDefaultWorld()

	s := w.Objects[0]
	(*s).GetMaterial().SetTransparency(1.0).SetRefractiveIndex(1.5)

	r := NewRay(NewPoint(0, 0, -5), NewVec(0, 0, 1))

	xs := []Intersection{
		NewIntersection(4, *s),
		NewIntersection(6, *s),
	}

	comps := PrepareComputationsWithHit(xs[0], r, xs)

	c := w.RefractedColor(comps, 0)

	if !c.Eq(NewColor(0, 0, 0)) {
		t.Errorf("Got %v, want %v", c, NewColor(0, 0, 0))
	}
}

func TestFindRefractedColorUnderTotalInternalReflection(t *testing.T) {
	w := NewDefaultWorld()
	s := w.Objects[0]

	(*s).GetMaterial().SetTransparency(1).SetRefractiveIndex(1.5)

	r := NewRay(NewPoint(0, 0, math.Sqrt(2)/2), NewVec(0, 1, 0))

	xs := []Intersection{
		NewIntersection(-math.Sqrt(2)/2, *s),
		NewIntersection(math.Sqrt(2)/2, *s),
	}

	comps := PrepareComputationsWithHit(xs[1], r, xs)

	c := w.RefractedColor(comps, 5)

	if !c.Eq(NewColor(0, 0, 0)) {
		t.Errorf("Got %v, want %v", c, NewColor(0, 0, 0))
	}
}

func TestFindRefractedColor(t *testing.T) {
	w := NewDefaultWorld()
	a := w.Objects[0]
	(*a).GetMaterial().SetPattern(NewTestPattern())

	b := w.Objects[1]
	(*b).GetMaterial().SetTransparency(1).SetRefractiveIndex(1.5)

	r := NewRay(NewPoint(0, 0, 0.1), NewVec(0, 1, 0))

	xs := []Intersection{
		NewIntersection(-0.9899, *a),
		NewIntersection(-0.4899, *b),
		NewIntersection(0.4899, *b),
		NewIntersection(0.9899, *a),
	}

	comps := PrepareComputationsWithHit(xs[2], r, xs)

	c := w.RefractedColor(comps, 5)

	if !c.Eq(NewColor(0, 0.998879, 0.047217)) {
		t.Errorf("Got %v, want %v", c, NewColor(0, 0.998879, 0.047217))
	}
}

func TestShadeHitWithATransparentMaterial(t *testing.T) {
	w := NewDefaultWorld()

	floor := NewPlane().SetTransform(NewTranslation(0, -1, 0))
	floor.GetMaterial().SetTransparency(0.5).SetRefractiveIndex(1.5)

	w.AddObject(floor)

	ball := NewSphere().SetTransform(NewTranslation(0, -3.5, -0.5))
	ball.GetMaterial().SetColor(NewColor(1, 0, 0)).SetAmbient(0.5)

	w.AddObject(ball)

	r := NewRay(NewPoint(0, 0, -3), NewVec(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))
	xs := []Intersection{NewIntersection(math.Sqrt(2), floor)}

	comps := PrepareComputationsWithHit(xs[0], r, xs)

	color := w.ShadeHit(comps, 5)

	if !color.Eq(NewColor(0.93642, 0.68642, 0.68642)) {
		t.Errorf("Got %v, want %v", color, NewColor(0.93642, 0.68642, 0.68642))
	}
}

func TestShadeHitWithAReflectiveAndTransparentMaterial(t *testing.T) {
	w := NewDefaultWorld()

	r := NewRay(NewPoint(0, 0, -3), NewVec(0, -math.Sqrt(2)/2, math.Sqrt(2)/2))

	floor := NewPlane().SetTransform(NewTranslation(0, -1, 0))
	floor.GetMaterial().SetReflective(0.5).SetTransparency(0.5).SetRefractiveIndex(1.5)

	w.AddObject(floor)

	ball := NewSphere().SetTransform(NewTranslation(0, -3.5, -0.5))
	ball.GetMaterial().SetColor(NewColor(1, 0, 0)).SetAmbient(0.5)

	w.AddObject(ball)

	xs := []Intersection{
		NewIntersection(math.Sqrt(2), floor),
	}

	comps := PrepareComputationsWithHit(xs[0], r, xs)

	color := w.ShadeHit(comps, 5)

	if !color.Eq(NewColor(0.93391, 0.69643, 0.69243)) {
		t.Errorf("Got %v, want %v", color, NewColor(0.93391, 0.69643, 0.69243))
	}
}

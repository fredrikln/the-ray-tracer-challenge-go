package raytracer

import (
	"math"
)

type Pattern interface {
	ColorAt(Point) Color
	ColorAtObject(Intersectable, Point) Color
	SetTransform(m *Matrix)
	GetTransform() *Matrix
}

type StripePattern struct {
	A         Color
	B         Color
	Transform *Matrix
}

func NewStripePattern(a, b Color) *StripePattern {
	return &StripePattern{a, b, NewIdentityMatrix()}
}

func (sp *StripePattern) GetTransform() *Matrix {
	return sp.Transform
}
func (sp *StripePattern) SetTransform(m *Matrix) {
	sp.Transform = m
}
func (sp *StripePattern) ColorAt(p Point) Color {
	if math.Mod(math.Floor(p.X), 2) == 0 {
		return sp.A
	} else {
		return sp.B
	}
}

func (sp *StripePattern) ColorAtObject(object Intersectable, worldPoint Point) Color {
	objectPoint := object.GetTransform().Inverse().MulPoint(worldPoint)
	patternPoint := sp.GetTransform().Inverse().MulPoint(objectPoint)

	return sp.ColorAt(patternPoint)
}

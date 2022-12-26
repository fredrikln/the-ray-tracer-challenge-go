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

type GradientPattern struct {
	A         Color
	B         Color
	Transform *Matrix
}

func NewGradientPattern(a, b Color) *GradientPattern {
	return &GradientPattern{a, b, NewIdentityMatrix()}
}

func (gp *GradientPattern) GetTransform() *Matrix {
	return gp.Transform
}
func (gp *GradientPattern) SetTransform(m *Matrix) {
	gp.Transform = m
}
func (gp *GradientPattern) ColorAt(p Point) Color {
	distance := gp.B.Sub(gp.A)
	fraction := p.X - math.Floor(p.X)

	return gp.A.Add(distance.MulFloat(fraction))
}
func (gp *GradientPattern) ColorAtObject(object Intersectable, worldPoint Point) Color {
	objectPoint := object.GetTransform().Inverse().MulPoint(worldPoint)
	patternPoint := gp.GetTransform().Inverse().MulPoint(objectPoint)

	return gp.ColorAt(patternPoint)
}

type RingPattern struct {
	A         Color
	B         Color
	Transform *Matrix
}

func NewRingPattern(a, b Color) *RingPattern {
	return &RingPattern{a, b, NewIdentityMatrix()}
}

func (gp *RingPattern) GetTransform() *Matrix {
	return gp.Transform
}
func (gp *RingPattern) SetTransform(m *Matrix) {
	gp.Transform = m
}
func (gp *RingPattern) ColorAt(p Point) Color {
	if math.Mod(math.Floor(math.Sqrt(math.Pow(p.X, 2)+math.Pow(p.Z, 2))), 2) == 0 {
		return gp.A
	}

	return gp.B
}
func (gp *RingPattern) ColorAtObject(object Intersectable, worldPoint Point) Color {
	objectPoint := object.GetTransform().Inverse().MulPoint(worldPoint)
	patternPoint := gp.GetTransform().Inverse().MulPoint(objectPoint)

	return gp.ColorAt(patternPoint)
}

type CheckerPattern struct {
	A         Color
	B         Color
	Transform *Matrix
}

func NewCheckerPattern(a, b Color) *CheckerPattern {
	return &CheckerPattern{a, b, NewIdentityMatrix()}
}

func (gp *CheckerPattern) GetTransform() *Matrix {
	return gp.Transform
}
func (gp *CheckerPattern) SetTransform(m *Matrix) {
	gp.Transform = m
}
func (gp *CheckerPattern) ColorAt(p Point) Color {
	if math.Mod(math.Floor(p.X)+math.Floor(p.Y)+math.Floor(p.Z), 2) == 0 {
		return gp.A
	}

	return gp.B
}
func (gp *CheckerPattern) ColorAtObject(object Intersectable, worldPoint Point) Color {
	objectPoint := object.GetTransform().Inverse().MulPoint(worldPoint)
	patternPoint := gp.GetTransform().Inverse().MulPoint(objectPoint)

	return gp.ColorAt(patternPoint)
}

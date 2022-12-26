package raytracer

import (
	"math"
)

type Pattern interface {
	ColorAt(Point) Color
	ColorAtObject(Intersectable, Point) Color
	SetTransform(m *Matrix) Pattern
	GetTransform() *Matrix
}

// StripePattern

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
func (sp *StripePattern) SetTransform(m *Matrix) Pattern {
	sp.Transform = m

	return sp
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

// GradientPattern

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
func (gp *GradientPattern) SetTransform(m *Matrix) Pattern {
	gp.Transform = m

	return gp
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

// RingPattern

type RingPattern struct {
	A         Color
	B         Color
	Transform *Matrix
}

func NewRingPattern(a, b Color) *RingPattern {
	return &RingPattern{a, b, NewIdentityMatrix()}
}

func (rp *RingPattern) GetTransform() *Matrix {
	return rp.Transform
}
func (rp *RingPattern) SetTransform(m *Matrix) Pattern {
	rp.Transform = m

	return rp
}
func (rp *RingPattern) ColorAt(p Point) Color {
	if math.Mod(math.Floor(math.Sqrt(math.Pow(p.X, 2)+math.Pow(p.Z, 2))), 2) == 0 {
		return rp.A
	}

	return rp.B
}
func (rp *RingPattern) ColorAtObject(object Intersectable, worldPoint Point) Color {
	objectPoint := object.GetTransform().Inverse().MulPoint(worldPoint)
	patternPoint := rp.GetTransform().Inverse().MulPoint(objectPoint)

	return rp.ColorAt(patternPoint)
}

// CheckerPattern

type CheckerPattern struct {
	A         Color
	B         Color
	Transform *Matrix
}

func NewCheckerPattern(a, b Color) *CheckerPattern {
	return &CheckerPattern{a, b, NewIdentityMatrix()}
}

func (cp *CheckerPattern) GetTransform() *Matrix {
	return cp.Transform
}
func (cp *CheckerPattern) SetTransform(m *Matrix) Pattern {
	cp.Transform = m

	return cp
}
func (cp *CheckerPattern) ColorAt(p Point) Color {
	if math.Mod(math.Floor(p.X)+math.Floor(p.Y)+math.Floor(p.Z), 2) == 0 {
		return cp.A
	}

	return cp.B
}
func (cp *CheckerPattern) ColorAtObject(object Intersectable, worldPoint Point) Color {
	objectPoint := object.GetTransform().Inverse().MulPoint(worldPoint)
	patternPoint := cp.GetTransform().Inverse().MulPoint(objectPoint)

	return cp.ColorAt(patternPoint)
}

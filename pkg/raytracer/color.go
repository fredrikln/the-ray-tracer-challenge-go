package raytracer

import (
	"image/color"
	"math"

	c "github.com/fredrikln/the-ray-tracer-challenge-go/common"
)

type Color struct {
	R, G, B float64
}

func NewColor(r, g, b float64) Color {
	return Color{
		r, g, b,
	}
}

func (a Color) Eq(b Color) bool {
	if !c.WithinTolerance(a.R, b.R, 1e-5) {
		return false
	}
	if !c.WithinTolerance(a.G, b.G, 1e-5) {
		return false
	}
	if !c.WithinTolerance(a.B, b.B, 1e-5) {
		return false
	}

	return true
}

func (a Color) Add(b Color) Color {
	return Color{
		a.R + b.R,
		a.G + b.G,
		a.B + b.B,
	}
}

func (a Color) Sub(b Color) Color {
	return Color{
		a.R - b.R,
		a.G - b.G,
		a.B - b.B,
	}
}

func (a Color) MulFloat(b float64) Color {
	return Color{
		a.R * b,
		a.G * b,
		a.B * b,
	}
}

func (a Color) Mul(b Color) Color {
	return Color{
		a.R * b.R,
		a.G * b.G,
		a.B * b.B,
	}
}

func (c Color) GetRGBA() color.Color {
	r := uint8(math.Min(math.Max(math.Round(c.R*255), 0), 255))
	g := uint8(math.Min(math.Max(math.Round(c.G*255), 0), 255))
	b := uint8(math.Min(math.Max(math.Round(c.B*255), 0), 255))

	return color.RGBA{r, g, b, 0xFF}
}

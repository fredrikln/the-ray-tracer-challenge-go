package raytracer

import (
	"math"
)

type Camera struct {
	Hsize      int
	Vsize      int
	Fov        float64
	Transform  *Matrix
	PixelSize  float64
	HalfWidth  float64
	HalfHeight float64
}

func NewCamera(hsize, vsize int, fov float64) *Camera {
	halfView := math.Tan(fov / 2)
	aspect := float64(hsize) / float64(vsize)

	var halfWidth, halfHeight float64

	if aspect > 1 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfWidth = halfView * aspect
		halfHeight = halfView
	}

	return &Camera{
		Hsize:      hsize,
		Vsize:      vsize,
		Fov:        fov,
		Transform:  NewIdentityMatrix(),
		PixelSize:  (halfWidth * 2) / float64(hsize),
		HalfWidth:  halfWidth,
		HalfHeight: halfHeight,
	}
}

func (c *Camera) RayForPixel(x, y int) Ray {
	xOffset := (float64(x) + 0.5) * c.PixelSize
	yOffset := (float64(y) + 0.5) * c.PixelSize

	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset

	pixel := c.Transform.Inverse().MulPoint(NewPoint(worldX, worldY, -1))
	origin := c.Transform.Inverse().MulPoint(NewPoint(0, 0, 0))

	direction := pixel.Sub(origin).Norm()

	return NewRay(origin, direction)
}

func (c *Camera) SetTransform(m *Matrix) *Camera {
	c.Transform = m

	return c
}

func (c *Camera) Render(w *World) *Canvas {
	canvas := NewCanvas(c.Hsize, c.Vsize)

	for y := 0; y < c.Vsize; y++ {
		for x := 0; x < c.Hsize; x++ {
			ray := c.RayForPixel(x, y)
			color := w.ColorAt(ray)

			canvas.SetPixel(x, y, color)
		}
	}

	return canvas
}

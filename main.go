package main

import (
	"math"

	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
	r "github.com/fredrikln/the-ray-tracer-challenge-go/render"
)

type Projectile struct {
	Position g.Vec
	Velocity g.Vec
}

type Environment struct {
	Gravity g.Vec
	Wind    g.Vec
}

func tick(p Projectile, e Environment) Projectile {
	return Projectile{
		p.Position.Add(p.Velocity),
		p.Velocity.Add(e.Gravity).Add(e.Wind),
	}
}

func main() {
	// Set up canvas
	canvas := r.NewCanvas(500, 500)
	c := r.NewColor(1.0, 0.0, 0.0)

	point := g.NewVec(0, 1, 0)
	rotation := g.NewRotationZ(2 * math.Pi / 60)

	for i := 0; i < 60; i += 1 {
		point = point.MulMat(rotation)

		x := int(point.X*200) + 250
		y := int(point.Y*200) + 250
		canvas.SetPixel(x, y, c)
	}

	canvas.SavePNG("test.png")
}

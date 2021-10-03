package main

import (
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
	pixels := 500
	canvas := r.NewCanvas(pixels, pixels)
	c := r.NewColor(1.0, 0.0, 0.0)

	rayOrigin := g.NewPoint(0, 0, -5)

	wallZ := 10.0
	wallSize := 7.0

	pixelSize := wallSize / float64(pixels)

	half := wallSize / 2

	shape := g.NewSphere()

	for y := 0; y < pixels; y += 1 {
		worldY := half - pixelSize*float64(y)

		for x := 0; x < pixels; x += 1 {
			worldX := -half + pixelSize*float64(x)

			position := g.NewPoint(worldX, worldY, wallZ)

			r := g.NewRay(rayOrigin, position.Sub(rayOrigin).Norm())
			xs := shape.Intersect(r)

			if _, hit := g.GetHit(xs); hit {
				canvas.SetPixel(x, y, c)
			}
		}
	}

	canvas.SavePNG("test.png")
}

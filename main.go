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

	// Ray starting at Z -5
	rayOrigin := g.NewPoint(0, 0, -5)

	// Wall at Z 10
	wallZ := 10.0
	wallSize := 7.0

	// 7/500 = 0.014
	pixelSize := wallSize / float64(pixels)

	// 3.5
	half := wallSize / 2

	// Shape is at Z 0
	shape := g.NewSphere()

	for y := 0; y < pixels; y += 1 {
		// 3.5 - (0.014 * 0) to 3.5 - (0.014 * 500) = 3.5 to 3.5-7 = 3.5 to -3.5
		worldY := half - pixelSize*float64(y)

		for x := 0; x < pixels; x += 1 {
			// -3.5 + (0.014 * 0) to -3.5 + (0.014 * 500) = -3.5 to -3.5+7 = -3.5 to 3.5
			worldX := -half + pixelSize*float64(x)

			// (-3.5 to 3.5), (3.5 to -3.5), 10
			position := g.NewPoint(worldX, worldY, wallZ)

			// a ray pointing from 0,0,0 to a point at (-3.5 to 3.5), (3.5 to -3.5), 10
			r := g.NewRay(rayOrigin, position.Sub(rayOrigin).Norm())
			// get intersections for ray and sphere
			xs := shape.Intersect(r)

			// check if we have a hit
			if _, hit := g.GetHit(xs); hit {
				// draw hit to canvas at x, y with color
				canvas.SetPixel(x, y, c)
			}
		}
	}

	canvas.SavePNG("test.png")
}

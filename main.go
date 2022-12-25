package main

import (
	"fmt"
	"time"

	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
	m "github.com/fredrikln/the-ray-tracer-challenge-go/material"
	r "github.com/fredrikln/the-ray-tracer-challenge-go/render"
	s "github.com/fredrikln/the-ray-tracer-challenge-go/shapes"
)

func main() {
	// Set up canvas
	pixels := 500
	canvas := r.NewCanvas(pixels, pixels)

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
	shape := s.NewSphere()
	mat := m.NewMaterial()
	mat.SetColor(m.NewColor(1, 0.2, 1))
	shape.SetMaterial(mat)

	light := m.NewPointLight(g.NewPoint(-10, 10, -10), m.NewColor(1, 1, 1))

	timeBefore := time.Now()

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
			if hit, didHit := g.GetHit(xs); didHit {
				point := r.Position(hit.Time)
				normal := (*hit.Object).NormalAt(point)
				eyev := r.Direction.Mul(-1)

				color := shape.Material.Lighting(light, point, eyev, normal)

				// draw hit to canvas at x, y with color
				canvas.SetPixel(x, y, color)
			}
		}
	}

	timeAfter := time.Now()

	diff := timeAfter.Sub(timeBefore)

	fmt.Println("Render time:", diff)

	canvas.SavePNG("test.png")
}

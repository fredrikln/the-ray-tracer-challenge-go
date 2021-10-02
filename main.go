package main

import (
	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
	r "github.com/fredrikln/the-ray-tracer-challenge-go/render"
)

type Projectile struct {
	Position *g.Vec
	Velocity *g.Vec
}

type Environment struct {
	Gravity *g.Vec
	Wind    *g.Vec
}

func tick(p Projectile, e Environment) Projectile {
	return Projectile{
		p.Position.Add(p.Velocity),
		p.Velocity.Add(e.Gravity).Add(e.Wind),
	}
}

func main() {
	// Set up canvas
	canvas := r.NewCanvas(900, 550)
	c := r.NewColor(1.0, 0.0, 0.0)

	// Set up projectile and environment
	projectile := Projectile{Position: g.NewVec(0, 1, 0), Velocity: g.NewVec(1, 1.8, 0).Norm().Mul(11.25)}
	environment := Environment{g.NewVec(0, -0.1, 0), g.NewVec(-0.01, 0, 0)}

	for projectile.Position.Y > 0 {
		projectile = tick(projectile, environment)

		if projectile.Position.Y > 0 {
			canvas.SetPixel(int(projectile.Position.X), canvas.Height-int(projectile.Position.Y), c)
		}
	}

	canvas.SavePNG("test.png")
}

package main

import (
	"fmt"

	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
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
	projectile := Projectile{Position: g.NewVec(0, 1, 0), Velocity: g.NewVec(1, 1, 0).Norm()}
	environment := Environment{g.NewVec(0, -0.1, 0), g.NewVec(-0.01, 0, 0)}

	for projectile.Position.Y > 0 {
		projectile = tick(projectile, environment)

		fmt.Printf("%+v\n", projectile)
	}
}

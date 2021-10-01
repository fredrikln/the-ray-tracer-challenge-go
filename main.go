package main

import (
	"fmt"

	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
)

func main() {
	vector := g.NewVec(1, 2, 3)
	point := g.NewVec(2, 3, 4)

	newVector := vector.Add(vector)
	newPoint := vector.Add(point)

	fmt.Println("Hello world", newVector, newPoint)
}

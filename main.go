package main

import (
	"fmt"
	"math"
	"time"

	r "github.com/fredrikln/the-ray-tracer-challenge-go/pkg/raytracer"
)

func main() {
	// Set up canvas
	w := r.NewWorld()

	p0 := r.NewRingPattern(r.NewColor(1, 0.9, 0.9), r.NewColor(1, 1, 1))
	m1 := r.NewMaterial().SetColor(r.NewColor(1, 0.9, 0.9)).SetSpecular(0).SetPattern(p0)
	floor := r.NewPlane().SetMaterial(m1)
	w.AddObject(floor)

	t1 := r.NewTranslation(0, 0, 5).Mul(r.NewRotationY(-(math.Pi / 4))).Mul(r.NewRotationX(math.Pi / 2))
	leftWall := r.NewPlane().SetTransform(t1).SetMaterial(floor.Material)
	w.AddObject(leftWall)

	t2 := r.NewTranslation(0, 0, 5).Mul(r.NewRotationY((math.Pi / 4))).Mul(r.NewRotationX(math.Pi / 2))
	rightWall := r.NewPlane().SetTransform(t2).SetMaterial(floor.Material)
	w.AddObject(rightWall)

	p1 := r.NewStripePattern(r.NewColor(0.25, 0.5, 0.25), r.NewColor(1, 1, 1))
	p1.SetTransform(r.NewRotationY(math.Pi / 4).Mul(r.NewScaling(0.35, 0.35, 0.35)))
	t3 := r.NewTranslation(-0.5, 1, 0.5)
	m2 := r.NewMaterial().SetColor(r.NewColor(0.1, 1, 0.5)).SetDiffuse(0.7).SetSpecular(0.3).SetPattern(p1)
	middle := r.NewSphere().SetTransform(t3).SetMaterial(m2)
	w.AddObject(middle)

	p2 := r.NewGradientPattern(r.NewColor(0.5, 0.25, 0.25), r.NewColor(1, 1, 1))
	p2.SetTransform(r.NewTranslation(1, 0, 0).Mul(r.NewScaling(2, 2, 2)))
	t4 := r.NewTranslation(1.5, 0.5, -0.5).Mul(r.NewScaling(0.5, 0.5, 0.5))
	m3 := r.NewMaterial().SetColor(r.NewColor(0.5, 1, 0.1)).SetDiffuse(0.7).SetSpecular(0.3).SetPattern(p2)
	right := r.NewSphere().SetTransform(t4).SetMaterial(m3)
	w.AddObject(right)

	p3 := r.NewCheckerPattern(r.NewColor(0.25, 0.25, 0.5), r.NewColor(1, 1, 1))
	p3.SetTransform(r.NewRotationZ(-math.Pi / 4).Mul(r.NewScaling(0.5, 0.15, 0.15)))
	t5 := r.NewTranslation(-1.5, 0.33, -0.75).Mul(r.NewScaling(0.33, 0.33, 0.33))
	m4 := r.NewMaterial().SetColor(r.NewColor(1, 0.8, 0.1)).SetDiffuse(0.7).SetSpecular(0.3).SetPattern(p3)
	left := r.NewSphere().SetTransform(t5).SetMaterial(m4)
	w.AddObject(left)

	w.AddLight(r.NewPointLight(r.NewPoint(-10, 10, -10), r.NewColor(1, 1, 1)))

	ct := r.ViewTransform(r.NewPoint(0, 1.5, -5), r.NewPoint(0, 1, 0), r.NewVec(0, 1, 0))
	camera := r.NewCamera(500, 500, math.Pi/3).SetTransform(ct)

	timeBefore := time.Now()

	canvas := camera.Render(w)

	timeAfter := time.Now()

	diff := timeAfter.Sub(timeBefore)

	fmt.Println("Render time:", diff)

	canvas.SavePNG("test.png")
}

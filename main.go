package main

import (
	"fmt"
	"math"
	"runtime"
	"time"

	r "github.com/fredrikln/the-ray-tracer-challenge-go/pkg/raytracer"
)

func main() {
	// Set up canvas
	w := r.NewWorld()

	whiteMaterial := r.NewMaterial().SetColor(r.NewColor(1, 1, 1)).SetDiffuse(0.7).SetSpecular(0.0).SetReflective(0.1).SetAmbient(0.1)
	blueMaterial := r.NewMaterial().SetColor(r.NewColor(0.537, 0.831, 0.914)).SetDiffuse(0.7).SetSpecular(0.0).SetReflective(0.1).SetAmbient(0.1)
	redMaterial := r.NewMaterial().SetColor(r.NewColor(0.941, 0.322, 0.388)).SetDiffuse(0.7).SetSpecular(0.0).SetReflective(0.1).SetAmbient(0.1)
	purpleMaterial := r.NewMaterial().SetColor(r.NewColor(0.373, 0.404, 0.550)).SetDiffuse(0.7).SetSpecular(0.0).SetReflective(0.1).SetAmbient(0.1)

	floor := r.NewPlane().SetMaterial(whiteMaterial)
	w.AddObject(floor)

	t1 := r.NewTranslation(0, 0, 5).Mul(r.NewRotationY(-(math.Pi / 4))).Mul(r.NewRotationX(math.Pi / 2))
	leftWall := r.NewPlane().SetTransform(t1).SetMaterial(whiteMaterial)
	w.AddObject(leftWall)

	t2 := r.NewTranslation(0, 0, 5).Mul(r.NewRotationY((math.Pi / 4))).Mul(r.NewRotationX(math.Pi / 2))
	rightWall := r.NewPlane().SetTransform(t2).SetMaterial(whiteMaterial)
	w.AddObject(rightWall)

	t3 := r.NewTranslation(-0.5, 1, 0.5)
	middle := r.NewSphere().SetTransform(t3).SetMaterial(blueMaterial)
	w.AddObject(middle)

	t4 := r.NewTranslation(1.5, 0.5, -0.5).Mul(r.NewScaling(0.5, 0.5, 0.5))
	right := r.NewSphere().SetTransform(t4).SetMaterial(redMaterial)
	w.AddObject(right)

	t5 := r.NewTranslation(-1.5, 0.33, -0.75).Mul(r.NewScaling(0.33, 0.33, 0.33))
	left := r.NewSphere().SetTransform(t5).SetMaterial(purpleMaterial)
	w.AddObject(left)

	w.AddLight(r.NewPointLight(r.NewPoint(-10, 10, -10), r.NewColor(1, 1, 1)))

	ct := r.ViewTransform(r.NewPoint(0, 1.5, -5), r.NewPoint(0, 1, 0), r.NewVec(0, 1, 0))
	width := 640
	ratio := 16.0 / 9.0

	camera := r.NewCamera(width, int(float64(width)/ratio), math.Pi/3).SetTransform(ct)
	//camera.Bounces = 10
	//camera.Antialiasing = true

	timeBefore := time.Now()

	canvas := camera.RenderMultiThreaded(w, runtime.NumCPU())

	timeAfter := time.Now()

	diff := timeAfter.Sub(timeBefore)

	fmt.Println("Render time:", diff)

	canvas.SavePNG("test.png")
}

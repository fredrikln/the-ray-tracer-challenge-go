package main

import (
	"fmt"

	"math"
	"math/rand"

	"runtime"

	"time"

	r "github.com/fredrikln/the-ray-tracer-challenge-go/pkg/raytracer"
)

func GetTestScene() (*r.World, *r.Matrix) {
	w := r.NewWorld()

	whiteMaterial := r.NewMaterial().SetColor(r.NewColor(1, 1, 1)).SetDiffuse(0.7).SetSpecular(0.0).SetAmbient(0.1).SetReflective(0.1)
	// blueMaterial := r.NewMaterial().SetColor(r.NewColor(0.537, 0.831, 0.914)).SetDiffuse(0.7).SetSpecular(0.0).SetReflective(0.1).SetAmbient(0.1)
	redMaterial := r.NewMaterial().SetColor(r.NewColor(0.941, 0.322, 0.388)).SetDiffuse(0.7).SetSpecular(0.0).SetReflective(0.1).SetAmbient(0.1)
	purpleMaterial := r.NewMaterial().SetColor(r.NewColor(0.373, 0.404, 0.550)).SetDiffuse(0.7).SetSpecular(0.0).SetReflective(0.1).SetAmbient(0.1)

	mirror := r.NewMaterial().SetColor(r.NewColor(0, 0.1, 0)).SetReflective(0.75).SetDiffuse(0.05).SetAmbient(0.05).SetSpecular(1).SetShininess(300)
	glass := r.NewMaterial().SetColor(r.NewColor(0, 0.1, 0)).SetTransparency(0.9).SetReflective(0.9).SetRefractiveIndex(1.5).SetReflective(0.9).SetDiffuse(0.05).SetAmbient(0.05).SetSpecular(1).SetShininess(300)

	floor := r.NewPlane().SetMaterial(whiteMaterial)
	w.AddObject(floor)

	t1 := r.NewTranslation(0, 0, 5).Mul(r.NewRotationY(-(math.Pi / 4))).Mul(r.NewRotationX(math.Pi / 2))
	leftWall := r.NewPlane().SetTransform(t1).SetMaterial(mirror)
	w.AddObject(leftWall)

	t2 := r.NewTranslation(0, 0, 5).Mul(r.NewRotationY((math.Pi / 4))).Mul(r.NewRotationX(math.Pi / 2))
	rightWall := r.NewPlane().SetTransform(t2).SetMaterial(whiteMaterial)
	w.AddObject(rightWall)

	t3 := r.NewTranslation(-0.5, 1, 0.5)
	middle := r.NewSphere().SetTransform(t3).SetMaterial(glass)
	w.AddObject(middle)

	t4 := r.NewTranslation(1.5, 0.5, -0.5).Mul(r.NewScaling(0.5, 0.5, 0.5))
	right := r.NewSphere().SetTransform(t4).SetMaterial(redMaterial)
	w.AddObject(right)

	t5 := r.NewTranslation(-1.5, 0.33, -0.75).Mul(r.NewScaling(0.33, 0.33, 0.33))
	left := r.NewSphere().SetTransform(t5).SetMaterial(purpleMaterial)
	w.AddObject(left)

	w.AddLight(r.NewPointLight(r.NewPoint(-10, 10, -10), r.NewColor(1, 1, 1)))

	return w, r.ViewTransform(r.NewPoint(0, 1.5, -5), r.NewPoint(0, 1, 0), r.NewVec(0, 1, 0))
}

func GetTestScene2() (*r.World, *r.Matrix) {
	w := r.NewWorld()

	t := r.NewTranslation(0, 0, 500).Mul(r.NewRotationX(math.Pi / 2))
	wall := r.NewPlane().SetTransform(t)
	w.AddObject(wall)

	seed := 1672236286765 //time.Now().UnixMilli()
	rand.Seed(int64(seed))
	fmt.Println("Seed", seed)

	stepsEachSide := 2
	offset := 3

	for y := -stepsEachSide; y <= stepsEachSide; y++ {
		group := r.NewGroup()

		for x := -stepsEachSide; x <= stepsEachSide; x++ {
			for z := -stepsEachSide; z <= stepsEachSide; z++ {

				var object r.Intersectable

				objectType := rand.Intn(4)
				switch objectType {
				case 0:
					object = r.NewSphere()
				case 1:
					object = r.NewCube()
				case 2:
					cy := r.NewCylinder()
					cy.Minimum = -1
					cy.Maximum = 1
					cy.Closed = true
					object = cy
				case 3:
					co := r.NewCone()
					co.Minimum = -1
					co.Maximum = 0
					co.Closed = true

					t := r.NewScaling(1, 2, 1).Mul(r.NewTranslation(0, 0.5, 0))
					co.SetTransform(t)

					object = co
				default:
					object = r.NewSphere()
				}

				objTransform := r.NewTranslation(float64(offset*x), 0, float64(offset*z)).Mul(object.GetTransform())

				object.SetTransform(objTransform)

				var material *r.Material

				materialType := rand.Intn(4)
				switch materialType {
				case 0:
					material = r.NewMaterial().SetColor(r.NewColor(0.537, 0.831, 0.914)).SetDiffuse(0.7).SetSpecular(0.0).SetReflective(0.1).SetAmbient(0.1)
				case 1:
					material = r.NewMaterial().SetColor(r.NewColor(0.941, 0.322, 0.388)).SetDiffuse(0.7).SetSpecular(0.0).SetReflective(0.1).SetAmbient(0.1)
				case 2:
					material = r.NewMaterial().SetColor(r.NewColor(0.373, 0.404, 0.550)).SetDiffuse(0.7).SetSpecular(0.0).SetReflective(0.1).SetAmbient(0.1)
				case 3:
					material = r.NewMaterial().SetColor(r.NewColor(0.373, 0.404, 0.550)).SetTransparency(1).SetReflective(1).SetRefractiveIndex(1.5).SetSpecular(1).SetShininess(300).SetDiffuse(0.05).SetAmbient(0.05)
				default:
					material = r.NewMaterial().SetColor(r.NewColor(1, 0, 0)).SetAmbient(1).SetDiffuse(0).SetShininess(0).SetSpecular(0)
				}

				object.SetMaterial(material)

				group.AddChild(object)
			}
		}
		group.SetTransform(r.NewTranslation(0, float64(offset*y), 0))
		w.AddObject(group)
	}

	w.AddLight(r.NewPointLight(r.NewPoint(50, 100, -50), r.NewColor(1, 1, 1)))
	w.AddLight(r.NewPointLight(r.NewPoint(-400, 50, -10), r.NewColor(0.2, 0.2, 0.2)))

	return w, r.ViewTransform(r.NewPoint(-10, 10, -30), r.NewPoint(0, -1.5, 0), r.NewVec(0, 1, 0))
}

func GetTestScene3() (*r.World, *r.Matrix) {
	w := r.NewWorld()

	t := r.NewTranslation(0, 0, 10).Mul(r.NewRotationX(math.Pi / 2))
	p := r.NewCheckerPattern(r.NewColor(0, 0, 0), r.NewColor(1, 1, 1))
	m := r.NewMaterial().SetPattern(p)
	wall := r.NewPlane().SetTransform(t).SetMaterial(m)

	w.AddObject(wall)

	glass := r.NewMaterial().SetColor(r.NewColor(0.373, 0.404, 0.550)).SetTransparency(0.9).SetReflective(0.9).SetRefractiveIndex(1.5).SetSpecular(0.9).SetShininess(300).SetDiffuse(0.05).SetAmbient(0.05)
	air := r.NewMaterial().SetTransparency(0.9).SetReflective(0.9).SetRefractiveIndex(1).SetSpecular(0.9).SetShininess(300).SetDiffuse(0.05).SetAmbient(0.05)

	s1 := r.NewSphere().SetMaterial(air)
	w.AddObject(s1)

	s2 := r.NewSphere().SetMaterial(glass).SetTransform(r.NewScaling(2, 2, 2))
	w.AddObject(s2)

	w.AddLight(r.NewPointLight(r.NewPoint(0, 100, -100), r.NewColor(1, 1, 1)))

	return w, r.ViewTransform(r.NewPoint(0, 0, -10), r.NewPoint(0, 0, 0), r.NewVec(0, 1, 0))
}

func GetTestScene4() (*r.World, *r.Matrix) {
	w := r.NewWorld()

	farleft := r.NewPoint(-1, 0, -1)
	farright := r.NewPoint(1, 0, -1)
	nearleft := r.NewPoint(-1, 0, 1)
	nearright := r.NewPoint(1, 0, 1)
	top := r.NewPoint(0, 1, 0)

	m1 := r.NewMaterial().SetColor(r.NewColor(0.373, 0.404, 0.550)).SetTransparency(0.7).SetReflective(0.7).SetRefractiveIndex(1.5).SetSpecular(1).SetShininess(200).SetDiffuse(0.2).SetAmbient(0.0)
	// m1 := r.NewMaterial().SetColor(r.NewColor(1, 0, 0))
	t1 := r.NewTriangle(nearleft, farleft, farright)
	t1.SetMaterial(m1)
	t2 := r.NewTriangle(nearleft, nearright, farright)
	t2.SetMaterial(m1)

	t3 := r.NewTriangle(nearleft, top, farleft)
	t3.SetMaterial(m1)
	t4 := r.NewTriangle(farleft, top, farright)
	t4.SetMaterial(m1)
	t5 := r.NewTriangle(farright, top, nearright)
	t5.SetMaterial(m1)
	t6 := r.NewTriangle(nearright, top, nearleft)
	t6.SetMaterial(m1)

	g := r.NewGroup()
	g.AddChild(t1)
	g.AddChild(t2)
	g.AddChild(t3)
	g.AddChild(t4)
	g.AddChild(t5)
	g.AddChild(t6)

	g.SetTransform(r.NewRotationY(math.Pi / 4))
	w.AddObject(g)

	m2 := r.NewMaterial().SetColor(r.NewColor(1, 1, 1)).SetPattern(r.NewCheckerPattern(r.NewColor(0.5, 0.5, 0.5), r.NewColor(1, 1, 1)))
	p := r.NewPlane().SetTransform(r.NewTranslation(0, -10, 0)).SetMaterial(m2)
	w.AddObject(p)

	w.AddLight(r.NewPointLight(r.NewPoint(50, 100, -50), r.NewColor(1, 1, 1)))
	w.AddLight(r.NewPointLight(r.NewPoint(-400, 50, -10), r.NewColor(0.2, 0.2, 0.2)))

	return w, r.ViewTransform(r.NewPoint(0, 1, -4), r.NewPoint(0, 0.25, 0), r.NewVec(0, 1, 0))
}

func main() {
	// Set up canvas
	w, ct := GetTestScene2()

	width := 640
	ratio := 16.0 / 9.0

	camera := r.NewCamera(width, int(float64(width)/ratio), math.Pi/3).SetTransform(ct)
	// camera.Bounces = 6
	// camera.Antialiasing = true

	timeBefore := time.Now()

	canvas := camera.RenderMultiThreaded(w, runtime.NumCPU())

	timeAfter := time.Now()

	diff := timeAfter.Sub(timeBefore)

	fmt.Println("Render time:", diff)

	canvas.SavePNG("test.png")
}

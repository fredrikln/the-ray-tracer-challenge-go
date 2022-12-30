package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"strings"
	"time"

	"github.com/fredrikln/the-ray-tracer-challenge-go/pkg/objparser"
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
	teapotData, err := os.ReadFile("models/teapot2.obj")

	if err != nil {
		panic(err)
	}

	w := r.NewWorld()

	t := r.NewTranslation(0, 0, 500).Mul(r.NewRotationX(math.Pi / 2))
	wall := r.NewPlane().SetTransform(t)
	w.AddObject(wall)

	seed := 1672354185622 //time.Now().UnixMilli()
	rand.Seed(int64(seed))
	fmt.Println("Seed", seed)

	stepsEachSide := 2
	offset := 3

	globalGroup := r.NewGroup()

	for x := -stepsEachSide; x <= stepsEachSide; x++ {
		for y := -stepsEachSide; y <= stepsEachSide; y++ {
			for z := -stepsEachSide; z <= stepsEachSide; z++ {

				var object r.Intersectable

				objectType := rand.Intn(5)
				// objectType := rand.Intn(4)
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
				case 4:
					g := r.NewGroup()
					g.SetTransform(r.NewScaling(0.1, 0.1, 0.1))

					object = g
				default:
					object = r.NewSphere()
				}

				objTransform := r.NewTranslation(float64(offset*x), float64(offset*y), float64(offset*z)).Mul(object.GetTransform())

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

				if objectType == 4 {
					p := objparser.NewParser()
					p.SetMaterial(material)
					teapot := p.Parse(strings.Trim(string(teapotData), "\n"))

					object.(*r.Group).AddChild(teapot)
				} else {
					object.SetMaterial(material)
				}

				globalGroup.AddChild(object)
			}
		}
	}

	globalGroup.Divide(1)

	w.AddObject(globalGroup)

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

func GetTestScene5() (*r.World, *r.Matrix) {
	w := r.NewWorld()

	content, err := os.ReadFile("models/teapot2.obj")

	if err != nil {
		panic(err)
	}

	p := objparser.NewParser()
	// m := r.NewMaterial().SetColor(r.NewColor(0.373, 0.404, 0.550)).SetTransparency(1).SetReflective(1).SetRefractiveIndex(1.5).SetSpecular(1).SetShininess(300).SetDiffuse(0.05).SetAmbient(0.05)
	m := r.NewMaterial().SetColor(r.NewColor(0.75, 0.75, 0.75))
	p.SetMaterial(m)
	g := p.Parse(strings.Trim(string(content), "\n"))
	t1 := r.NewScaling(0.25, 0.25, 0.25)
	g.SetTransform(t1)

	w.AddObject(g)

	pl := r.NewPlane()
	t2 := r.NewTranslation(0, 0, 30).Mul(r.NewRotationX(math.Pi / 2))
	pl.SetTransform(t2)
	pl.SetMaterial(r.NewMaterial().SetPattern(r.NewCheckerPattern(r.NewColor(0.5, 0.5, 0.5), r.NewColor(1, 1, 1))))

	w.AddObject(pl)

	w.AddLight(r.NewPointLight(r.NewPoint(50, 100, -50), r.NewColor(1, 1, 1)))
	w.AddLight(r.NewPointLight(r.NewPoint(-400, 50, -10), r.NewColor(0.2, 0.2, 0.2)))

	return w, r.ViewTransform(r.NewPoint(0, 5, -10), r.NewPoint(0, 0, 0), r.NewVec(0, 1, 0))
}

func GetTestScene6() (*r.World, *r.Matrix) {
	w := r.NewWorld()

	pl := r.NewPlane()
	pl.SetTransform(r.NewTranslation(0, -1, 0))
	pl.SetMaterial(r.NewMaterial().SetPattern(r.NewCheckerPattern(r.NewColor(0.5, 0.5, 0.5), r.NewColor(1, 1, 1))))
	w.AddObject(pl)

	c := r.NewCube()
	c.SetMaterial(r.NewMaterial().SetColor(r.NewColor(1, 1, 0)))
	sp := r.NewSphere()
	sp.SetTransform(r.NewScaling(math.Sqrt(2), math.Sqrt(2), math.Sqrt(2)))
	sp.SetMaterial(r.NewMaterial().SetColor(r.NewColor(1, 0, 1)))

	csg1 := r.NewCSG(r.Intersect, c, sp)

	cy1 := r.NewCylinder()
	cy1.Minimum = -10
	cy1.Maximum = 10
	cy1.Closed = true
	cy1.SetTransform(r.NewScaling(0.5, 0.5, 0.5))
	cy1.SetMaterial(r.NewMaterial().SetColor(r.NewColor(1, 0, 0)))

	cy2 := r.NewCylinder()
	cy2.Minimum = -10
	cy2.Maximum = 10
	cy2.Closed = true
	cy2.SetTransform(r.NewScaling(0.5, 0.5, 0.5).RotateX(math.Pi / 2))
	cy2.SetMaterial(r.NewMaterial().SetColor(r.NewColor(0, 1, 0)))

	csg2 := r.NewCSG(r.Union, cy1, cy2)

	cy3 := r.NewCylinder()
	cy3.Minimum = -10
	cy3.Maximum = 10
	cy3.Closed = true

	cy3.SetTransform(r.NewScaling(0.5, 0.5, 0.5).RotateX(math.Pi / 2).RotateZ(math.Pi / 2))
	cy3.SetMaterial(r.NewMaterial().SetColor(r.NewColor(0, 0, 1)))

	csg3 := r.NewCSG(r.Union, csg2, cy3)
	csg4 := r.NewCSG(r.Difference, csg1, csg3)
	t := r.NewTranslation(0, 1, 0).RotateY(math.Pi/5).Scale(2, 2, 2)
	csg4.SetTransform(t)

	w.AddObject(csg4)

	w.AddLight(r.NewPointLight(r.NewPoint(50, 100, -50), r.NewColor(1, 1, 1)))
	w.AddLight(r.NewPointLight(r.NewPoint(-400, 50, -10), r.NewColor(0.2, 0.2, 0.2)))

	return w, r.ViewTransform(r.NewPoint(0, 5, -10), r.NewPoint(0, 0.5, 0), r.NewVec(0, 1, 0))
}

func GetTestScene7() (*r.World, *r.Matrix) {
	w := r.NewWorld()

	content, err := os.ReadFile("models/dragon.obj")

	if err != nil {
		panic(err)
	}

	transform := r.NewTranslation(0, -2.25, 0)
	checker := r.NewCheckerPattern(r.NewColor(0.5, 0.5, 0.5), r.NewColor(1, 1, 1))
	material := r.NewMaterial().SetPattern(checker)
	wall := r.NewPlane().SetTransform(transform).SetMaterial(material)

	w.AddObject(wall)

	p := objparser.NewParser()
	// m := r.NewMaterial().SetColor(r.NewColor(0.373, 0.404, 0.550)).SetTransparency(1).SetReflective(1).SetRefractiveIndex(1.5).SetSpecular(1).SetShininess(300).SetDiffuse(0.05).SetAmbient(0.05)
	m := r.NewMaterial().SetColor(r.NewColor(0.135, 0.2225, 0.1575)).SetDiffuse(0.63).SetSpecular(0.316228).SetShininess(12.9).SetReflective(0.05)
	p.SetMaterial(m)
	g := p.Parse(strings.Trim(string(content), "\n"))
	t1 := r.NewTranslation(0, -2.25, 0)
	g.SetTransform(t1)

	g.Divide(1)

	w.AddObject(g)

	w.AddLight(r.NewPointLight(r.NewPoint(50, 100, -50), r.NewColor(1, 1, 1)))
	w.AddLight(r.NewPointLight(r.NewPoint(-400, 50, -10), r.NewColor(0.2, 0.2, 0.2)))

	return w, r.ViewTransform(r.NewPoint(0, 5, -10), r.NewPoint(0, 0, 0), r.NewVec(0, 1, 0))
}

func GetTestScene8() (*r.World, *r.Matrix) {
	w := r.NewWorld()

	g := r.NewGroup()

	c := r.NewColor(0.1, 0.1, 0.1)
	w.Background = &c

	// floor := r.NewPlane()
	// floor.NewMaterial = r.NewDiffuse(r.NewColor(0.7, 0.8, 0.7))
	// floor.SetTransform(r.NewTranslation(0, -1, 0))
	// g.AddChild(floor)

	floor := r.NewGroup()

	m := r.NewDiffuse(r.NewColor(0.7, 0.8, 0.7))
	localSource := rand.New(rand.NewSource(1337))
	count := 20
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			b := r.NewCube()
			b.NewMaterial = m

			b.SetTransform(r.NewTranslation(float64(i)*2, -1*localSource.Float64(), float64(j)*2).Translate(-float64(count), -2, -float64(count)))

			floor.AddChild(b)
		}
	}

	g.AddChild(floor)

	// container := r.NewSphere()
	// container.SetTransform(r.NewScaling(50, 50, 50))
	// containerMaterial := r.NewDiffuse(r.NewColor(0.1, 0.1, 0.1))
	// container.NewMaterial = containerMaterial

	// g.AddChild(container)

	// leftwall := r.NewPlane()
	// leftwall.NewMaterial = r.NewDiffuse(r.NewColor(0.8, 0.7, 0.7))
	// leftwall.SetTransform(r.NewTranslation(-15, 0, 0).RotateX(math.Pi / 2).RotateZ(math.Pi / 2))
	// g.AddChild(leftwall)

	// rightwall := r.NewPlane()
	// rightwall.NewMaterial = &r.Diffuse{Albedo: r.NewColor(0.7, 0.7, 0.8)}
	// rightwall.SetTransform(r.NewTranslation(15, 0, 0).RotateX(math.Pi / 2).RotateZ(math.Pi / 2))
	// g.AddChild(rightwall)

	// backwall := r.NewPlane()
	// backwall.NewMaterial = &r.Diffuse{Albedo: r.NewColor(0.8, 0.8, 0.7)}
	// backwall.SetTransform(r.NewTranslation(0, 0, 15).RotateX(math.Pi / 2))
	// g.AddChild(backwall)

	// behindcamerawall := r.NewPlane()
	// behindcamerawall.NewMaterial = &r.Diffuse{Albedo: r.NewColor(0.7, 0.8, 0.8)}
	// behindcamerawall.SetTransform(r.NewTranslation(0, 0, -15).RotateX(math.Pi / 2))
	// g.AddChild(behindcamerawall)

	light := r.NewCube()
	light.NewMaterial = r.NewEmissive(r.NewColor(1, 1, 1))
	// light.NewMaterial = r.NewDiffuse(r.NewColor(0.5, 0.5, 0.3))
	light.SetTransform(r.NewTranslation(0, 10, 0).Scale(10, 0.01, 10))
	g.AddChild(light)

	// content, err := os.ReadFile("models/dragon.obj")

	// if err != nil {
	// 	panic(err)
	// }

	// p := objparser.NewParser()
	// p.NewMaterial = r.NewMetal(r.NewColor(0.135, 0.2225, 0.1575), 0.75)
	// g2 := p.Parse(strings.Trim(string(content), "\n"))
	// t1 := r.NewTranslation(0, -2.25, 0)
	// g2.SetTransform(t1)

	// g.AddChild(g2)

	s1 := r.NewSphere()
	m2 := r.NewDiffuse(r.NewColor(0.8, 0.5, 0.5))
	s1.NewMaterial = m2
	g.AddChild(s1)

	s2 := r.NewSphere()
	s2.SetTransform(r.NewTranslation(-2.5, 0, 0))
	m3 := r.NewMetal(r.NewColor(0.5, 0.8, 0.8), 0.2)
	s2.NewMaterial = m3
	g.AddChild(s2)

	s3 := r.NewSphere()
	s3.SetTransform(r.NewTranslation(2.5, 0, 0))
	m4 := r.NewDielectric(1.5)
	s3.NewMaterial = m4
	g.AddChild(s3)

	g.Divide(1)

	w.AddObject(g)

	return w, r.ViewTransform(r.NewPoint(0, 4, -10), r.NewPoint(0, 0, 0), r.NewVec(0, 1, 0))
}

func startProfiling() func() {
	f, err := os.Create("profile-cpu.pb.gz")
	if err != nil {
		log.Fatal(err)
	}

	f2, err := os.Create("profile-mem.pb.gz")
	if err != nil {
		log.Fatal(err)
	}

	err = pprof.StartCPUProfile(f)
	if err != nil {
		log.Fatal(err)
	}

	return func() {
		pprof.StopCPUProfile()
		pprof.WriteHeapProfile(f2)
		f.Close()
		f2.Close()
	}
}

func main() {
	stop := startProfiling()
	defer stop()

	// Set up canvas
	_, ct := GetTestScene8()

	width := 1280
	ratio := 16.0 / 9.0

	camera := r.NewCamera(width, int(float64(width)/ratio), math.Pi/3).SetTransform(ct)
	camera.Samples = 1000
	camera.Depth = 10
	camera.GammaCorrection = true

	timeBefore := time.Now()

	canvas := camera.RenderMultiThreaded(GetTestScene8, runtime.NumCPU())

	timeAfter := time.Now()

	diff := timeAfter.Sub(timeBefore)

	fmt.Println("Render time:", diff)
	filename := fmt.Sprintf("render-%d-%d-%d-%s.png", time.Now().UnixMilli(), camera.Samples, camera.Depth, diff)
	fmt.Println("Saved:", filename)

	canvas.SavePNG(filename)
}

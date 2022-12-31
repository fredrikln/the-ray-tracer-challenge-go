package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	r "github.com/fredrikln/the-ray-tracer-challenge-go/pkg/raytracer"
)

func GetTestScene1() (*r.World, *r.Matrix) {
	w := r.NewWorld()

	g := r.NewGroup()

	c := r.NewColor(0, 0, 0)
	w.Background = &c

	// -- Plane as floor ---
	// floor := r.NewPlane()
	// floor.NewMaterial = r.NewDiffuse(r.NewColor(0.7, 0.8, 0.7))
	// floor.SetTransform(r.NewTranslation(0, -1, 0))
	// g.AddChild(floor)

	// --- Floor of boxes ---
	floor := r.NewGroup()

	m := r.NewDiffuse(r.NewColor(0.7, 0.8, 0.7))
	localSource := rand.New(rand.NewSource(1338))
	count := 20
	for i := 0; i < count; i++ {
		for j := 0; j < count; j++ {
			b := r.NewCube()
			b.SetNewMaterial(m)

			// fmt.Println(b.GetParentObject())

			b.SetTransform(r.NewTranslation(float64(i)*2, -1*localSource.Float64(), float64(j)*2).Translate(-float64(count), -2, -float64(count)))

			floor.AddChild(b)
		}
	}
	floor.SetTransform(r.NewScaling(2, 1, 2).Translate(0, -0.25, 0))
	g.AddChild(floor)

	// --- Walls ---
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

	// --- Light sphere ---
	light := r.NewCube()
	light.SetNewMaterial(r.NewEmissive(r.NewColor(1, 1, 1)))
	// light.NewMaterial = r.NewDiffuse(r.NewColor(0.5, 0.5, 0.3))
	light.SetTransform(r.NewTranslation(0, 10, 0).Scale(10, 0.01, 10))
	g.AddChild(light)

	// --- Dragon ---
	// content, err := os.ReadFile("models/dragon.obj")

	// if err != nil {
	// 	panic(err)
	// }

	// p := objparser.NewParser()
	// p.SetNewMaterial(r.NewMetal(r.NewColor(0.135, 0.2225, 0.1575), 0.75))
	// g2 := p.Parse(strings.Trim(string(content), "\n"))
	// t1 := r.NewTranslation(0, -2.25, 0)
	// g2.SetTransform(t1)

	// floor.SetTransform(floor.GetTransform().Translate(0, -1, 0))

	// g.AddChild(g2)

	// --- Spheres ---
	s1 := r.NewSphere()
	m2 := r.NewDiffuse(r.NewColor(0.8, 0.5, 0.5))
	s1.SetNewMaterial(m2)
	g.AddChild(s1)

	s2 := r.NewSphere()
	s2.SetTransform(r.NewTranslation(-2.5, 0, 0))
	m3 := r.NewMetal(r.NewColor(0.5, 0.8, 0.8), 0.2)
	s2.SetNewMaterial(m3)
	g.AddChild(s2)

	s3 := r.NewSphere()
	s3.SetTransform(r.NewTranslation(2.5, 0, 0))
	m4 := r.NewDielectric(1.5)
	s3.SetNewMaterial(m4)
	g.AddChild(s3)

	g.Divide(1)

	w.AddObject(g)

	return w, r.ViewTransform(r.NewPoint(0, 5, -10), r.NewPoint(0, 0, 0), r.NewVec(0, 1, 0))
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
	_, ct := GetTestScene1()

	width := 400
	// ratio := 43.0 / 18.0
	ratio := 16.0 / 9.0
	// ratio := 1.0

	camera := r.NewCamera(width, int(float64(width)/ratio), math.Pi/3).SetTransform(ct)
	camera.Samples = 10
	camera.Depth = 10
	camera.GammaCorrection = true

	timeBefore := time.Now()

	canvas := camera.RenderMultiThreaded(GetTestScene1, runtime.NumCPU())

	timeAfter := time.Now()

	diff := timeAfter.Sub(timeBefore)

	fmt.Println("Render time:", diff)
	filename := fmt.Sprintf("render-%d-%d-%d-%s.png", time.Now().UnixMilli(), camera.Samples, camera.Depth, diff)
	fmt.Println("Saved:", filename)

	canvas.SavePNG(filename)
}

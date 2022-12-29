package raytracer

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type Camera struct {
	Hsize             int
	Vsize             int
	Fov               float64
	Transform         *Matrix
	PixelSize         float64
	HalfWidth         float64
	HalfHeight        float64
	Antialiasing      bool
	AntiAliasingSteps int
	Bounces           int
}

func NewCamera(hsize, vsize int, fov float64) *Camera {
	halfView := math.Tan(fov / 2)
	aspect := float64(hsize) / float64(vsize)

	var halfWidth, halfHeight float64

	if aspect > 1 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfWidth = halfView * aspect
		halfHeight = halfView
	}

	return &Camera{
		Hsize:             hsize,
		Vsize:             vsize,
		Fov:               fov,
		Transform:         NewIdentityMatrix(),
		PixelSize:         (halfWidth * 2) / float64(hsize),
		HalfWidth:         halfWidth,
		HalfHeight:        halfHeight,
		Antialiasing:      false,
		AntiAliasingSteps: 2,
		Bounces:           4,
	}
}

func (c *Camera) RayForPixel(x, y float64) Ray {
	xOffset := x * c.PixelSize
	yOffset := y * c.PixelSize

	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset

	pixel := c.Transform.Inverse().MulPoint(NewPoint(worldX, worldY, -1))
	origin := c.Transform.Inverse().MulPoint(NewPoint(0, 0, 0))

	direction := pixel.Sub(origin).Norm()

	return NewRay(origin, direction)
}

func (c *Camera) SetTransform(m *Matrix) *Camera {
	c.Transform = m

	return c
}

func (c *Camera) Render(w *World) *Canvas {
	canvas := NewCanvas(c.Hsize, c.Vsize)

	var linesRendered int
	start := time.Now()
	prev := start

	for y := 0; y < c.Vsize; y++ {
		for x := 0; x < c.Hsize; x++ {
			color := c.getColorForPixels(float64(x), float64(y), w)

			canvas.SetPixel(x, y, color)
		}

		linesRendered++

		if time.Since(prev).Seconds() > 1 {
			elapsed := time.Since(start)

			linesLeft := c.Vsize - linesRendered
			avgTimePerLine := (elapsed.Seconds() / float64(linesRendered))

			timeLeft := avgTimePerLine * float64(linesLeft)

			fmt.Printf(
				"Progress: %d of %d (%.2f%%) (%.0fs/%.0fs)",
				linesRendered,
				c.Vsize,
				float64(linesRendered)/float64(c.Vsize)*100,
				elapsed.Seconds(),
				timeLeft+elapsed.Seconds(),
			)
			fmt.Println()

			prev = time.Now()
		}

	}

	return canvas
}

func (c *Camera) getColorForPixels(x, y float64, w *World) Color {
	x = x + 0.5
	y = y + 0.5
	if !c.Antialiasing {
		ray := c.RayForPixel(x, y)
		color := w.ColorAt(ray, c.Bounces)

		return color
	}

	var outColor Color

	var steps float64 = float64(c.AntiAliasingSteps)

	for i := 0.0; i < steps; i++ {
		for j := 0.0; j < steps; j++ {
			offsetX := (1 / (steps * 2)) + i*1/steps
			offsetY := (1 / (steps * 2)) + j*1/steps

			ray := c.RayForPixel(x+offsetX, y+offsetY)
			color := w.ColorAt(ray, c.Bounces)

			outColor = outColor.Add(color)
		}
	}

	// gamma correction
	scale := 1.0 / (steps * steps)
	outColor = NewColor(math.Sqrt(outColor.R*scale), math.Sqrt(outColor.G*scale), math.Sqrt(outColor.B*scale))

	return outColor
}

type response struct {
	Y    int
	line []Color
}

type job struct {
	Y      int
	Camera *Camera
	World  *World
}

func worker(jobChan chan job, responseChan chan response, wg *sync.WaitGroup) {
	for job := range jobChan {
		y := job.Y
		line := make([]Color, 0, job.Camera.Hsize)

		for x := 0; x < job.Camera.Hsize; x++ {
			color := job.Camera.getColorForPixels(float64(x), float64(y), job.World)

			line = append(line, color)
		}
		responseChan <- response{y, line}

		wg.Done()
	}
}

func (c *Camera) RenderMultiThreaded(w *World, cores int) *Canvas {
	canvas := NewCanvas(c.Hsize, c.Vsize)

	jobChan := make(chan job)
	responseChan := make(chan response)

	wg := sync.WaitGroup{}
	wg.Add(c.Vsize)

	defer close(jobChan)
	defer close(responseChan)

	// Handle responses
	go func() {
		var linesRendered int
		start := time.Now()
		prev := start

		for response := range responseChan {
			for x := 0; x < c.Hsize; x++ {
				canvas.SetPixel(x, response.Y, response.line[x])
			}

			linesRendered++

			if time.Since(prev).Seconds() > 5 {
				elapsed := time.Since(start)

				linesLeft := c.Vsize - linesRendered
				avgTimePerLine := (elapsed.Seconds() / float64(linesRendered))

				timeLeft := avgTimePerLine * float64(linesLeft)

				fmt.Printf(
					"Progress: %d of %d (%.2f%%) (%.0fs/%.0fs)",
					linesRendered,
					c.Vsize,
					float64(linesRendered)/float64(c.Vsize)*100,
					elapsed.Seconds(),
					timeLeft+elapsed.Seconds(),
				)
				fmt.Println()

				prev = time.Now()
			}
		}
	}()

	// Start workers
	for i := 0; i < cores; i++ {
		go worker(jobChan, responseChan, &wg)
	}

	// Send jobs
	go func() {
		for y := 0; y < c.Vsize; y++ {
			jobChan <- job{y, c, w}
		}
	}()

	wg.Wait()

	return canvas
}

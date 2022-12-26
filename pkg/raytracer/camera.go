package raytracer

import (
	"math"
	"sync"
)

type Camera struct {
	Hsize        int
	Vsize        int
	Fov          float64
	Transform    *Matrix
	PixelSize    float64
	HalfWidth    float64
	HalfHeight   float64
	Antialiasing bool
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
		Hsize:        hsize,
		Vsize:        vsize,
		Fov:          fov,
		Transform:    NewIdentityMatrix(),
		PixelSize:    (halfWidth * 2) / float64(hsize),
		HalfWidth:    halfWidth,
		HalfHeight:   halfHeight,
		Antialiasing: false,
	}
}

func (c *Camera) RayForPixel(x, y, offsetX, offsetY float64) Ray {
	xOffset := (x + offsetX) * c.PixelSize
	yOffset := (y + offsetY) * c.PixelSize

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

	for y := 0; y < c.Vsize; y++ {
		for x := 0; x < c.Hsize; x++ {
			color := c.getColorForPixels(float64(x), float64(y), w)

			canvas.SetPixel(x, y, color)
		}
	}

	return canvas
}

func (c *Camera) getColorForPixels(x, y float64, w *World) Color {
	if !c.Antialiasing {
		ray := c.RayForPixel(float64(x), float64(y), 0.5, 0.5)
		color := w.ColorAt(ray)

		return color
	}

	var outColor Color

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			offsetX := 0.25 + float64(i)*0.5
			offsetY := 0.25 + float64(j)*0.5
			ray := c.RayForPixel(float64(x), float64(y), offsetX, offsetY)
			color := w.ColorAt(ray)

			outColor = outColor.Add(color)
		}
	}

	outColor = outColor.MulFloat(0.25)

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
		for response := range responseChan {
			for x := 0; x < c.Hsize; x++ {
				canvas.SetPixel(x, response.Y, response.line[x])
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

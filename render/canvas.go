package render

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
)

type Canvas struct {
	Width  int
	Height int
	Pixels []Color
}

func NewCanvas(width, height int) Canvas {
	pixels := make([]Color, width*height)

	return Canvas{
		width, height, pixels,
	}
}

func (c *Canvas) GetPixel(x, y int) Color {
	return c.Pixels[y*c.Width+x]
}

func (c *Canvas) SetPixel(x, y int, color Color) {
	c.Pixels[y*c.Width+x] = color
}

func getColorValue(c float64) int {
	return int(math.Min(math.Max(math.Round(c*255), 0), 255))
}

func (c *Canvas) getPPMString() string {
	data := "P3\n"
	data += fmt.Sprintf("%d %d\n", c.Width, c.Height)
	data += "255\n"

	for y := 0; y < c.Height; y += 1 {
		pixelValues := make([]string, 0)

		for x := 0; x < c.Width; x += 1 {
			pixel := c.GetPixel(x, y)
			r := getColorValue(pixel.R)
			g := getColorValue(pixel.G)
			b := getColorValue(pixel.B)

			pixelValues = append(pixelValues, fmt.Sprint(r))
			pixelValues = append(pixelValues, fmt.Sprint(g))
			pixelValues = append(pixelValues, fmt.Sprint(b))
		}

		line := ""

		for _, value := range pixelValues {
			if len(line)+len(value) >= 70 {
				data += line + "\n"
				line = ""
			}

			if line != "" {
				line += " "
			}

			line += value
		}

		data += line + "\n"
	}

	return data
}

func (c *Canvas) SavePPM(filename string) {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	data := c.getPPMString()

	bytes, err := f.Write([]byte(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(bytes, "bytes written successfully")

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (c *Canvas) SavePNG(filename string) {
	img := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))

	for y := 0; y < c.Height; y += 1 {
		for x := 0; x < c.Width; x += 1 {
			pixel := c.GetPixel(x, y)
			color := pixel.GetRGBA()

			img.Set(x, y, color)
		}
	}

	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	png.Encode(f, img)
}

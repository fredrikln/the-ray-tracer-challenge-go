package raytracer

import (
	"math"
	"testing"
)

func TestCamera(t *testing.T) {
	camera := NewCamera(160, 120, math.Pi/2)

	if camera.Hsize != 160 {
		t.Errorf("Wrong camera hsize, got %v, want %v", camera.Hsize, 160)
	}

	if camera.Vsize != 120 {
		t.Errorf("Wrong camera vsize, got %v, want %v", camera.Vsize, 120)
	}

	if camera.Fov != math.Pi/2 {
		t.Errorf("Wrong camera fov, got %v, want %v", camera.Fov, math.Pi/2)
	}

	if !camera.Transform.Eq(NewIdentityMatrix()) {
		t.Errorf("Wrong camera default transform, got %v, want %v", camera.Transform, NewIdentityMatrix())
	}
}

func TestCameraPixelSize(t *testing.T) {
	camera := NewCamera(200, 125, math.Pi/2)

	if camera.PixelSize != 0.01 {
		t.Errorf("Got wrong pixel size, got %v, want %v", camera.PixelSize, 0.01)
	}

	camera2 := NewCamera(125, 200, math.Pi/2)

	if camera2.PixelSize != 0.01 {
		t.Errorf("Got wrong pixel size, got %v, want %v", camera2.PixelSize, 0.01)
	}
}

func TestRayThroughCanvas(t *testing.T) {
	testCases := []struct {
		desc   string
		camera *Camera
		x      float64
		y      float64
		want   Ray
	}{
		{
			desc:   "Center of canvas",
			camera: NewCamera(201, 101, math.Pi/2),
			x:      100,
			y:      50,
			want:   NewRay(NewPoint(0, 0, 0), NewVec(0, 0, -1)),
		},
		{
			desc:   "Corner of canvas",
			camera: NewCamera(201, 101, math.Pi/2),
			x:      0,
			y:      0,
			want:   NewRay(NewPoint(0, 0, 0), NewVec(0.66519, 0.33259, -0.66851)),
		},
		{
			desc:   "Ray when camera transformed",
			camera: NewCamera(201, 101, math.Pi/2).SetTransform(NewRotationY(math.Pi / 4).Mul(NewTranslation(0, -2, 5))),
			x:      100,
			y:      50,
			want:   NewRay(NewPoint(0, 2, -5), NewVec(math.Sqrt(2)/2, 0, -(math.Sqrt(2)/2))),
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			ray := (*tC.camera).RayForPixel(tC.x, tC.y, 0.5, 0.5)

			if !ray.Direction.Eq(tC.want.Direction) || !ray.Origin.Eq(tC.want.Origin) {
				t.Errorf("Got %v, want %v", ray, tC.want)
			}
		})
	}
}

func TestRenderWorldWithCamera(t *testing.T) {
	w := NewDefaultWorld()
	c := NewCamera(11, 11, math.Pi/2)

	from := NewPoint(0, 0, -5)
	to := NewPoint(0, 0, 0)
	up := NewVec(0, 1, 0)

	c.SetTransform(ViewTransform(from, to, up))

	canvas := c.Render(w)
	got := canvas.GetPixel(5, 5)
	want := NewColor(0.380661, 0.475826, 0.285495)

	if !got.Eq(want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}

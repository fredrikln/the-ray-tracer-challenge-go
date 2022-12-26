package raytracer

type Light interface {
	GetIntensity() Color
	GetPosition() Point
}

type PointLight struct {
	Intensity Color
	Position  Point
}

func NewPointLight(position Point, intensity Color) *PointLight {
	return &PointLight{
		intensity,
		position,
	}
}

func (pl *PointLight) GetIntensity() Color {
	return pl.Intensity
}

func (pl *PointLight) GetPosition() Point {
	return pl.Position
}

package geom_test

import (
	"testing"

	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
)

func TestNewRay(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		origin := g.NewPoint(1.0, 2.0, 3.0)
		direction := g.NewVec(4, 5, 6)
		ray := g.NewRay(origin, direction)

		if !ray.Origin.Eq(origin) || !ray.Direction.Eq(direction) {
			t.Error("Ray initialized incorrectly")
		}
	})
}

func TestRayPosition(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		tests := []struct {
			name string
			a    g.Ray
			time float64
			want g.Point
		}{
			{
				"Test 1",
				g.NewRay(g.NewPoint(2, 3, 4), g.NewVec(1, 0, 0)),
				0,
				g.NewPoint(2, 3, 4),
			},
			{
				"Test 2",
				g.NewRay(g.NewPoint(2, 3, 4), g.NewVec(1, 0, 0)),
				1,
				g.NewPoint(3, 3, 4),
			},
			{
				"Test 3",
				g.NewRay(g.NewPoint(2, 3, 4), g.NewVec(1, 0, 0)),
				-1,
				g.NewPoint(1, 3, 4),
			},
			{
				"Test 4",
				g.NewRay(g.NewPoint(2, 3, 4), g.NewVec(1, 0, 0)),
				2.5,
				g.NewPoint(4.5, 3, 4),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if pos := tt.a.Position(tt.time); !pos.Eq(tt.want) {
					t.Errorf("Got %v, Want %v", pos, tt.want)
				}
			})
		}
	})
}

func TestRayTranslate(t *testing.T) {
	r := g.NewRay(g.NewPoint(1, 2, 3), g.NewVec(0, 1, 0))
	m := g.NewTranslation(3, 4, 5)

	r2 := r.Mul(m)

	if !r2.Origin.Eq(g.NewPoint(4, 6, 8)) || !r2.Direction.Eq(g.NewVec(0, 1, 0)) {
		t.Errorf("Invalid ray translation, got %v %v, want %v, %v", r2.Origin, r2.Direction, g.NewPoint(4, 6, 8), g.NewVec(0, 1, 0))
	}
}

func TestRayScale(t *testing.T) {
	r := g.NewRay(g.NewPoint(1, 2, 3), g.NewVec(0, 1, 0))
	m := g.NewScaling(2, 3, 4)

	r2 := r.Mul(m)

	if !r2.Origin.Eq(g.NewPoint(2, 6, 12)) || !r2.Direction.Eq(g.NewVec(0, 3, 0)) {
		t.Errorf("Invalid ray scaling, Got %v %v, Want %v %v", r2.Origin, r2.Direction, g.NewPoint(2, 6, 12), g.NewVec(0, 3, 0))
	}
}

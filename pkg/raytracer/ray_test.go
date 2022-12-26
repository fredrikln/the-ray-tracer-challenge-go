package raytracer

import (
	"testing"
)

func TestNewRay(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		origin := NewPoint(1.0, 2.0, 3.0)
		direction := NewVec(4, 5, 6)
		ray := NewRay(origin, direction)

		if !ray.Origin.Eq(origin) || !ray.Direction.Eq(direction) {
			t.Error("Ray initialized incorrectly")
		}
	})
}

func TestRayPosition(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		tests := []struct {
			name string
			a    Ray
			time float64
			want Point
		}{
			{
				"Test 1",
				NewRay(NewPoint(2, 3, 4), NewVec(1, 0, 0)),
				0,
				NewPoint(2, 3, 4),
			},
			{
				"Test 2",
				NewRay(NewPoint(2, 3, 4), NewVec(1, 0, 0)),
				1,
				NewPoint(3, 3, 4),
			},
			{
				"Test 3",
				NewRay(NewPoint(2, 3, 4), NewVec(1, 0, 0)),
				-1,
				NewPoint(1, 3, 4),
			},
			{
				"Test 4",
				NewRay(NewPoint(2, 3, 4), NewVec(1, 0, 0)),
				2.5,
				NewPoint(4.5, 3, 4),
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
	r := NewRay(NewPoint(1, 2, 3), NewVec(0, 1, 0))
	m := NewTranslation(3, 4, 5)

	r2 := r.Mul(m)

	if !r2.Origin.Eq(NewPoint(4, 6, 8)) || !r2.Direction.Eq(NewVec(0, 1, 0)) {
		t.Errorf("Invalid ray translation, got %v %v, want %v, %v", r2.Origin, r2.Direction, NewPoint(4, 6, 8), NewVec(0, 1, 0))
	}
}

func TestRayScale(t *testing.T) {
	r := NewRay(NewPoint(1, 2, 3), NewVec(0, 1, 0))
	m := NewScaling(2, 3, 4)

	r2 := r.Mul(m)

	if !r2.Origin.Eq(NewPoint(2, 6, 12)) || !r2.Direction.Eq(NewVec(0, 3, 0)) {
		t.Errorf("Invalid ray scaling, Got %v %v, Want %v %v", r2.Origin, r2.Direction, NewPoint(2, 6, 12), NewVec(0, 3, 0))
	}
}

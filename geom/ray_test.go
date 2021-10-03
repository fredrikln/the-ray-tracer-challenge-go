package geom

import "testing"

func TestNewRay(t *testing.T) {
	t.Run("Test", func(t *testing.T) {
		origin := NewVec(1.0, 2.0, 3.0)
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
			want Vec
		}{
			{
				"Test 1",
				NewRay(NewVec(2, 3, 4), NewVec(1, 0, 0)),
				0,
				NewVec(2, 3, 4),
			},
			{
				"Test 2",
				NewRay(NewVec(2, 3, 4), NewVec(1, 0, 0)),
				1,
				NewVec(3, 3, 4),
			},
			{
				"Test 3",
				NewRay(NewVec(2, 3, 4), NewVec(1, 0, 0)),
				-1,
				NewVec(1, 3, 4),
			},
			{
				"Test 4",
				NewRay(NewVec(2, 3, 4), NewVec(1, 0, 0)),
				2.5,
				NewVec(4.5, 3, 4),
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

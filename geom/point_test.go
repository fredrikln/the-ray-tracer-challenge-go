package geom_test

import (
	"testing"

	g "github.com/fredrikln/the-ray-tracer-challenge-go/geom"
)

func TestPointMulMat(t *testing.T) {
	tests := []struct {
		name string
		a    g.Point
		b    *g.Matrix
		want g.Point
	}{
		{"Test 1",
			g.NewPoint(1, 2, 3),
			g.NewMatrix(
				1, 2, 3, 4,
				2, 4, 4, 2,
				8, 6, 4, 1,
				0, 0, 0, 1,
			),
			g.NewPoint(18, 24, 33),
		},
		{"Test 2",
			g.NewPoint(1, 2, 3),
			g.NewIdentityMatrix(),
			g.NewPoint(1, 2, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.MulMat(tt.b); !got.Eq(tt.want) {
				t.Errorf("Got %v, want %v", got, tt.want)
			}
		})
	}
}

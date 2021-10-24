package geom

import (
	"testing"
)

func TestPointMulMat(t *testing.T) {
	tests := []struct {
		name string
		a    Point
		b    *Matrix
		want Point
	}{
		{"Test 1",
			NewPoint(1, 2, 3),
			NewMatrix(
				1, 2, 3, 4,
				2, 4, 4, 2,
				8, 6, 4, 1,
				0, 0, 0, 1,
			),
			NewPoint(18, 24, 33),
		},
		{"Test 2",
			NewPoint(1, 2, 3),
			NewIdentityMatrix(),
			NewPoint(1, 2, 3),
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

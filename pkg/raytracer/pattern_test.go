package raytracer

import (
	"testing"
)

var black, white Color = NewColor(0, 0, 0), NewColor(1, 1, 1)

func TestNewStripePattern(t *testing.T) {
	p := NewStripePattern(white, black)

	if !p.A.Eq(white) {
		t.Error("Invalid A color")
	}

	if !p.B.Eq(black) {
		t.Error("Invalid B color")
	}
}

func TestStripePatternConstantInY(t *testing.T) {
	p := NewStripePattern(white, black)

	if !p.ColorAt(NewPoint(0, 0, 0)).Eq(white) {
		t.Error("Invalid color at 0,0,0")
	}
	if !p.ColorAt(NewPoint(0, 1, 0)).Eq(white) {
		t.Error("Invalid color at 0,1,0")
	}
	if !p.ColorAt(NewPoint(0, 2, 0)).Eq(white) {
		t.Error("Invalid color at 0,2,0")
	}
}

func TestStripePatternConstantInZ(t *testing.T) {
	p := NewStripePattern(white, black)

	if !p.ColorAt(NewPoint(0, 0, 0)).Eq(white) {
		t.Error("Invalid color at 0,0,0")
	}
	if !p.ColorAt(NewPoint(0, 0, 1)).Eq(white) {
		t.Error("Invalid color at 0,0,1")
	}
	if !p.ColorAt(NewPoint(0, 0, 2)).Eq(white) {
		t.Error("Invalid color at 0,0,2")
	}
}

func TestStripePatternAlternatesInX(t *testing.T) {
	p := NewStripePattern(white, black)

	if !p.ColorAt(NewPoint(0, 0, 0)).Eq(white) {
		t.Error("Invalid color at 0,0,0")
	}
	if !p.ColorAt(NewPoint(0.9, 0, 0)).Eq(white) {
		t.Error("Invalid color at 0.9,0,0")
	}
	if !p.ColorAt(NewPoint(1, 0, 0)).Eq(black) {
		t.Error("Invalid color at 1,0,0")
	}
	if !p.ColorAt(NewPoint(-0.1, 0, 0)).Eq(black) {
		t.Error("Invalid color at -0.1,0,0")
	}
	if !p.ColorAt(NewPoint(-1, 0, 0)).Eq(black) {
		t.Error("Invalid color at -1,0,0")
	}
	if !p.ColorAt(NewPoint(-1.1, 0, 0)).Eq(white) {
		t.Error("Invalid color at -1.1,0,0")
	}
}

func TestStripesWithObjectTransformation(t *testing.T) {
	s := NewSphere().SetTransform(NewScaling(2, 2, 2))
	p := NewStripePattern(white, black)

	c := p.ColorAtObject(s, NewPoint(1.5, 0, 0))

	if !c.Eq(white) {
		t.Errorf("Invalid color, got %v, want %v", c, white)
	}
}

func TestStripesWithPatternTransformation(t *testing.T) {
	s := NewSphere()
	p := NewStripePattern(white, black)
	p.SetTransform(NewScaling(2, 2, 2))

	c := p.ColorAtObject(s, NewPoint(1.5, 0, 0))

	if !c.Eq(white) {
		t.Errorf("Invalid color, got %v, want %v", c, white)
	}
}

func TestStripesWithBothObjectAndPatternTransformation(t *testing.T) {
	s := NewSphere().SetTransform(NewScaling(2, 2, 2))
	p := NewStripePattern(white, black)
	p.SetTransform(NewScaling(2, 2, 2))

	c := p.ColorAtObject(s, NewPoint(1.5, 0, 0))

	if !c.Eq(white) {
		t.Errorf("Invalid color, got %v, want %v", c, white)
	}
}

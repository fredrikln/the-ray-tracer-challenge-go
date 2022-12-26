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

func TestGradientPattern(t *testing.T) {
	p := NewGradientPattern(white, black)

	c1 := p.ColorAt(NewPoint(0, 0, 0))
	c2 := p.ColorAt(NewPoint(0.25, 0, 0))
	c3 := p.ColorAt(NewPoint(0.5, 0, 0))
	c4 := p.ColorAt(NewPoint(0.75, 0, 0))

	if !c1.Eq(white) {
		t.Errorf("Invalid color c1, got %v want %v", c1, white)
	}
	if !c2.Eq(NewColor(0.75, 0.75, 0.75)) {
		t.Errorf("Invalid color c2, got %v want %v", c2, NewColor(0.75, 0.75, 0.75))
	}
	if !c3.Eq(NewColor(0.5, 0.5, 0.5)) {
		t.Errorf("Invalid color c3, got %v want %v", c3, NewColor(0.5, 0.5, 0.5))
	}
	if !c4.Eq(NewColor(0.25, 0.25, 0.25)) {
		t.Errorf("Invalid color c4, got %v want %v", c4, NewColor(0.25, 0.25, 0.25))
	}
}

func TestRingPattern(t *testing.T) {
	p := NewRingPattern(white, black)

	c1 := p.ColorAt(NewPoint(0, 0, 0))
	c2 := p.ColorAt(NewPoint(1, 0, 0))
	c3 := p.ColorAt(NewPoint(0, 0, 1))
	c4 := p.ColorAt(NewPoint(0.708, 0, 0.708))

	if !c1.Eq(white) {
		t.Errorf("Wrong color at c1, got %v, want %v", c1, white)
	}
	if !c2.Eq(black) {
		t.Errorf("Wrong color at c2, got %v, want %v", c2, black)
	}
	if !c3.Eq(black) {
		t.Errorf("Wrong color at c3, got %v, want %v", c3, black)
	}
	if !c4.Eq(black) {
		t.Errorf("Wrong color at c4, got %v, want %v", c4, black)
	}
}

func TestCheckerPatternRepeatX(t *testing.T) {
	p := NewCheckerPattern(white, black)

	c1 := p.ColorAt(NewPoint(0, 0, 0))
	c2 := p.ColorAt(NewPoint(0.99, 0, 0))
	c3 := p.ColorAt(NewPoint(1.01, 0, 0))

	if !c1.Eq(white) {
		t.Errorf("Wrong color at c1, got %v, want %v", c1, white)
	}
	if !c2.Eq(white) {
		t.Errorf("Wrong color at c2, got %v, want %v", c2, white)
	}
	if !c3.Eq(black) {
		t.Errorf("Wrong color at c3, got %v, want %v", c3, black)
	}
}

func TestCheckerPatternRepeatY(t *testing.T) {
	p := NewCheckerPattern(white, black)

	c1 := p.ColorAt(NewPoint(0, 0, 0))
	c2 := p.ColorAt(NewPoint(0, 0.99, 0))
	c3 := p.ColorAt(NewPoint(0, 1.01, 0))

	if !c1.Eq(white) {
		t.Errorf("Wrong color at c1, got %v, want %v", c1, white)
	}
	if !c2.Eq(white) {
		t.Errorf("Wrong color at c2, got %v, want %v", c2, white)
	}
	if !c3.Eq(black) {
		t.Errorf("Wrong color at c3, got %v, want %v", c3, black)
	}
}

func TestCheckerPatternRepeatZ(t *testing.T) {
	p := NewCheckerPattern(white, black)

	c1 := p.ColorAt(NewPoint(0, 0, 0))
	c2 := p.ColorAt(NewPoint(0, 0, 0.99))
	c3 := p.ColorAt(NewPoint(0, 0, 1.01))

	if !c1.Eq(white) {
		t.Errorf("Wrong color at c1, got %v, want %v", c1, white)
	}
	if !c2.Eq(white) {
		t.Errorf("Wrong color at c2, got %v, want %v", c2, white)
	}
	if !c3.Eq(black) {
		t.Errorf("Wrong color at c3, got %v, want %v", c3, black)
	}
}

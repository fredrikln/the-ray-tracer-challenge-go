package objparser

import (
	"testing"

	r "github.com/fredrikln/the-ray-tracer-challenge-go/pkg/raytracer"
)

func TestIgnoreGibberish(t *testing.T) {
	parser := NewParser()

	input := `There was a young lady named Bright
	who traveled much faster than light.
	She set out one day
	in a relative way,
	and came back the previous night.`

	parser.Parse(input)

	if parser.numIgnoredLines != 5 {
		t.Errorf("Got %v, want %v", parser.numIgnoredLines, 5)
	}
}

func TestParseVertexRecords(t *testing.T) {
	parser := NewParser()

	input := `v -1 1 0
	v -1.0000 0.5000 0.0000
	v 1 0 0
	v 1 1 0`

	parser.Parse(input)

	if len(parser.vertices) != 5 {
		t.Errorf("Invalid vertices length, got %v, want %v", len(parser.vertices), 4)
	}

	if !parser.vertices[1].Eq(r.NewPoint(-1, 1, 0)) {
		t.Errorf("Invalid vertex 1, got %v, want %v", parser.vertices[1], r.NewPoint(-1, 1, 0))
	}

	if !parser.vertices[2].Eq(r.NewPoint(-1, 0.5, 0)) {
		t.Errorf("Invalid vertex 1, got %v, want %v", parser.vertices[2], r.NewPoint(-1, 0.5, 0))
	}

	if !parser.vertices[3].Eq(r.NewPoint(1, 0, 0)) {
		t.Errorf("Invalid vertex 1, got %v, want %v", parser.vertices[3], r.NewPoint(1, 0, 0))
	}

	if !parser.vertices[4].Eq(r.NewPoint(1, 1, 0)) {
		t.Errorf("Invalid vertex 1, got %v, want %v", parser.vertices[4], r.NewPoint(1, 1, 0))
	}
}

func TestParseTriangleFaces(t *testing.T) {
	parser := NewParser()

	input := `v -1 1 0
	v -1 0 0
	v 1 0 0
	v 1 1 0

	f 1 2 3
	f 1 3 4`

	parser.Parse(input)

	g := parser.defaultGroup

	t1 := g.Items[0]
	t2 := g.Items[1]
	if !t1.(*r.Triangle).P1.Eq(parser.vertices[1]) {
		t.Error("T1P1 wrong vertex")
	}
	if !t1.(*r.Triangle).P2.Eq(parser.vertices[2]) {
		t.Error("T1P2 wrong vertex")
	}
	if !t1.(*r.Triangle).P3.Eq(parser.vertices[3]) {
		t.Error("T1P3 wrong vertex")
	}
	if !t2.(*r.Triangle).P1.Eq(parser.vertices[1]) {
		t.Error("T2P1 wrong vertex")
	}
	if !t2.(*r.Triangle).P2.Eq(parser.vertices[3]) {
		t.Error("T2P2 wrong vertex")
	}
	if !t2.(*r.Triangle).P3.Eq(parser.vertices[4]) {
		t.Error("T2P3 wrong vertex")
	}
}

func TestParsePolygonData(t *testing.T) {
	parser := NewParser()

	input := `v -1 1 0
	v -1 0 0
	v 1 0 0
	v 1 1 0
	v 0 2 0

	f 1 2 3 4 5`

	parser.Parse(input)

	g := parser.defaultGroup

	t1 := g.Items[0]
	t2 := g.Items[1]
	t3 := g.Items[2]

	if !t1.(*r.Triangle).P1.Eq(parser.vertices[1]) {
		t.Error("T1P1 wrong vertex")
	}
	if !t1.(*r.Triangle).P2.Eq(parser.vertices[2]) {
		t.Error("T1P2 wrong vertex")
	}
	if !t1.(*r.Triangle).P3.Eq(parser.vertices[3]) {
		t.Error("T1P3 wrong vertex")
	}
	if !t2.(*r.Triangle).P1.Eq(parser.vertices[1]) {
		t.Error("T2P1 wrong vertex")
	}
	if !t2.(*r.Triangle).P2.Eq(parser.vertices[3]) {
		t.Error("T2P2 wrong vertex")
	}
	if !t2.(*r.Triangle).P3.Eq(parser.vertices[4]) {
		t.Error("T2P3 wrong vertex")
	}
	if !t3.(*r.Triangle).P1.Eq(parser.vertices[1]) {
		t.Error("T3P1 wrong vertex")
	}
	if !t3.(*r.Triangle).P2.Eq(parser.vertices[4]) {
		t.Error("T3P2 wrong vertex")
	}
	if !t3.(*r.Triangle).P3.Eq(parser.vertices[5]) {
		t.Error("T3P3 wrong vertex")
	}
}

func TestNamedGroups(t *testing.T) {
	parser := NewParser()

	input := `v -1 1 0
	v -1 0 0
	v 1 0 0
	v 1 1 0

	g FirstGroup
	f 1 2 3

	g SecondGroup
	f 1 3 4`

	g := parser.Parse(input)

	if len(g.Items) != 2 {
		t.Errorf("Invalid main group length, got %v, want %v", len(g.Items), 2)
	}

	g1 := g.Items[0]
	t1 := g1.(*r.Group).Items[0]

	g2 := g.Items[1]
	t2 := g2.(*r.Group).Items[0]

	if !t1.(*r.Triangle).P1.Eq(parser.vertices[1]) {
		t.Error("T1P1 wrong vertex")
	}
	if !t1.(*r.Triangle).P2.Eq(parser.vertices[2]) {
		t.Error("T1P2 wrong vertex")
	}
	if !t1.(*r.Triangle).P3.Eq(parser.vertices[3]) {
		t.Error("T1P3 wrong vertex")
	}
	if !t2.(*r.Triangle).P1.Eq(parser.vertices[1]) {
		t.Error("T2P1 wrong vertex")
	}
	if !t2.(*r.Triangle).P2.Eq(parser.vertices[3]) {
		t.Error("T2P2 wrong vertex")
	}
	if !t2.(*r.Triangle).P3.Eq(parser.vertices[4]) {
		t.Error("T2P3 wrong vertex")
	}
}

func TestSmoothTrianglesInObj(t *testing.T) {
	parser := NewParser()

	input := `vn 0 0 1
	vn 0.707 0 -0.707
	vn 1 2 3`

	parser.Parse(input)

	if len(parser.normals) != 4 {
		t.Error("Wrong normal count")
	}

	if !parser.normals[1].Eq(r.NewVec(0, 0, 1)) {
		t.Errorf("Invalid normal 1, got %v, want %v", parser.normals[1], r.NewVec(0, 0, 1))
	}
	if !parser.normals[2].Eq(r.NewVec(0.707, 0, -0.707)) {
		t.Errorf("Invalid normal 1, got %v, want %v", parser.normals[2], r.NewVec(0.707, 0, -0.707))
	}
	if !parser.normals[3].Eq(r.NewVec(1, 2, 3)) {
		t.Errorf("Invalid normal 1, got %v, want %v", parser.normals[3], r.NewVec(1, 2, 3))
	}
}

func TestFacesWithNormals(t *testing.T) {
	parser := NewParser()

	input := `v 0 1 0
	v -1 0 0
	v 1 0 0

	vn -1 0 0
	vn 1 0 0
	vn 0 1 0

	f 1//3 2//1 3//2
	f 1/0/3 2/102/1 3/14/2`

	g := parser.Parse(input)

	t1 := g.Items[0]
	//t2 := g.Items[1]

	if !t1.(*r.SmoothTriangle).P1.Eq(parser.vertices[1]) {
		t.Error("T1P1 wrong vertex")
	}
	if !t1.(*r.SmoothTriangle).P2.Eq(parser.vertices[2]) {
		t.Error("T1P2 wrong vertex")
	}
	if !t1.(*r.SmoothTriangle).P3.Eq(parser.vertices[3]) {
		t.Error("T1P3 wrong vertex")
	}
	if !t1.(*r.SmoothTriangle).N1.Eq(parser.normals[3]) {
		t.Error("T1N1 wrong normal")
	}
	if !t1.(*r.SmoothTriangle).N2.Eq(parser.normals[1]) {
		t.Error("T1N2 wrong normal")
	}
	if !t1.(*r.SmoothTriangle).N3.Eq(parser.normals[2]) {
		t.Error("T1N3 wrong normal")
	}
}

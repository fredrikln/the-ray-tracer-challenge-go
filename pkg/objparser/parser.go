package objparser

import (
	"fmt"
	"strconv"
	"strings"

	r "github.com/fredrikln/the-ray-tracer-challenge-go/pkg/raytracer"
)

type Objparser struct {
	vertices        []r.Point
	normals         []r.Vec
	numIgnoredLines int
	defaultGroup    *r.Group
	material        *r.Material
	NewMaterial     r.Scatters
}

func NewParser() *Objparser {
	return &Objparser{
		vertices:        make([]r.Point, 1),
		normals:         make([]r.Vec, 1),
		numIgnoredLines: 0,
		defaultGroup:    r.NewGroup(),
	}
}

func (p *Objparser) SetMaterial(m *r.Material) *Objparser {
	p.material = m

	return p
}

func (p *Objparser) Parse(input string) *r.Group {
	lines := strings.Split(input, "\n")

	var curGroup *r.Group

	for _, line := range lines {
		line = strings.TrimSpace(line)

		fields := strings.Fields(line)

		if len(fields) == 0 {
			p.numIgnoredLines++
			continue
		}

		switch fields[0] {
		case "v":
			coords := mapStringSliceToFloat(fields[1:])
			point := r.NewPoint(coords[0], coords[1], coords[2])
			p.vertices = append(p.vertices, point)
		case "f":
			indices, normalIndices := mapFieldsToIndicesAndNormals(fields[1:])

			vertices := getVerticesForIndexes(p.vertices, indices)
			normals := getNormalsForIndexes(p.normals, normalIndices)

			triangles := createTriangles(vertices, normals)

			for _, triangle := range triangles {
				if p.material != nil {
					triangle.SetMaterial(p.material)
				}

				if p.NewMaterial != nil {
					switch triangle.(type) {
					case (*r.Triangle):
						triangle.(*r.Triangle).NewMaterial = p.NewMaterial
					case (*r.SmoothTriangle):
						triangle.(*r.SmoothTriangle).NewMaterial = p.NewMaterial
					}
				}

				if curGroup != nil {
					curGroup.AddChild(triangle)
				} else {
					p.defaultGroup.AddChild(triangle)
				}
			}
		case "g":
			if curGroup != nil {
				p.defaultGroup.AddChild(curGroup)
			}

			curGroup = r.NewGroup()
		case "vn":
			coords2 := mapStringSliceToFloat(fields[1:])
			vec := r.NewVec(coords2[0], coords2[1], coords2[2])

			p.normals = append(p.normals, vec)
		default:
			p.numIgnoredLines++
		}
	}

	if curGroup != nil {
		p.defaultGroup.AddChild(curGroup)
	}

	return p.defaultGroup
}

func createTriangles(vertices []r.Point, normals []r.Vec) []r.Intersectable {
	triangles := make([]r.Intersectable, 0)

	for i := 1; i < len(vertices)-1; i++ {
		if len(normals) > 1 {
			triangle := r.NewSmoothTriangle(vertices[0], vertices[i], vertices[i+1], normals[0], normals[i], normals[i+1])
			triangles = append(triangles, triangle)
		} else {
			triangle := r.NewTriangle(vertices[0], vertices[i], vertices[i+1])
			triangles = append(triangles, triangle)
		}

	}

	return triangles
}

func getVerticesForIndexes(vertices []r.Point, indices []int) []r.Point {
	out := make([]r.Point, 0, len(indices))

	for _, index := range indices {
		out = append(out, vertices[index])
	}

	return out
}

func getNormalsForIndexes(normals []r.Vec, indices []int) []r.Vec {
	out := make([]r.Vec, 0, len(indices))

	for _, index := range indices {
		out = append(out, normals[index])
	}

	return out
}

func mapFieldsToIndicesAndNormals(slice []string) ([]int, []int) {
	indices, normalIndices := make([]int, 0), make([]int, 0)

	for _, item := range slice {
		numbers := strings.Split(item, "/")

		index, err := strconv.Atoi(numbers[0])

		if err != nil {
			panic(fmt.Sprintf("Could not parse file, invalid value %v", item))
		}

		indices = append(indices, index)

		if len(numbers) > 1 {
			normal, err := strconv.Atoi(numbers[2])

			if err != nil {
				panic(fmt.Sprintf("Could not parse file, invalid value %v", item))
			}

			normalIndices = append(normalIndices, normal)
		}
	}

	return indices, normalIndices
}

// func mapStringSliceToInt(slice []string) []int {
// 	out := make([]int, 0)

// 	for _, item := range slice {
// 		value, err := strconv.Atoi(item)

// 		if err != nil {
// 			panic(fmt.Sprintf("Could not parse file, invalid value %v", item))
// 		}

// 		out = append(out, value)
// 	}

// 	return out
// }

func mapStringSliceToFloat(slice []string) []float64 {
	out := make([]float64, 0)

	for _, item := range slice {
		value, err := strconv.ParseFloat(item, 64)

		if err != nil {
			panic(fmt.Sprintf("Could not parse file, invalid value %v", item))
		}

		out = append(out, value)
	}

	return out
}

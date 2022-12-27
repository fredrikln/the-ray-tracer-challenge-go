package raytracer

import (
	"math"
	"sort"

	c "github.com/fredrikln/the-ray-tracer-challenge-go/common"
	common "github.com/fredrikln/the-ray-tracer-challenge-go/common"
)

type Cone struct {
	Transform *Matrix
	Material  *Material
	Minimum   float64
	Maximum   float64
	Closed    bool
}

func NewCone() *Cone {
	return &Cone{
		NewIdentityMatrix(),
		NewMaterial(),
		math.Inf(-1),
		math.Inf(1),
		false,
	}
}

func NewGlassCone() *Cone {
	return &Cone{
		NewIdentityMatrix(),
		NewMaterial().SetTransparency(1.0).SetRefractiveIndex(1.5),
		math.Inf(-1),
		math.Inf(1),
		false,
	}
}

func (co *Cone) GetMaterial() *Material {
	return co.Material
}

func (co *Cone) SetMaterial(m *Material) Intersectable {
	co.Material = m

	return co
}

func (co *Cone) GetTransform() *Matrix {
	return co.Transform
}

func (co *Cone) SetTransform(m *Matrix) Intersectable {
	co.Transform = m

	return co
}

func (co *Cone) Intersect(worldRay Ray) []Intersection {
	localRay := worldRay.Mul(co.Transform.Inverse())

	a := math.Pow(localRay.Direction.X, 2) - math.Pow(localRay.Direction.Y, 2) + math.Pow(localRay.Direction.Z, 2)
	b := 2*localRay.Origin.X*localRay.Direction.X - 2*localRay.Origin.Y*localRay.Direction.Y + 2*localRay.Origin.Z*localRay.Direction.Z
	c := math.Pow(localRay.Origin.X, 2) - math.Pow(localRay.Origin.Y, 2) + math.Pow(localRay.Origin.Z, 2)

	disc := math.Pow(b, 2) - 4*a*c

	if disc < 0 {
		return []Intersection{}
	}

	xs := make([]Intersection, 0)

	if common.WithinTolerance(a, 0, 1e-5) && !common.WithinTolerance(b, 0, 1e-5) {
		t := -c / (2 * b)

		xs = append(xs, NewIntersection(t, co))
	}

	t0 := (-b - math.Sqrt(disc)) / (2 * a)
	t1 := (-b + math.Sqrt(disc)) / (2 * a)

	if t0 > t1 {
		t0, t1 = t1, t0
	}

	y0 := localRay.Origin.Y + t0*localRay.Direction.Y

	if co.Minimum < y0 && y0 < co.Maximum {
		xs = append(xs, NewIntersection(t0, co))
	}

	y1 := localRay.Origin.Y + t1*localRay.Direction.Y
	if co.Minimum < y1 && y1 < co.Maximum {
		xs = append(xs, NewIntersection(t1, co))
	}

	xs = append(xs, intersectCaps2(co, localRay)...)

	sort.Slice(xs, func(i, j int) bool {
		return xs[i].Time < xs[j].Time
	})

	return xs
}

func (co *Cone) NormalAt(worldPoint Point) Vec {
	objectPoint := worldPoint.MulMat(co.Transform.Inverse())

	var objectNormal Vec

	dist := math.Pow(objectPoint.X, 2) + math.Pow(objectPoint.Z, 2)

	if dist < 1 && objectPoint.Y >= co.Maximum-1e-5 {
		objectNormal = NewVec(0, 1, 0)
	} else if dist < 1 && objectPoint.Y <= co.Minimum+1e-5 {
		objectNormal = NewVec(0, -1, 0)
	} else {
		y := math.Sqrt(math.Pow(objectPoint.X, 2) + math.Pow(objectPoint.Z, 2))

		if objectPoint.Y > 0 {
			y = -y
		}

		objectNormal = NewVec(objectPoint.X, y, objectPoint.Z)
	}

	worldNormal := objectNormal.MulMat(co.Transform.Inverse().Transpose())

	return worldNormal
}

func checkCap2(r Ray, t float64, radius float64) bool {
	x := r.Origin.X + t*r.Direction.X
	z := r.Origin.Z + t*r.Direction.Z

	return (math.Pow(x, 2) + math.Pow(z, 2)) <= math.Abs(radius)
}

func intersectCaps2(co *Cone, r Ray) []Intersection {
	xs := make([]Intersection, 0)

	if !co.Closed || c.WithinTolerance(r.Direction.Y, 0, 1e-5) {
		return []Intersection{}
	}

	t0 := (co.Minimum - r.Origin.Y) / r.Direction.Y
	if checkCap2(r, t0, co.Minimum) {
		xs = append(xs, NewIntersection(t0, co))
	}

	t1 := (co.Maximum - r.Origin.Y) / r.Direction.Y
	if checkCap2(r, t1, co.Maximum) {
		xs = append(xs, NewIntersection(t1, co))
	}

	return xs
}

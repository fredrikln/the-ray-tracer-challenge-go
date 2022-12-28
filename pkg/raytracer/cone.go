package raytracer

import (
	"math"
)

type Cone struct {
	Transform *Matrix
	Material  *Material
	Minimum   float64
	Maximum   float64
	Closed    bool
	Parent    *Group
}

func NewCone() *Cone {
	return &Cone{
		Transform: NewIdentityMatrix(),
		Material:  NewMaterial(),
		Minimum:   math.Inf(-1),
		Maximum:   math.Inf(1),
		Closed:    false,
	}
}

func NewGlassCone() *Cone {
	return &Cone{
		Transform: NewIdentityMatrix(),
		Material:  NewMaterial().SetTransparency(1.0).SetRefractiveIndex(1.5),
		Minimum:   math.Inf(-1),
		Maximum:   math.Inf(1),
		Closed:    false,
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

func (co *Cone) GetParent() *Group {
	return co.Parent
}
func (co *Cone) SetParent(g *Group) Intersectable {
	co.Parent = g

	return co
}

func (co *Cone) Intersect(worldRay Ray) []Intersection {
	localRay := worldRay.Mul(co.Transform.Inverse())

	a := localRay.Direction.X*localRay.Direction.X - localRay.Direction.Y*localRay.Direction.Y + localRay.Direction.Z*localRay.Direction.Z
	b := 2*localRay.Origin.X*localRay.Direction.X - 2*localRay.Origin.Y*localRay.Direction.Y + 2*localRay.Origin.Z*localRay.Direction.Z
	c := localRay.Origin.X*localRay.Origin.X - localRay.Origin.Y*localRay.Origin.Y + localRay.Origin.Z*localRay.Origin.Z

	disc := b*b - 4*a*c

	if disc < 0 {
		return []Intersection{}
	}

	var xs []Intersection

	if WithinTolerance(a, 0, 1e-5) && !WithinTolerance(b, 0, 1e-5) {
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

	return xs
}

func (co *Cone) LocalNormalAt(objectPoint Point) Vec {
	var objectNormal Vec

	dist := objectPoint.X*objectPoint.X + objectPoint.Z*objectPoint.Z

	if dist < 1 && objectPoint.Y >= co.Maximum-1e-5 {
		objectNormal = NewVec(0, 1, 0)
	} else if dist < 1 && objectPoint.Y <= co.Minimum+1e-5 {
		objectNormal = NewVec(0, -1, 0)
	} else {
		y := math.Sqrt(objectPoint.X*objectPoint.X + objectPoint.Z*objectPoint.Z)

		if objectPoint.Y > 0 {
			y = -y
		}

		objectNormal = NewVec(objectPoint.X, y, objectPoint.Z)
	}

	return objectNormal
}

func (co *Cone) NormalAt(worldPoint Point, i Intersection) Vec {
	objectPoint := co.WorldToObject(worldPoint)

	objectNormal := co.LocalNormalAt(objectPoint)

	worldNormal := co.NormalToWorld(objectNormal)

	return worldNormal
}

func checkCap2(r Ray, t float64, radius float64) bool {
	x := r.Origin.X + t*r.Direction.X
	z := r.Origin.Z + t*r.Direction.Z

	return (x*x + z*z) <= math.Abs(radius)
}

func intersectCaps2(co *Cone, r Ray) []Intersection {
	var xs []Intersection

	if !co.Closed || WithinTolerance(r.Direction.Y, 0, 1e-5) {
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

func (co *Cone) WorldToObject(p Point) Point {
	parent := co.GetParent()

	if parent != nil {
		p = parent.WorldToObject(p)
	}

	return co.GetTransform().Inverse().MulPoint(p)
}

func (co *Cone) NormalToWorld(n Vec) Vec {
	normal := co.GetTransform().Inverse().Transpose().MulVec(n).Norm()

	parent := co.GetParent()

	if parent != nil {
		normal = parent.NormalToWorld(normal)
	}

	return normal
}

package raytracer

import "math"

type Cube struct {
	Transform *Matrix
	Material  *Material
	Parent    *Group
}

func NewCube() *Cube {
	return &Cube{
		Transform: NewIdentityMatrix(),
		Material:  NewMaterial(),
	}
}

func NewGlassCube() *Cube {
	return &Cube{
		Transform: NewIdentityMatrix(),
		Material:  NewMaterial().SetTransparency(1.0).SetRefractiveIndex(1.5),
	}
}

func (c *Cube) GetMaterial() *Material {
	return c.Material
}
func (c *Cube) SetMaterial(m *Material) Intersectable {
	c.Material = m

	return c
}

func (c *Cube) GetTransform() *Matrix {
	return c.Transform
}
func (c *Cube) SetTransform(m *Matrix) Intersectable {
	c.Transform = m

	return c
}

func (c *Cube) GetParent() *Group {
	return c.Parent
}
func (c *Cube) SetParent(g *Group) Intersectable {
	c.Parent = g

	return c
}

func (c *Cube) Intersect(worldRay Ray) []Intersection {
	localRay := worldRay.Mul(c.Transform.Inverse())

	xtmin, xtmax := checkAxis(localRay.Origin.X, localRay.Direction.X)
	ytmin, ytmax := checkAxis(localRay.Origin.Y, localRay.Direction.Y)
	ztmin, ztmax := checkAxis(localRay.Origin.Z, localRay.Direction.Z)

	tmin := math.Max(math.Max(xtmin, ytmin), ztmin)
	tmax := math.Min(math.Min(xtmax, ytmax), ztmax)

	if tmin > tmax {
		return []Intersection{}
	}

	return []Intersection{
		NewIntersection(tmin, c),
		NewIntersection(tmax, c),
	}
}

func (c *Cube) NormalAt(worldPoint Point) Vec {
	objectPoint := c.WorldToObject(worldPoint)

	maxC := math.Max(math.Max(math.Abs(objectPoint.X), math.Abs(objectPoint.Y)), math.Abs(objectPoint.Z))

	var objectNormal Vec

	if maxC == math.Abs(objectPoint.X) {
		objectNormal = NewVec(objectPoint.X, 0, 0)
	} else if maxC == math.Abs(objectPoint.Y) {
		objectNormal = NewVec(0, objectPoint.Y, 0)
	} else {
		objectNormal = NewVec(0, 0, objectPoint.Z)
	}

	worldNormal := c.NormalToWorld(objectNormal)

	return worldNormal.Norm()
}

func checkAxis(origin, direction float64) (float64, float64) {
	tmin_numerator := -1 - origin
	tmax_numerator := 1 - origin

	var tmin, tmax float64

	if math.Abs(direction) > 1e-5 {
		tmin = tmin_numerator / direction
		tmax = tmax_numerator / direction
	} else {
		tmin = tmin_numerator * math.Inf(1)
		tmax = tmax_numerator * math.Inf(1)
	}

	if tmin > tmax {
		tmin, tmax = tmax, tmin
	}

	return tmin, tmax
}

func (cu *Cube) WorldToObject(p Point) Point {
	parent := cu.GetParent()

	if parent != nil {
		p = parent.WorldToObject(p)
	}

	return cu.GetTransform().Inverse().MulPoint(p)
}

func (cu *Cube) NormalToWorld(n Vec) Vec {
	inv := cu.GetTransform().Inverse()
	trans := inv.Transpose()
	normal := trans.MulVec(n).Norm()

	parent := cu.GetParent()

	if parent != nil {
		normal = parent.NormalToWorld(normal)
	}

	return normal
}

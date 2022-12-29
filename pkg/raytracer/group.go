package raytracer

import (
	"sort"
)

type Group struct {
	Transform *Matrix
	Items     []Intersectable
	Parent    Intersectable
}

func NewGroup() *Group {
	return &Group{
		Transform: NewIdentityMatrix(),
		Items:     make([]Intersectable, 0),
	}
}

func (g *Group) AddChild(c Intersectable) *Group {
	c.SetParent(g)
	g.Items = append(g.Items, c)

	return g
}

func (g *Group) SetMaterial(*Material) Intersectable {
	panic("Should never happen")
}
func (g *Group) GetMaterial() *Material {
	panic("Should never happen")
}

func (g *Group) SetTransform(m *Matrix) Intersectable {
	g.Transform = m

	return g
}
func (g *Group) GetTransform() *Matrix {

	return g.Transform
}

func (g *Group) GetParent() Intersectable {
	return g.Parent
}
func (g *Group) SetParent(p Intersectable) Intersectable {
	g.Parent = p

	return g
}

func (g *Group) Intersect(worldRay Ray) []Intersection {
	if len(g.Items) == 0 {
		return []Intersection{}
	}

	localRay := worldRay.Mul(g.Transform.Inverse())

	xs := make([]Intersection, 0)

	for _, childObject := range g.Items {
		xs = append(xs, childObject.Intersect(localRay)...)
	}

	sort.Sort(IntersectonSorter(xs))

	return xs
}
func (g *Group) NormalAt(Point, Intersection) Vec {
	panic("Should never happen")
}

func (g *Group) WorldToObject(p Point) Point {
	parent := g.GetParent()

	if parent != nil {
		p = parent.WorldToObject(p)
	}

	return g.GetTransform().Inverse().MulPoint(p)
}

func (g *Group) NormalToWorld(n Vec) Vec {
	inv := g.GetTransform().Inverse()
	trans := inv.Transpose()
	normal := trans.MulVec(n).Norm()

	parent := g.GetParent()

	if parent != nil {
		normal = parent.NormalToWorld(normal)
	}

	return normal
}

func (g *Group) Includes(object Intersectable) bool {
	for _, child := range g.Items {
		switch child.(type) {
		case *Group:
			return (child.(*Group)).Includes(object)
		case *CSG:
			return (child.(*CSG)).Includes(object)
		default:
			return object == child
		}
	}

	return false
}

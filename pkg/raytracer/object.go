package raytracer

type object struct {
	transform    *Matrix
	parent       Intersectable
	material     Scatters
	parentObject Intersectable
}

func newObject() object {
	return object{
		transform:    NewIdentityMatrix(),
		parent:       nil,
		material:     NewDiffuse(NewColor(1, 1, 1)),
		parentObject: nil,
	}
}

func (o *object) Intersect(worldRay Ray) []Intersection {
	localRay := worldRay.Mul(o.transform.Inverse())

	xs := o.parentObject.LocalIntersect(localRay)

	return xs
}
func (o *object) LocalIntersect(Ray) []Intersection {
	panic("Should not happen")
}

func (o *object) NormalAt(worldPoint Point, i Intersection) Vec {
	objectPoint := o.WorldToObject(worldPoint)
	objectNormal := o.parentObject.LocalNormalAt(objectPoint, i)
	worldNormal := o.NormalToWorld(objectNormal)

	return worldNormal.Norm()
}
func (o *object) LocalNormalAt(Point, Intersection) Vec {
	panic("Should not happen")
}

func (o *object) SetTransform(t *Matrix) Intersectable {
	o.transform = t

	return o
}
func (o *object) GetTransform() *Matrix {
	return o.transform
}

func (o *object) GetParent() Intersectable {
	return o.parent
}
func (o *object) SetParent(i Intersectable) Intersectable {
	o.parent = i

	return o
}

func (o *object) WorldToObject(p Point) Point {
	parent := o.GetParent()

	if parent != nil {
		p = parent.WorldToObject(p)
	}

	return o.GetTransform().Inverse().MulPoint(p)
}
func (o *object) NormalToWorld(n Vec) Vec {
	normal := o.GetTransform().Inverse().Transpose().MulVec(n).Norm()

	parent := o.GetParent()

	if parent != nil {
		normal = parent.NormalToWorld(normal)
	}

	return normal
}

func (o *object) Bounds() *BoundingBox {
	// Does nothing
	return NewBoundingBox()
}
func (o *object) Divide(int) {
	// Does nothing
}

func (o *object) GetNewMaterial() Scatters {
	return o.material
}
func (o *object) SetNewMaterial(s Scatters) Intersectable {
	o.material = s

	return o
}

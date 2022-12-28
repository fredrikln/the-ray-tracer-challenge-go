package raytracer

type Group struct {
	Transform *Matrix
	Items     []*Intersectable
}

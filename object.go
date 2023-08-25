package fisiks

type Object struct {
	OriginalX float64
	OriginalY float64
	Mass      float64
	Forces    []*Force

	Width  float64
	Height float64

	resource string
	x        Ordinate
	y        Ordinate
}

func NewObject(width float64, height float64, mass float64, resource string) *Object {
	obj := &Object{
		Mass:     mass,
		Width:    width,
		Height:   height,
		resource: resource,
	}

	return obj
}

func (obj *Object) ComputeForces() (x float64, y float64) {
	for _, f := range obj.Forces {
		fx, fy := (*f)(obj)
		x += fx
		y += fy
	}
	return
}

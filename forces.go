package fisiks

import "errors"

// Force takes an object and its properties as parameters and return the forces (in Newton) on the x and y axes
type Force func(obj *Object) (float64, float64)

type ForceConfig func(args ...interface{}) (Force, error)

// Weight takes no arguments
func Weight(_ ...interface{}) (Force, error) {
	return func(obj *Object) (_ float64, y float64) {
		y = obj.Mass * g
		return
	}, nil
}

// FluidResistance takes as arguments :
// - the volumic mass of the fluid
// - the drag coefficient of the object
func FluidResistance(args ...interface{}) (Force, error) {
	if len(args) < 2 {
		return nil, errors.New("not enough arguments")
	}
	pho, ok := args[0].(float64)
	if !ok {
		return nil, WrongArgumentType{
			Argument: "pho",
			Expected: "float64",
		}
	}
	CD, ok := args[1].(float64)
	if !ok {
		return nil, WrongArgumentType{
			Argument: "CD",
			Expected: "float64",
		}
	}

	return func(obj *Object) (x float64, y float64) {
		x = pho * CD * obj.Height * obj.Height * obj.x.velocity * obj.x.velocity / 2
		y = pho * CD * obj.Width * obj.Width * obj.y.velocity * obj.y.velocity / 2
		return
	}, nil
}

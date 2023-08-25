package fisiks

import (
	"errors"
	"math"
)

// Force takes an object and its properties as parameters and return the forces (in Newton) on the x and y axes
type Force func(obj *Object) (fx float64, fy float64)

type ForceConfig func(args ...interface{}) (Force, error)

// Weight takes one argument : the gravity value
func Weight(args ...interface{}) (Force, error) {
	g := args[0].(float64)
	return func(obj *Object) (_ float64, fy float64) {
		fy = obj.Mass * g
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

	return func(obj *Object) (fx float64, fy float64) {
		fx = pho * CD * obj.Height * obj.Height * obj.x.velocity * obj.x.velocity / 2
		fy = pho * CD * obj.Width * obj.Width * obj.y.velocity * obj.y.velocity / 2
		return
	}, nil
}

// ElectricCharge takes three arguments :
// - x ordinate of the charge
// - y ordinate of the charge
// - charge value (in Coulomb)
// - scale (for distance calculation)
// TODO: verify arguments
func ElectricCharge(args ...interface{}) (Force, error) {
	x := args[0].(float64)
	y := args[1].(float64)
	q := args[2].(float64)
	scale := args[3].(float64)
	return func(obj *Object) (fx float64, fy float64) {
		d := DistanceToObj(obj, x, y) * scale
		if d == 0 {
			return
		}

		radiusNorm := K * obj.Charge * q / d
		if x == obj.x.pos {
			fy = radiusNorm * Sign(obj.y.pos-y)
			return
		} else {
			ratio := (obj.y.pos - y) / (obj.x.pos - x)
			// atan(x) is the angle between the unit vector and the axes
			// cos(atan(x)) = 1/sqrt(1+x²)
			fx = radiusNorm / math.Sqrt(1+ratio*ratio)
			// sin(atan(x)) = x/sqrt(1+x²)
			fy = radiusNorm * ratio / math.Sqrt(1+ratio*ratio)
			// Adding a +π after the atan so inverting the signs of cos and sin (angle is now at the opposite on the circle)
			if obj.x.pos < x {
				fx *= -1
				fy *= -1
			}
		}
		return
	}, nil
}

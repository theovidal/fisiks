package fisiks

import "math"

func Distance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2))
}

func DistanceToObj(obj *Object, x float64, y float64) float64 {
	return Distance(obj.x.pos, obj.y.pos, x, y)
}

func Sign(x float64) float64 {
	if x < 0 {
		return -1
	} else {
		return 1
	}
}

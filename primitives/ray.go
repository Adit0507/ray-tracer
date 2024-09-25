package primitives

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) Point(t float64) Vector {
	b := r.Direction.MultiplyScaler(t)
	
	return r.Origin.Add(b)
}

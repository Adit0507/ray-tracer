package primitives

type Metal struct {
	C Vector
	Fuzz float64
}

func (m Metal) Bounce(input Ray, hit Hit) (bool, Ray) {
	direction := reflect(input.Direction, hit.Normal)
	bouncedRay := Ray{hit.P, direction.Add(VectorInUnitSphere().MultiplyScaler(m.Fuzz))}
	bounced := direction.Dot(hit.Normal) > 0
	return bounced, bouncedRay
}

func (m Metal) Color() Vector {
	return m.C
}

func reflect(v Vector, n Vector) Vector {
	b := 2 * v.Dot(n)
	return v.Subtract(n.MultiplyScaler(b))
}

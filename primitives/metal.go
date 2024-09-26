package primitives

type Metal struct {
	C Color
	Fuzz float64
}

func (m Metal) Bounce(input Ray, hit Hit) (bool, Ray) {
	direction := input.Direction.Reflect(hit.Normal)
	bouncedRay := Ray{hit.P, direction.Add(VectorInUnitSphere().MultiplyScaler(m.Fuzz))}
	bounced := direction.Dot(hit.Normal) > 0
	return bounced, bouncedRay
}

func (m Metal) Color() Color {
	return m.C
}

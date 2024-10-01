package primitives

import "math/rand"

type Metal struct {
	Attenuation Color
	Fuzz float64
}

func (m Metal) Bounce(input Ray, hit Hit, rnd *rand.Rand) (bool, Ray) {
	direction := input.Direction.Reflect(hit.Normal)
	bouncedRay := Ray{hit.P, direction.Add(VectorInUnitSphere(rnd).MultiplyScaler(m.Fuzz))}
	bounced := direction.Dot(hit.Normal) > 0
	return bounced, bouncedRay
}

func (m Metal) Color() Color {
	return m.Attenuation
}

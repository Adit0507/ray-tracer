package primitives

import "math/rand"

type Lambertian struct {
	Attenuation Color
}

func (l Lambertian) Bounce(input Ray, hit Hit, rnd *rand.Rand) (bool, Ray) {
	direction := hit.Normal.Add(VectorInUnitSphere(rnd))
	return true, Ray{hit.P, direction}
}

func (l Lambertian) Color() Color {
	return l.Attenuation
}

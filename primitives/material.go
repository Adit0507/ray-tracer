package primitives

import "math/rand"

type Material interface {
	Bounce(input Ray, hit Hit, rnd *rand.Rand) (bool, Ray)
	Color() Color
}
package render

import (
	"math/rand"
	"ray/primitives"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomScene() *primitives.World {
	world := &primitives.World{}

	floor := primitives.NewSphere(0, -1000, 0, 1000,
		primitives.Lambertian{Attenuation: primitives.Color{
			R: 0.5,
			G: 0.5,
			B: 0.5}})

	world.Add(floor)

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			material := rand.Float64()

			center := primitives.Vector{
				X: float64(a) + 0.9*rand.Float64(),
				Y: 0.2,
				Z: float64(b) + 0.9*rand.Float64()}

			if center.Subtract(primitives.Vector{X: 4, Y: 0.2, Z: 0}).Length() > 0.9 {
				if material < 0.8 {
					lambertian := primitives.NewSphere(center.X, center.Y, center.Z, 0.2,
						primitives.Lambertian{Attenuation: primitives.Color{
							R: rand.Float64() * rand.Float64(),
							G: rand.Float64() * rand.Float64(),
							B: rand.Float64() * rand.Float64()}})

					world.Add(lambertian)
				} else if material < 0.95 {
					metal := primitives.NewSphere(center.X, center.Y, center.Z, 0.2,
						primitives.Metal{Attenuation: primitives.Color{
							R: 0.5 * (1.0 + rand.Float64()),
							G: 0.5 * (1.0 + rand.Float64()),
							B: 0.5 * (1.0 + rand.Float64())},
							Fuzz: 0.5 + rand.Float64()})

					world.Add(metal)
				} else {
					glass := primitives.NewSphere(center.X, center.Y, center.Z, 0.2,
						primitives.Dielectric{Index: 1.5})

					world.Add(glass)
				}
			}
		}
	}

	glass := primitives.NewSphere(0, 1, 0, 1.0, primitives.Dielectric{Index: 1.5})
	lambertian := primitives.NewSphere(-4, 1, 0, 1.0, primitives.Lambertian{Attenuation: primitives.Color{R: 0.4, G: 0.0, B: 0.1}})
	metal := primitives.NewSphere(4, 1, 0, 1.0, primitives.Metal{Attenuation: primitives.Color{R: 0.7, G: 0.6, B: 0.5}, Fuzz: 0.0})

	world.AddAll(glass, lambertian, metal)
	return world
}
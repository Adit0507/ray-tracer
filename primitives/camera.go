package primitives

import (
	"math"
	"math/rand"
)

type Camera struct {
	lowerLeft, horizontal, vertical, origin, u, v, w Vector
	lensRadius                                       float64
}

var vUp = Vector{X: 0, Y: 1, Z: 0}

func NewCamera(lookFrom, lookAt Vector, vFov, aspect, aperture float64) *Camera {
	c := Camera{}

	c.origin = lookFrom
	c.lensRadius = aperture / 2

	theta := vFov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	w := lookFrom.Subtract(lookAt).Normalize()
	u := vUp.Cross(w).Normalize()
	v := w.Cross(u)

	focusDist := lookFrom.Subtract(lookAt).Length()
	x := u.MultiplyScaler(halfWidth * focusDist)
	y := v.MultiplyScaler(halfHeight * focusDist)

	c.lowerLeft = c.origin.Subtract(x).Subtract(y).Subtract(w.MultiplyScaler(focusDist))
	c.horizontal = x.MultiplyScaler(2)
	c.vertical = y.MultiplyScaler(2)

	c.w = w
	c.u = u
	c.v = v

	return &c
}

func (c *Camera) RayAt(u, v float64, rnd *rand.Rand) Ray {
	var randomInUnitDisc Vector
	for {
		randomInUnitDisc = Vector{rnd.Float64(), rnd.Float64(), 0}.MultiplyScaler(2).Subtract(Vector{1, 1, 0})
		if randomInUnitDisc.Dot(randomInUnitDisc) < 1 {
			break
		}
	}

	rd := randomInUnitDisc.MultiplyScaler(c.lensRadius)
	offset := c.u.MultiplyScaler(rd.X).Add(c.v.MultiplyScaler(rd.Y))

	horizontal := c.horizontal.MultiplyScaler(u)
	vertical := c.vertical.MultiplyScaler(v)

	origin := c.origin.Add(offset)
	direction := c.lowerLeft.Add(horizontal).Add(vertical).Subtract(c.origin).Subtract(offset)
	return Ray{origin, direction}
}

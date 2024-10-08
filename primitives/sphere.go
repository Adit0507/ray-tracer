package primitives

import "math"

type Sphere struct {
	Center Vector
	Radius float64
	Material
}

func NewSphere(x, y,z, radius float64, m Material) *Sphere{
	return &Sphere{Vector{x,y,z}, radius, m}
}

func (s *Sphere) Hit(r Ray, tMin float64, tMax float64) (bool, Hit) {
	oc := r.Origin.Subtract(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
	c := oc.Dot(oc) - s.Radius*s.Radius
	discriminant := b*b - a*c

	hit := Hit{Material: s.Material}

	if discriminant > 0.0 {
		t := (-b - math.Sqrt(discriminant)) / a
		if t < tMax && t > tMin {
			hit.T = t
			hit.P = r.Point(t)
			hit.Normal = hit.P.Subtract(s.Center).DivideScaler(s.Radius)
			return true, hit
		}

		t = (-b + math.Sqrt(discriminant)) / a
		if t < tMax && t > tMin {
			hit.T = t
			hit.P = r.Point(t)
			hit.Normal = hit.P.Subtract(s.Center).DivideScaler(s.Radius) 
			return true, hit
		}
	}

	return false, Hit{}
}

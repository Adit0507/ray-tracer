package geometry

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) Point(t float64) Vector {
	b := r.Direction.MultiplyScaler(t)
	a := r.Origin
	
	return a.Add(b)
}

func (r Ray) HitSphere(s Sphere) bool {
	oc := r.Origin.Subtract(s.Center)
	a := r.Direction.Dot(r.Direction)
	b := 2.0 * oc.Dot(r.Direction)
    c := oc.Dot(oc) - s.Radius*s.Radius
    discriminant := b*b - 4*a*c
	
	return discriminant > 0
}


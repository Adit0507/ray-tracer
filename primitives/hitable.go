package primitives

type Hit struct {
	T         float64
	P, Normal Vector
	Material
}

type Hitable interface {
	Hit(r *Ray, tMin float64, tMax float64) (bool, Hit)
}
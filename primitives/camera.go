package primitives

type Camera struct {
	lowerLeft, horizontal, vertical, origin Vector
}

func NewCamera() Camera {
	c := Camera{}

	c.lowerLeft = Vector{-2, -1, -1}
	c.horizontal = Vector{4, 0, 0}
	c.vertical = Vector{0, 4, 0}
	c.origin = Vector{0, 0, 0}

	return c
}

func (c Camera) position(u float64, v float64) Vector {
	horizontal := c.horizontal.MultiplyScaler(u)
	vertical := c.vertical.MultiplyScaler(v)

	return horizontal.Add(vertical)
}

func (c Camera) direction(position Vector) Vector {
	return c.lowerLeft.Add(position)
}

func (c Camera) RayAt(u float64, v float64) Ray {
	position := c.position(u, v)
	direction := c.direction(position)

	return Ray{c.origin, direction}
}

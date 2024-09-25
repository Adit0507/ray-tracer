package primitives

type Camera struct {
	lowerLeft, horizontal, vertical, origin Vector	
}

func NewCamera() *Camera {
	c :=new(Camera)

	c.lowerLeft  =Vector{-2, -1, -1}
	c.horizontal = Vector{4,  0,0}
	c.vertical = Vector{0, 4, 0}
	c.origin = Vector{0, 0, 0}

	return c
}


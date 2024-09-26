package primitives

type Color struct {
	R, G, B float64
}

var (
	Black = Color{}
	White = Color{1, 1, 1}
	Blue  = Color{0.5, 0.7, 1}
)

func (c Color) Add(o Color) Color {
	return Color{c.R + o.R, c.G + o.G, c.B + o.B}
}

func (c Color) Multiply(o Color) Color {
	return Color{c.R * o.R, c.G * o.G, c.B * o.B}
}

func (c Color) AddScalar(f float64) Color {
	return Color{c.R * f, c.G * f, c.B * f}
}

func (c Color) MultiplyScalar(f float64) Color {
	return Color{c.R * f, c.G * f, c.B * f}
}

func (c Color) DivideScalar(f float64) Color {
	return Color{c.R / f, c.G / f, c.B / f}
}

func Gradient(a, b Color, f float64) Color {
	f = 0.5 * (f + 1.0)
	return a.MultiplyScalar(1.0 - f).Add(b.MultiplyScalar(f))
}
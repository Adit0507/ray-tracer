package geometry

import (
	"math"
)

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) Dot(o Vector) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

func (v Vector) Normalize() Vector {
	l := v.Length()
	return Vector{v.X / l, v.Y / l, v.Z / l}
}

func (v Vector) Add(o Vector) Vector {
	return Vector{v.X + o.X, v.Y + o.Y, v.Z + o.Z}
}

func (v Vector) Subtract(o Vector) Vector {
	return Vector{v.X - o.X, v.Y - o.Y, v.Z - o.Z}
}

func (v Vector) AddScaler(t float64) Vector {
	return Vector{v.X + t, v.Y + t, v.Z + t}
}

func (v Vector) SubtractScaler(t float64) Vector {
	return Vector{v.X - t, v.Y - t, v.Z - t}
}

func (v Vector) MultiplyScaler(t float64) Vector {
	return Vector{v.X *t, v.Y * t, v.Z * t}
}

func (v Vector) DivideScaler(t float64) Vector {
	return Vector{v.X / t, v.Y / t, v.Z / t}
}

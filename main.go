package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	p "ray/primitives"
)

const (
	sx = 500
	sy = 300
	ns = 100
	c  = 255.99
)

var (
	white = p.Vector{1.0, 1.0, 1.0}
	blue  = p.Vector{0.5, 0.7, 1.0}

	camera = p.NewCamera()

	sphere = p.Sphere{p.Vector{0, 0, -1}, 0.5}
	floor  = p.Sphere{p.Vector{0, -100.5, -1}, 100}

	world = p.World{[]p.Hitable{&sphere, &floor}}
)

func color(r *p.Ray, h p.Hitable) p.Vector {
	hit, record := h.Hit(r, 0.0, math.MaxFloat64)

	if hit {
		return record.Normal.AddScaler(1.0).MultiplyScaler(0.5)
	}

	unitDirection := r.Direction.Normalize()

	return gradient(&unitDirection)
}

func gradient(v *p.Vector) p.Vector {
	t := 0.5 * (v.Y + 1)

	return white.MultiplyScaler(1 - t).Add(blue.MultiplyScaler(t))
}

func main() {

	f, err := os.Create("out.ppm")
	defer f.Close()
	check(err, "Error opening file: %v\n")

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", sx, sy)
	check(err, "Error writing to file: %v\n")

	for j := sy - 1; j >= 0; j-- {
		for i := 0; i < sx; i++ {
			rgb := p.Vector{}

			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(sx)
				v := (float64(j) + rand.Float64()) / float64(sy)

				r := camera.RayAt(u, v)
				color := color(&r, &world)
				rgb = rgb.Add(color)
			}

			rgb = rgb.DivideScaler(float64(ns))

			ir := int(c * rgb.X)
			ig := int(c * rgb.Y)
			ib := int(c * rgb.Z)

			_, err = fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
			check(err, "Error writing to file: %v \n")
		}
	}
}

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

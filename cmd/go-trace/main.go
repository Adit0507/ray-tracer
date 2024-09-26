package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	p "ray/primitives"
	"time"
)

const (
	sx = 500
	sy = 300
	ns = 100
	c  = 255.99
)

func color(r p.Ray, world p.Hitable, depth int) p.Color {
	hit, record := world.Hit(r, 0.001, math.MaxFloat64)

	if hit {
		if depth < 50 {
			bounced, bouncedRay := record.Bounce(r, record)
			if bounced {
				newColor := color(bouncedRay, world, depth+1)
				return record.Material.Color().Multiply(newColor)
			}
		}
		return p.Black
	}

	return p.Gradient(p.White, p.Blue, r.Direction.Normalize().Y)
}

func createFile() *os.File {
	f, err := os.Create("out.ppm")
	check(err, "Error opening file: %v\n")
	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", sx, sy)
	check(err, "Error writting to file: %v\n")
	return f
}

func writeFile(f *os.File, rgb p.Color) {
	ir := int(c * math.Sqrt(rgb.R))
	ig := int(c * math.Sqrt(rgb.G))
	ib := int(c * math.Sqrt(rgb.B))

	_, err := fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
	check(err, "Error writing to file: %v\n")
}

func sample(world *p.World, camera *p.Camera, i, j int) p.Color {
	rgb := p.Color{}

	for s := 0; s < ns; s++ {
		u := (float64(i) + rand.Float64()) / float64(sx)
		v := (float64(j) + rand.Float64()) / float64(sy)

		ray := camera.RayAt(u, v)
		rgb = rgb.Add(color(ray, world, 0))
	}

	return rgb.DivideScalar(float64(ns))
}

func render(world *p.World, camera *p.Camera) {
	ticker := time.NewTicker(time.Millisecond * 100)

	go func() {
		for {
			<-ticker.C
			fmt.Print(".")
		}
	}()

	f := createFile()
	defer f.Close()

	start := time.Now()

	for j := sy - 1; j >= 0; j-- {
		for i := 0; i < sx; i++ {
			rgb := sample(world, camera, i, j)
			writeFile(f, rgb)
		}
	}
	ticker.Stop()
	fmt.Printf("\nDone.\nElapsed: %v\n", time.Since(start))
}

func main() {
	camera := p.NewCamera()
	world := p.World{}

	sphere := p.NewSphere(0, 0, -1, 0.5, p.Lambertian{p.Color{0.8, 0.3, 0.3}})
	floor := p.NewSphere(0, -100.5, -1, 100, p.Lambertian{p.Color{0.8, 0.8, 0.0}})
	metal := p.NewSphere(1, 0, -1, 0.5, p.Metal{p.Color{0.8, 0.6, 0.2}, 0.3})
	glass := p.NewSphere(-1, 0, -1, 0.5, p.Dielectric{})
	bubble := p.NewSphere(-1, 0, -1, -0.45, p.Dielectric{1.5})

	world.AddAll(&sphere, &floor, &metal, &glass, &bubble)
	render(&world, &camera)
}

func check(e error, s string) {
	if e != nil {
		fmt.Fprintf(os.Stderr, s, e)
		os.Exit(1)
	}
}

package render

import (
	"image"
	"math"
	"math/rand"
	p "ray/primitives"
	"sync"
	"time"
)

const (
	maxDepth = 50
	tMin     = 0.001
)

func color(r p.Ray, h p.Hitable, rnd *rand.Rand, depth int) p.Color {
	hit, record := h.Hit(r, tMin, math.MaxFloat64)

	if hit {
		if depth < maxDepth {
			bounced, bouncedRay := record.Bounce(r, record, rnd)
			if bounced {
				newColor := color(bouncedRay, h, rnd, depth+1)
				return record.Material.Color().Multiply(newColor)
			}
		}
		return p.Black
	}

	return p.Gradient(p.White, p.Blue, r.Direction.Normalize().Y)
}

func sample(h p.Hitable, camera *p.Camera, rnd *rand.Rand, samples, width, height, i, j int) p.Color {
	rgb := p.Color{}

	for s := 0; s < samples; s++ {
		u := (float64(i) + rnd.Float64()) / float64(width)
		v := (float64(j) + rnd.Float64()) / float64(height)

		ray := camera.RayAt(u, v, rnd)
		rgb = rgb.Add(color(ray, h, rnd, 0))
	}

	return rgb.DivideScalar(float64(samples))
}

// performs render 
func Do(h p.Hitable, camera *p.Camera, cpus, samples, width, height int, ch chan<- int) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	var wg sync.WaitGroup

	for i := 0; i < cpus; i++ {
		wg.Add(1)
		go func (i int)  {
			defer wg.Done()
			rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
			
			for row := i; row <height; row += cpus {
				for col := 0; col < width; col++ {
					rgb := sample(h, camera, rnd, samples, width, height, col, row)
					img.Set(col, height-row-1, rgb)
				}
				ch <- 1
			}
		}(i)
	}

	return img
}
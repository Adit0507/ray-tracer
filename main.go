package main

import (
	"fmt"
	"os"
	"ray/geometry"
)

func main() {
	sx := 500
	sy := 300

	const color = 255.99

	f, err := os.Create("out.ppm")
	defer f.Close()
	check(err, "Error opening file: %v\n")

	_, err = fmt.Fprintf(f, "P3\n%d %d\n255\n", sx, sy)
	check(err, "Error writing to file: %v\n")

	lowerLeft := geometry.Vector{-2.0, -1.0, -1.0}
	horizontal := geometry.Vector{4.0, 0.0, 0.0}
	vertical := geometry.Vector{0.0, 2.0, 0.0}
	origin := geometry.Vector{0.0, 0.0, 0.0}

	for j := sy - 1; j >= 0; j-- {
		for i := 0; i < sx; i++ {
			u := float64(i) / float64(sx)
			v := float64(j) / float64(sy)

			position := horizontal.MultiplyScaler(u).Add(vertical.MultiplyScaler(v))
			direction := lowerLeft.Add(position)

			rgb := geometry.Ray{origin, direction}.Color()
			ir := int(color * rgb.X)
			ig := int(color * rgb.Y)
			ib := int(color * rgb.Z)

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

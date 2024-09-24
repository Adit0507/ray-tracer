package main

import (
	"fmt"
	"os"
	"ray/geometry"
)

func main() {
	sx := 400
	sy := 300

	const color = 240

	f, err := os.Create("out.ppm")
	defer f.Close()
	check(err, "Error writing to file: %v\n")

	for j := sy - 1; j >= 0; j-- {
		for i := 0; i < sx; i++ {
			v := geometry.Vector{X: float64(i) / float64(sx), Y: float64(j) / float64(sy), Z: 0.2}

			r := color * v.X
			g := color * v.Y
			b := color * v.Z

			_, err = fmt.Fprintf(f, "%d %d %d\n", r, g, b)
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

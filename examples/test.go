package main

import (
	"math/rand"

	"github.com/fogleman/ln/ln"
)

func cube(x, y, z float64) ln.Shape {
	a := ln.Vector{x - 0.5, y - 0.5, z - 0.5}
	b := ln.Vector{x + 0.5, y + 0.5, z + 0.5}
	return ln.NewCube(a, b)
}

func main() {
	scene := ln.Scene{}
	n := 20
	for x := -n; x <= n; x++ {
		for y := -n; y <= n; y++ {
			z := rand.Float64() * 3
			scene.Add(cube(float64(x), float64(y), float64(z)))
			scene.Add(cube(float64(x), float64(y), float64(z+1)))
			scene.Add(cube(float64(x), float64(y), float64(z+2)))
		}
	}
	eye := ln.Vector{30, 50, 20}
	center := ln.Vector{}
	up := ln.Vector{0, 0, 1}
	paths := scene.Render(eye, center, up, 50, 1, 0.1, 100, 0.01)
	paths.Render("out.png", 1000)
}

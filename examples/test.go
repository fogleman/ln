package main

import "github.com/fogleman/ln/ln"

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
			// z := rand.Float64() * 3
			// scene.Add(cube(float64(x), float64(y), float64(z)))
			// scene.Add(cube(float64(x), float64(y), float64(z+1)))
			// scene.Add(cube(float64(x), float64(y), float64(z+2)))
		}
	}
	n = 8
	for x := -n; x <= n; x++ {
		for y := -n; y <= n; y++ {
			scene.Add(ln.NewSphere(ln.Vector{float64(x), float64(y), 0}, 0.45))
		}
	}
	// scene.Add(ln.NewSphere(ln.Vector{0, 4, 0}, 4))
	// scene.Add(ln.NewSphere(ln.Vector{-7, 0, 0}, 4))
	// scene.Add(ln.NewSphere(ln.Vector{7, 0, 0}, 4))
	eye := ln.Vector{8, 8, 1}
	center := ln.Vector{0, 0, -4.25}
	up := ln.Vector{0, 0, 1}
	width := 1024.0
	height := 1024.0
	paths := scene.Render(eye, center, up, width, height, 50, 0.1, 100, 0.01)
	paths.WriteToPNG("out.png", width, height)
	// paths.Print()
}

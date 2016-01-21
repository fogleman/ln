package main

import (
	"math"

	"github.com/fogleman/ln/ln"
)

func function(x, y float64) float64 {
	// x *= 3
	// y *= 3
	// return math.Sin(math.Sqrt(x*x + y*y))
	return math.Cos(x*y) * (x*x - y*y)
	// return -1 / (x*x + y*y)
}

func main() {
	scene := ln.Scene{}
	box := ln.Box{ln.Vector{-2, -2, -10}, ln.Vector{2, 2, 10}}
	scene.Add(ln.NewFunction(function, box))
	eye := ln.Vector{8, 8, 8}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	paths := scene.Render(eye, center, up, 50, 1, 0.1, 100, 0.01)
	paths.Render("out.png", 1024)
	// paths.Print()
}

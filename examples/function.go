package main

import "github.com/fogleman/ln/ln"

func function(x, y float64) float64 {
	return -1 / (x*x + y*y)
	// return math.Cos(x*y) * (x*x - y*y)
}

func render(matrix ln.Matrix) ln.Paths {
	scene := ln.Scene{}
	box := ln.Box{ln.Vector{-2, -2, -4}, ln.Vector{2, 2, 2}}
	shape := ln.NewFunction(function, box, ln.Below)
	scene.Add(ln.NewTransformedShape(shape, matrix))
	eye := ln.Vector{3, 0, 3}
	center := ln.Vector{1.1, 0, 0}
	up := ln.Vector{0, 0, 1}
	return scene.Render(eye, center, up, 50, 1, 0.1, 100, 0.01)
}

func main() {
	paths := render(ln.Identity())
	paths.WriteToPNG("out.png", 1024)
	// paths.Print()
}

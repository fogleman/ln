package main

import "github.com/fogleman/ln/ln"

func function(x, y float64) float64 {
	return -1 / (x*x + y*y)
	// return math.Cos(x*y) * (x*x - y*y)
}

func main() {
	scene := ln.Scene{}
	box := ln.Box{ln.Vector{-2, -2, -4}, ln.Vector{2, 2, 2}}
	scene.Add(ln.NewFunction(function, box, ln.Below))
	eye := ln.Vector{3, 0, 3}
	center := ln.Vector{1.1, 0, 0}
	up := ln.Vector{0, 0, 1}
	width := 1024.0
	height := 1024.0
	paths := scene.Render(eye, center, up, width, height, 50, 0.1, 100, 0.01)
	paths.WriteToPNG("out.png", width, height)
	paths.WriteToSVG("out.svg", width, height)
	// paths.Print()
}

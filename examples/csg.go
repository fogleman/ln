package main

import "github.com/fogleman/ln/ln"

func main() {
	scene := ln.Scene{}
	a := ln.NewSphere(ln.Vector{-1, -1, 0}, 2)
	b := ln.NewSphere(ln.Vector{1, 0, 0}, 2)
	// c := ln.NewSphere(ln.Vector{0, 0, 1.5}, 2)
	// scene.Add(ln.NewDifference(ln.NewIntersection(a, b), c))
	// scene.Add(ln.NewSphere(ln.Vector{0, -5, 0}, 2))
	scene.Add(ln.NewDifference(a, b))
	eye := ln.Vector{0, 6, 1}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	paths := scene.Render(eye, center, up, 50, 1, 0.1, 100, 0.01)
	paths.Render("out.png", 1024)
	// paths.Print()
}

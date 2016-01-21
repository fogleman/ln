package main

import (
	"fmt"
	"math"

	"github.com/fogleman/ln/ln"
)

func function(x, y float64) float64 {
	return math.Cos(x*y) * (x*x - y*y)
}

func render(matrix ln.Matrix) ln.Paths {
	scene := ln.Scene{}
	box := ln.Box{ln.Vector{-2, -2, -10}, ln.Vector{2, 2, 10}}
	shape := ln.NewFunction(function, box)
	scene.Add(ln.NewTransformedShape(shape, matrix))
	eye := ln.Vector{8, 8, 8}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	return scene.Render(eye, center, up, 50, 1, 0.1, 100, 0.01)
}

func main() {
	for i := 0; i < 360; i += 2 {
		fmt.Println(i)
		matrix := ln.Rotate(ln.Vector{0, 0, 1}, ln.Radians(float64(i)))
		paths := render(matrix)
		paths.Render(fmt.Sprintf("out%03d.png", i), 1024)
	}
}

package main

import (
	"fmt"

	"github.com/fogleman/ln/ln"
)

func main() {
	shape := ln.NewDifference(
		ln.NewIntersection(
			ln.NewSphere(ln.Vector{}, 1),
			ln.NewCube(ln.Vector{-0.8, -0.8, -0.8}, ln.Vector{0.8, 0.8, 0.8}),
		),
		ln.NewCylinder(0.4, -2, 2),
		ln.NewTransformedShape(ln.NewCylinder(0.4, -2, 2), ln.Rotate(ln.Vector{1, 0, 0}, ln.Radians(90))),
		ln.NewTransformedShape(ln.NewCylinder(0.4, -2, 2), ln.Rotate(ln.Vector{0, 1, 0}, ln.Radians(90))),
	)
	for i := 0; i < 90; i += 2 {
		fmt.Println(i)
		scene := ln.Scene{}
		m := ln.Rotate(ln.Vector{0, 0, 1}, ln.Radians(float64(i)))
		scene.Add(ln.NewTransformedShape(shape, m))
		eye := ln.Vector{0, 6, 2}
		center := ln.Vector{0, 0, 0}
		up := ln.Vector{0, 0, 1}
		width := 750.0
		height := 750.0
		paths := scene.Render(eye, center, up, width, height, 20, 0.1, 100, 0.01)
		paths.WriteToPNG(fmt.Sprintf("out%03d.png", i), width, height)
	}
}

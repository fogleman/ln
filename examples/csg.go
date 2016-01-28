package main

import (
	"fmt"

	"github.com/fogleman/ln/ln"
)

func main() {
	shape := ln.NewDifference(
		ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}),
		// ln.NewSphere(ln.Vector{}, 1),
		ln.NewSphere(ln.Vector{-1, -1, -1}, 0.5),
		ln.NewSphere(ln.Vector{1, 1, 1}, 1.5),
		ln.NewSphere(ln.Vector{-1, 0, 0}, 0.5),
		ln.NewSphere(ln.Vector{0, -1, 0}, 0.5),
	)
	for i := 0; i < 360; i += 2 {
		fmt.Println(i)
		scene := ln.Scene{}
		m := ln.Rotate(ln.Vector{0, 0, 1}, ln.Radians(float64(i)))
		scene.Add(ln.NewTransformedShape(shape, m))
		eye := ln.Vector{0, 6, 2}
		center := ln.Vector{0, 0, 0}
		up := ln.Vector{0, 0, 1}
		paths := scene.Render(eye, center, up, 35, 1, 0.1, 100, 0.01)
		paths.WriteToPNG(fmt.Sprintf("out%03d.png", i), 256)
	}
}

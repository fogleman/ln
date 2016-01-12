package main

import (
	"fmt"

	"github.com/fogleman/ln/ln"
)

func main() {
	a := ln.NewSphere(ln.Vector{-1, 0, 0}, 2)
	b := ln.NewSphere(ln.Vector{1, 0, 0}, 2)
	c := ln.NewSphere(ln.Vector{0, 0, 1.5}, 2)
	shape := ln.NewDifference(ln.NewIntersection(a, b), c)

	for i := 0; i < 360; i += 1 {
		fmt.Println(i)
		scene := ln.Scene{}
		m := ln.Rotate(ln.Vector{0, 0, 1}, ln.Radians(float64(i)))
		scene.Add(ln.NewTransformedShape(shape, m))
		eye := ln.Vector{0, 6, 1}
		center := ln.Vector{0, 0, 0}
		up := ln.Vector{0, 0, 1}
		paths := scene.Render(eye, center, up, 50, 1, 0.1, 100, 0.01)
		paths.Render(fmt.Sprintf("out%03d.png", i), 1024)
	}
}

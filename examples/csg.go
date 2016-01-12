package main

import (
	"fmt"

	"github.com/fogleman/ln/ln"
)

func main() {
	// a := ln.NewSphere(ln.Vector{-1, 0, 0}, 2)
	// b := ln.NewSphere(ln.Vector{1, 0, 0}, 2)
	// c := ln.NewSphere(ln.Vector{0, 0, 1.5}, 2)
	// shape := ln.NewDifference(ln.NewIntersection(a, b), c)
	a := ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1})
	// b := ln.NewSphere(ln.Vector{0, 0, 1}, 0.5)
	shape := ln.NewDifference(a,
		ln.NewSphere(ln.Vector{-1, -1, -1}, 0.8),
		ln.NewSphere(ln.Vector{1, 1, 1}, 0.8),
		ln.NewSphere(ln.Vector{0, 1, 0}, 0.5),
		ln.NewSphere(ln.Vector{0, -1, 0}, 0.5))
	// ln.NewSphere(ln.Vector{0, -1, 0}, 0.4),
	// ln.NewSphere(ln.Vector{1, 0, 0}, 0.4),
	// ln.NewSphere(ln.Vector{-1, 0, 0}, 0.4))
	for i := 0; i < 360; i += 1 {
		fmt.Println(i)
		scene := ln.Scene{}
		m := ln.Rotate(ln.Vector{0, 0, 1}, ln.Radians(float64(i)))
		scene.Add(ln.NewTransformedShape(shape, m))
		eye := ln.Vector{0, 6, 2}
		center := ln.Vector{0, 0, 0}
		up := ln.Vector{0, 0, 1}
		paths := scene.Render(eye, center, up, 40, 1, 0.1, 100, 0.01)
		paths.Render(fmt.Sprintf("out%03d.png", i), 256)
	}
}

package main

import (
	"math/rand"

	"github.com/fogleman/ln/ln"
)

func main() {
	scene := ln.Scene{}
	n := 15
	for x := -n; x <= n; x++ {
		for y := -n; y <= n; y++ {
			p := rand.Float64()*0.25 + 0.2
			dx := rand.Float64()*0.5 - 0.25
			dy := rand.Float64()*0.5 - 0.25
			fx := float64(x) + dx*0
			fy := float64(y) + dy*0
			fz := rand.Float64()*3 + 1
			shape := ln.NewCube(ln.Vector{fx - p, fy - p, 0}, ln.Vector{fx + p, fy + p, fz})
			if x == 2 && y == 1 {
				continue
			}
			scene.Add(shape)
		}
	}
	eye := ln.Vector{1.75, 1.25, 6}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	width := 1024.0
	height := 1024.0
	paths := scene.Render(eye, center, up, width, height, 100, 0.1, 100, 0.01)
	paths.WriteToPNG("out.png", width, height)
	// paths.Print()
}

package main

import (
	"math/rand"

	"github.com/fogleman/ln/ln"
)

func main() {
	eye := ln.Vector{8, 8, 8}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	scene := ln.Scene{}
	n := 10
	for x := -n; x <= n; x++ {
		for y := -n; y <= n; y++ {
			z := rand.Float64() * 3
			v := ln.Vector{float64(x), float64(y), z}
			sphere := ln.NewOutlineSphere(eye, up, v, 0.45)
			scene.Add(sphere)
		}
	}
	width := 1920.0
	height := 1200.0
	fovy := 50.0
	paths := scene.Render(eye, center, up, width, height, fovy, 0.1, 100, 0.01)
	paths.WriteToPNG("out.png", width, height)
}

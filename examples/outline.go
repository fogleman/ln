package main

import (
	"math/rand"

	"github.com/fogleman/ln/ln"
)

func main() {
	scene := ln.Scene{}
	eye := ln.Vector{8, 8, 8}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	n := 10
	for x := -n; x <= n; x++ {
		for y := -n; y <= n; y++ {
			z := rand.Float64() * 3
			v := ln.Vector{float64(x), float64(y), z}
			sphere := ln.NewOutlineSphere(eye, up, v, 0.45)
			scene.Add(sphere)
		}
	}
	paths := scene.Render(eye, center, up, 50, 1, 0.1, 100, 0.01)
	paths.Render("out.png", 1024)
}

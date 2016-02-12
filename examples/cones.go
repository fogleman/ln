package main

import (
	"math"
	"math/rand"

	"github.com/fogleman/ln/ln"
)

type Tree struct {
	ln.Shape
	V0, V1 ln.Vector
}

func (t *Tree) Paths() ln.Paths {
	paths := t.Shape.Paths()
	for i := 0; i < 128; i++ {
		p := math.Pow(rand.Float64(), 1.5)*0.5 + 0.5
		c := t.V0.Add(t.V1.Sub(t.V0).MulScalar(p))
		a := rand.Float64() * 2 * math.Pi
		l := (1 - p) * 8
		d := ln.Vector{math.Cos(a), math.Sin(a), -3}.Normalize()
		e := c.Add(d.MulScalar(l))
		paths = append(paths, ln.Path{c, e})
	}
	return paths
}

func main() {
	rand.Seed(111)
	eye := ln.Vector{}
	center := ln.Vector{0.5, 0, 8}
	up := ln.Vector{0, 0, 1}
	scene := ln.Scene{}
	n := 9
	for x := -n; x <= n; x += 3 {
		for y := -n; y <= n; y += 3 {
			if x == 0 && y == 0 {
				continue
			}
			z := rand.Float64()*5 + 20
			xx := float64(x) + (rand.Float64()*2-1)*1
			yy := float64(y) + (rand.Float64()*2-1)*1
			v0 := ln.Vector{xx, yy, 0}
			v1 := ln.Vector{xx, yy, z}
			c := ln.NewTransformedOutlineCone(eye, up, v0, v1, z/64)
			tree := Tree{c, v0, v1}
			scene.Add(&tree)
		}
	}
	width := 1024.0
	height := 1024.0
	fovy := 90.0
	paths := scene.Render(eye, center, up, width, height, fovy, 0.1, 100, 0.01)
	paths.WriteToPNG("out.png", width, height)
}

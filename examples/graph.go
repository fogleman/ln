package main

import (
	"math"

	"github.com/fogleman/ln/ln"
)

func main() {
	scene := ln.Scene{}
	eye := ln.Vector{1, 1, 1}.MulScalar(1.5)
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}

	nodes := []ln.Vector{
		{0, 0, 0},
		{-1, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
		{0, -1, 0},
		{0, 0, 1},
		{0, 0, -1},
	}

	// edges := [][2]int{
	// 	// {3, 4},
	// 	{0, 1},
	// 	{0, 2},
	// 	{1, 3},
	// 	{2, 3},
	// 	{0, 3},
	// 	{0, 4},
	// }

	for _, v := range nodes {
		scene.Add(ln.NewOutlineSphere(eye, up, v, 0.25))
	}

	for _, v0 := range nodes {
		for _, v1 := range nodes {
			if v0 == v1 {
				continue
			}
			// v0 := nodes[edge[0]]
			// v1 := nodes[edge[1]]
			d := v1.Sub(v0)
			z := d.Length()
			u := d.Cross(up).Normalize()
			a := math.Acos(d.Normalize().Dot(up))
			m := ln.Translate(v0)
			if a != 0 {
				m = ln.Rotate(u, a).Translate(v0)
			}
			c := ln.NewOutlineCylinder(m.Inverse().MulPosition(eye), up, 0.1/2, 0, z)
			scene.Add(ln.NewTransformedShape(c, m))
		}
	}

	width := 1024.0
	height := 1024.0
	paths := scene.Render(eye, center, up, width, height, 60, 0.1, 100, 0.01)
	paths.WriteToPNG("out.png", width, height)
}

package main

import (
	"fmt"
	"math"

	"github.com/fogleman/ln/ln"
)

func render(frame int) {
	cx := math.Cos(ln.Radians(float64(frame)))
	cy := math.Sin(ln.Radians(float64(frame)))
	scene := ln.Scene{}
	eye := ln.Vector{cx, cy, 0}.MulScalar(8)
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}

	nodes := []ln.Vector{
		{1.047, -0.000, -1.312},
		{-0.208, -0.000, -1.790},
		{2.176, 0.000, -2.246},
		{1.285, -0.001, 0.016},
		{-1.276, -0.000, -0.971},
		{-0.384, 0.000, -2.993},
		{-2.629, -0.000, -1.533},
		{-1.098, -0.000, 0.402},
		{0.193, 0.005, 0.911},
		{-1.934, -0.000, 1.444},
		{2.428, -0.000, 0.437},
		{0.068, -0.000, 2.286},
		{-1.251, -0.000, 2.560},
		{1.161, -0.000, 3.261},
		{1.800, 0.001, -3.269},
		{2.783, 0.890, -2.082},
		{2.783, -0.889, -2.083},
		{-2.570, -0.000, -2.622},
		{-3.162, -0.890, -1.198},
		{-3.162, 0.889, -1.198},
		{-1.679, 0.000, 3.552},
		{1.432, -1.028, 3.503},
		{2.024, 0.513, 2.839},
		{0.839, 0.513, 4.167},
		// {0.000000, 0.000000, 0.000000},
		// {0.000000, 0.000000, 1.089000},
		// {1.026719, 0.000000, -0.363000},
		// {-0.513360, -0.889165, -0.363000},
		// {-0.513360, 0.889165, -0.363000},
		//
		// {0, 0, 0},
		// {-1, 0, 0},
		// {1, 0, 0},
		// {0, 1, 0},
		// {0, -1, 0},
		// {0, 0, 1},
		// {0, 0, -1},
		//
		// {-1, 1, 1},
		// {-1, 1, -1},
		// {-1, -1, 1},
		// {-1, -1, -1},
		// {1, 1, 1},
		// {1, 1, -1},
		// {1, -1, 1},
		// {1, -1, -1},
	}

	edges := [][2]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 4},
		{1, 5},
		{2, 14},
		{2, 15},
		{2, 16},
		{3, 8},
		{3, 10},
		{4, 6},
		{4, 7},
		{6, 17},
		{6, 18},
		{6, 19},
		{7, 8},
		{7, 9},
		{8, 11},
		{9, 12},
		{11, 12},
		{11, 13},
		{12, 20},
		{13, 21},
		{13, 22},
		{13, 23},
	}

	for _, v := range nodes {
		scene.Add(ln.NewOutlineSphere(eye, up, v, 0.333))
	}

	// for _, v0 := range nodes {
	// 	for _, v1 := range nodes {
	// 		if v0 == v1 {
	// 			continue
	// 		}
	for _, edge := range edges {
		v0 := nodes[edge[0]]
		v1 := nodes[edge[1]]
		d := v1.Sub(v0)
		z := d.Length()
		u := d.Cross(up).Normalize()
		a := math.Acos(d.Normalize().Dot(up))
		m := ln.Translate(v0)
		if a != 0 {
			m = ln.Rotate(u, a).Translate(v0)
		}
		c := ln.NewOutlineCylinder(m.Inverse().MulPosition(eye), up, 0.1, 0, z)
		scene.Add(ln.NewTransformedShape(c, m))
	}
	// }

	width := 750.0
	height := 750.0
	paths := scene.Render(eye, center, up, width, height, 60, 0.1, 100, 0.01)
	paths.WriteToPNG(fmt.Sprintf("out%03d.png", frame), width, height)
}

func main() {
	for i := 0; i < 360; i += 2 {
		render(i)
	}
}

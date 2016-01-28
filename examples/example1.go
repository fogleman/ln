package main

import "github.com/fogleman/ln/ln"

func cube(x, y, z float64) ln.Shape {
	size := 0.5
	v := ln.Vector{x, y, z}
	return ln.NewCube(v.SubScalar(size), v.AddScalar(size))
}

func main() {
	scene := ln.Scene{}
	scene.Add(cube(0, 0, 0))
	scene.Add(cube(-1, 0, 0))
	scene.Add(cube(1, 0, 0))
	scene.Add(cube(0, -1, 0))
	scene.Add(cube(0, 1, 0))
	scene.Add(cube(0, 0, -1))
	scene.Add(cube(0, 0, 1))
	eye := ln.Vector{6, 5, 3}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	width := 1024.0
	height := 1024.0
	fovy := 50.0
	paths := scene.Render(eye, center, up, width, height, fovy, 0.1, 100, 0.01)
	paths.Render("out.png", width, height, 1)
}

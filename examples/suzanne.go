package main

import "github.com/fogleman/ln/ln"

func main() {
	scene := ln.Scene{}
	mesh, err := ln.LoadOBJ("examples/suzanne.obj")
	if err != nil {
		panic(err)
	}
	mesh.UnitCube()
	scene.Add(ln.NewTransformedShape(mesh, ln.Rotate(ln.Vector{0, 1, 0}, 0.5)))
	// scene.Add(mesh)
	eye := ln.Vector{-0.5, 0.5, 2}
	center := ln.Vector{}
	up := ln.Vector{0, 1, 0}
	paths := scene.Render(eye, center, up, 35, 1, 0.1, 100, 0.01)
	paths.Render("out.png", 1000)
}

package main

import "github.com/fogleman/ln/ln"

func main() {
	scene := ln.Scene{}
	mesh, err := ln.LoadBinarySTL("bowser.stl")
	// mesh, err := ln.LoadOBJ("../pt/examples/bunny.obj")
	if err != nil {
		panic(err)
	}
	mesh.FitInside(ln.Box{ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}}, ln.Vector{0.5, 0.5, 0.5})
	cubes := mesh.Voxelize(1.0 / 64)
	for _, cube := range cubes {
		scene.Add(cube)
	}
	eye := ln.Vector{-1, -2, 0}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	width := 1024.0 * 2
	height := 1024.0 * 2
	paths := scene.Render(eye, center, up, width, height, 60, 0.1, 100, 0.01)
	paths.WriteToPNG("out.png", width, height)
	// paths.WriteToSVG("out.svg", width, height)
	// paths.Print()
}

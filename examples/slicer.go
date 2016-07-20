package main

import (
	"fmt"

	"github.com/fogleman/ln/ln"
)

const Slices = 32
const Size = 1024

func main() {
	// mesh, err := ln.LoadBinarySTL("bowser.stl")
	mesh, err := ln.LoadOBJ("examples/suzanne.obj")
	if err != nil {
		panic(err)
	}
	mesh.FitInside(ln.Box{ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}}, ln.Vector{0.5, 0.5, 0.5})
	for i := 0; i < Slices; i++ {
		fmt.Printf("slice%04d\n", i)
		p := (float64(i)/(Slices-1))*2 - 1
		point := ln.Vector{0, 0, p}
		plane := ln.Plane{point, ln.Vector{0, 0, 1}}
		paths := plane.IntersectMesh(mesh)
		paths = paths.Transform(ln.Scale(ln.Vector{Size / 2, Size / 2, 1}).Translate(ln.Vector{Size / 2, Size / 2, 0}))
		paths.WriteToPNG(fmt.Sprintf("slice%04d.png", i), Size, Size)
		// paths.WriteToSVG(fmt.Sprintf("slice%04d.svg", i), Size, Size)
	}
}

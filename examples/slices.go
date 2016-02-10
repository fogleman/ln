package main

import (
	"fmt"

	"github.com/fogleman/ln/ln"
)

type Shape struct {
	ln.Mesh
}

func (s *Shape) Paths() ln.Paths {
	var result ln.Paths
	// for i := 0; i < 360; i++ {
	// 	fmt.Println(i)
	// 	a := ln.Radians(float64(i))
	// 	x := math.Cos(a)
	// 	y := math.Sin(a)
	// 	plane := ln.Plane{ln.Vector{}, ln.Vector{x, y, 0}}
	// 	paths := plane.IntersectMesh(&s.Mesh)
	// 	result = append(result, paths...)
	// }
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
		p := float64(i) / 100
		plane := ln.Plane{ln.Vector{0, 0, p*2 - 1}, ln.Vector{0, 0, 1}}
		result = append(result, plane.IntersectMesh(&s.Mesh)...)
		plane = ln.Plane{ln.Vector{p*2 - 1, 0, 0}, ln.Vector{1, 0, 0}}
		result = append(result, plane.IntersectMesh(&s.Mesh)...)
		plane = ln.Plane{ln.Vector{0, p*2 - 1, 0}, ln.Vector{0, 1, 0}}
		result = append(result, plane.IntersectMesh(&s.Mesh)...)
	}
	return result
}

func main() {
	scene := ln.Scene{}
	mesh, err := ln.LoadBinarySTL("bowser.stl")
	// mesh, err := ln.LoadOBJ("../pt/examples/bunny.obj")
	if err != nil {
		panic(err)
	}
	mesh.FitInside(ln.Box{ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}}, ln.Vector{0.5, 0.5, 0.5})
	scene.Add(&Shape{*mesh})
	// scene.Add(mesh)
	eye := ln.Vector{-2, 2, 1}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	width := 1024.0 * 2
	height := 1024.0 * 2
	paths := scene.Render(eye, center, up, width, height, 50, 0.1, 100, 0.01)
	paths.WriteToPNG("out.png", width, height)
	// paths.WriteToSVG("out.svg", width, height)
	// paths.Print()
}

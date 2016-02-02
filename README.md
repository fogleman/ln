# `ln` The 3D Line Art Engine

`ln` is a vector-based 3D renderer written in Go. It is used to produce 2D
vector graphics (think SVGs) depicting 3D scenes.

*The output of an OpenGL pipeline is a rastered image. The output of `ln` is
a set of 2D vector paths.*

![Examples](http://i.imgur.com/HY2Fg2t.png)

## Motivation

I created this so I could plot 3D drawings with my
[Makeblock XY Plotter](http://www.makeblock.cc/xy-plotter-robot-kit/).

Here's one of my drawings from the plotter...

![Example](http://i.imgur.com/NbgpUhQ.jpg)

## Dependencies

`ln` uses cairo for rendering.

### OS X

	brew install cairo pkg-config
	go get github.com/ungerik/go-cairo

### Ubuntu

	sudo apt-get install libcairo2-dev pkg-config
	go get github.com/ungerik/go-cairo

## Features

- Primitives
	- Sphere
	- Cube
	- Triangle
	- Cylinder
	- 3D Functions
- Triangle Meshes
	- OBJ & STL
- Vector-based "Texturing"
- CSG (Constructive Solid Geometry) Operations
	- Intersection
	- Difference
	- Union
- Output to PNG or SVG

## How it Works

To understand how `ln` works, it's useful to start with the `Shape` interface:

```go
type Shape interface {
	Paths() Paths
	Intersect(Ray) Hit
	Contains(Vector, float64) bool
	BoundingBox() Box
	Compile()
}
```

Each shape must provide some `Paths` which are 3D polylines on the surface
of the solid. Ultimately anything drawn in the final image is based on these
paths. These paths can be anything. For a sphere they could be lat/lng grid
lines, a triangulated-looking surface, dots on the surface, etc. This is what
we call vector-based texturing. Each built-in `Shape` ships with a default
`Paths` function (e.g. a `Cube` simply draws the outline of a cube) but you
can easily provide your own.

Each shape must also provide an `Intersect` method that lets the engine test
for ray-solid intersection. This is how the engine knows what is visible to the
camera and what is hidden.

All of the `Paths` are chopped up to some granularity and each point is tested
by shooting a ray toward the camera. If there is no intersection, that point is
visible. If there is an intersection, it is hidden and will not be rendered.

The visible points are then transformed into 2D space using transformation
matrices. The result can then be rendered as PNG or SVG.

The `Contains` method is only needed for CSG (Constructive Solid Geometry)
operations.

## Hello World: A Single Cube

### The Code

```go
package main

import "github.com/fogleman/ln/ln"

func main() {
	// create a scene and add a single cube
	scene := ln.Scene{}
	scene.Add(ln.NewCube(ln.Vector{-1, -1, -1}, ln.Vector{1, 1, 1}))

	// define camera parameters
	eye := ln.Vector{4, 3, 2}    // camera position
	center := ln.Vector{0, 0, 0} // camera looks at
	up := ln.Vector{0, 0, 1}     // up direction

	// define rendering parameters
	width := 1024.0  // rendered width
	height := 1024.0 // rendered height
	fovy := 50.0     // vertical field of view, degrees
	znear := 0.1     // near z plane
	zfar := 10.0     // far z plane
	step := 0.01     // how finely to chop the paths for visibility testing

	// compute 2D paths that depict the 3D scene
	paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

	// render the paths in an image
	paths.WriteToPNG("out.png", width, height)

	// save the paths as an svg
	paths.WriteToSVG("out.svg", width, height)
}
```

### The Output

![Cube](http://i.imgur.com/d2dGrOJ.png)

## Custom Texturing

Suppose we want to draw cubes with vertical stripes on their sides, as
shown in the skyscrapers example above. We can just define a new type
and override the `Paths()` function.

```go
type StripedCube struct {
	ln.Cube
}

func (c *StripedCube) Paths() ln.Paths {
	var paths ln.Paths
	x1, y1, z1 := c.Min.X, c.Min.Y, c.Min.Z
	x2, y2, z2 := c.Max.X, c.Max.Y, c.Max.Z
	for i := 0; i <= 10; i++ {
		p := float64(i) / 10
		x := x1 + (x2-x1)*p
		y := y1 + (y2-y1)*p
		paths = append(paths, ln.Path{{x, y1, z1}, {x, y1, z2}})
		paths = append(paths, ln.Path{{x, y2, z1}, {x, y2, z2}})
		paths = append(paths, ln.Path{{x1, y, z1}, {x1, y, z2}})
		paths = append(paths, ln.Path{{x2, y, z1}, {x2, y, z2}})
	}
	return paths
}
```

## Constructive Solid Geometry (CSG)

You can easily construct complex solids using Intersection, Difference, Union.

```go
shape := ln.NewDifference(
	ln.NewIntersection(
		ln.NewSphere(ln.Vector{}, 1),
		ln.NewCube(ln.Vector{-0.8, -0.8, -0.8}, ln.Vector{0.8, 0.8, 0.8}),
	),
	ln.NewCylinder(0.4, -2, 2),
	ln.NewTransformedShape(ln.NewCylinder(0.4, -2, 2), ln.Rotate(ln.Vector{1, 0, 0}, ln.Radians(90))),
	ln.NewTransformedShape(ln.NewCylinder(0.4, -2, 2), ln.Rotate(ln.Vector{0, 1, 0}, ln.Radians(90))),
)
```

This is `(Sphere & Cube) - (Cylinder | Cylinder | Cylinder)`.

Unfortunately, it's difficult to compute the joint formed at the boundaries of these combined shapes, so sufficient texturing is needed on the original solids for a decent result.

![Example](http://i.imgur.com/gk8UtVK.gif)

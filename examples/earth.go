package main

import (
	"github.com/fogleman/ln/ln"
	"github.com/jonas-p/go-shp"
)

func GetPaths(shape shp.Shape) ln.Paths {
	switch v := shape.(type) {
	case *shp.PolyLine:
		return getPaths(v)
	case *shp.Polygon:
		line := shp.PolyLine(*v)
		return getPaths(&line)
	}
	return nil
}

func getPaths(line *shp.PolyLine) ln.Paths {
	var result ln.Paths
	parts := append(line.Parts, line.NumPoints)
	for part := 0; part < len(parts)-1; part++ {
		var path ln.Path
		a := parts[part]
		b := parts[part+1]
		for i := a; i < b; i++ {
			pt := line.Points[i]
			path = append(path, ln.LatLngToXYZ(pt.Y, pt.X, 1))
		}
		result = append(result, path)
	}
	return result
}

func LoadPaths() ln.Paths {
	var result ln.Paths
	file, err := shp.Open("examples/ne_10m_coastline/ne_10m_coastline.shp")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for file.Next() {
		_, shape := file.Shape()
		paths := GetPaths(shape)
		result = append(result, paths...)
	}
	return result
}

type Earth struct {
	ln.Shape
	Lines ln.Paths
}

func (e *Earth) Paths() ln.Paths {
	// return append(e.Lines, e.Shape.Paths()...)
	return e.Lines
}

func Render(lines ln.Paths, matrix ln.Matrix) ln.Paths {
	scene := ln.Scene{}
	sphere := ln.NewSphere(ln.Vector{}, 1)
	earth := Earth{sphere, lines}
	shape := ln.NewTransformedShape(&earth, matrix)
	scene.Add(shape)
	eye := ln.LatLngToXYZ(35.7806, -78.6389, 1).Normalize().MulScalar(3.25)
	center := ln.Vector{}
	up := ln.Vector{0, 0, 1}
	return scene.Render(eye, center, up, 45, 1, 0.1, 100, 0.01)
}

func main() {
	lines := LoadPaths()
	m := ln.Identity()
	paths := Render(lines, m)
	paths.Render("earth.png", 1024)
	paths.Print()
}

package ln

import "math"

type Function struct {
	Function func(x, y float64) float64
	Box      Box
}

func NewFunction(function func(x, y float64) float64, box Box) Shape {
	return &Function{function, box}
}

func (f *Function) Compile() {
}

func (f *Function) BoundingBox() Box {
	return f.Box
}

func (f *Function) Contains(v Vector, eps float64) bool {
	return false
}

func (f *Function) Intersect(ray Ray) Hit {
	step := 1.0 / 64
	sign := f.Test(ray.Position(step))
	for t := step; t < 10; t += step {
		v := ray.Position(t)
		if f.Test(v) != sign && f.Box.Contains(v) {
			return Hit{f, t}
		}
	}
	return NoHit
}

func (f *Function) Test(v Vector) bool {
	return f.Function(v.X, v.Y) > v.Z
}

func (f *Function) Paths() Paths {
	var paths Paths
	step := 1.0 / 16
	fine := 1.0 / 64
	for x := f.Box.Min.X; x <= f.Box.Max.X; x += step {
		var path Path
		for y := f.Box.Min.Y; y <= f.Box.Max.Y; y += fine {
			z := f.Function(x, y)
			z = math.Min(z, f.Box.Max.Z)
			z = math.Max(z, f.Box.Min.Z)
			path = append(path, Vector{x, y, z})
		}
		paths = append(paths, path)
	}
	for y := f.Box.Min.Y; y <= f.Box.Max.Y; y += step {
		var path Path
		for x := f.Box.Min.X; x <= f.Box.Max.X; x += fine {
			z := f.Function(x, y)
			z = math.Min(z, f.Box.Max.Z)
			z = math.Max(z, f.Box.Min.Z)
			path = append(path, Vector{x, y, z})
		}
		paths = append(paths, path)
	}
	return paths
}

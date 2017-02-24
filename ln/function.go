package ln

import "math"

type Direction int

const (
	Above Direction = iota
	Below
)

type Function struct {
	Function  func(x, y float64) float64
	Box       Box
	Direction Direction
}

func NewFunction(function func(x, y float64) float64, box Box, direction Direction) Shape {
	return &Function{function, box, direction}
}

func (f *Function) Compile() {
}

func (f *Function) BoundingBox() Box {
	return f.Box
}

func (f *Function) Contains(v Vector, eps float64) bool {
	if f.Direction == Below {
		return v.Z < f.Function(v.X, v.Y)
	} else {
		return v.Z > f.Function(v.X, v.Y)
	}
}

func (f *Function) Intersect(ray Ray) Hit {
	step := 1.0 / 64
	sign := f.Contains(ray.Position(step), 0)
	for t := step; t < 10; t += step {
		v := ray.Position(t)
		if f.Contains(v, 0) != sign && f.Box.Contains(v) {
			return Hit{f, t}
		}
	}
	return NoHit
}

type SpiralFunction struct {
	Function
}

func NewSpiralFunction(function func(x, y float64) float64, box Box, direction Direction) Shape {
	return &SpiralFunction{Function{function, box, direction}}
}

func (f *SpiralFunction) Paths() Paths {
	var path Path
	n := 10000
	for i := 0; i < n; i++ {
		t := float64(i) / float64(n)
		r := 8 - math.Pow(t, 0.1)*8
		x := math.Cos(Radians(t*2*math.Pi*3000)) * r
		y := math.Sin(Radians(t*2*math.Pi*3000)) * r
		z := f.Function.Function(x, y)
		z = math.Min(z, f.Box.Max.Z)
		z = math.Max(z, f.Box.Min.Z)
		path = append(path, Vector{x, y, z})
	}
	return Paths{path}
}

type RadialFunction struct {
	Function
}

func NewRadialFunction(function func(x, y float64) float64, box Box, direction Direction) Shape {
	return &RadialFunction{Function{function, box, direction}}
}

func (f *RadialFunction) Paths() Paths {
	var paths Paths
	fine := 1.0 / 256
	for a := 0; a < 360; a += 5 {
		var path Path
		for r := 0.0; r <= 8.0; r += fine {
			x := math.Cos(Radians(float64(a))) * r
			y := math.Sin(Radians(float64(a))) * r
			z := f.Function.Function(x, y)
			o := -math.Pow(-z, 1.4)
			x = math.Cos(Radians(float64(a))-o) * r
			y = math.Sin(Radians(float64(a))-o) * r
			z = math.Min(z, f.Box.Max.Z)
			z = math.Max(z, f.Box.Min.Z)
			path = append(path, Vector{x, y, z})
		}
		paths = append(paths, path)
	}
	return paths
}

// Square grid, default for function base type.
func (f *Function) Paths() Paths {
	var paths Paths
	step := 1.0 / 8
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

type WavySpiralFunction struct {
	Function
}

func NewWavySpiralFunction(function func(x, y float64) float64, box Box, direction Direction) Shape {
	return &WavySpiralFunction{Function{function, box, direction}}
}

func (f *WavySpiralFunction) Paths() Paths {
	var path Path
	n := 50000
	for i := 0; i < n; i++ {
		t := float64(i) / float64(n)
                angle := Radians(t*2*math.Pi*3000)

                cos := math.Cos(angle)
                sin := math.Sin(angle)

                pow := math.Pow(t, 0.1)

                wave_angle := angle * 100 * (1-math.Pow(t,0.99))
                wave := math.Cos(wave_angle) * 0.1 * (1-pow)

		r := 8 - pow*(8 + wave)

		x := cos * r
		y := sin * r

		z := f.Function.Function(x, y)
		z = math.Min(z, f.Box.Max.Z)
		z = math.Max(z, f.Box.Min.Z)
		path = append(path, Vector{x, y, z})
	}
	return Paths{path}
}

type DoubleSpiralsFunction struct {
	SpiralFunction
	radial RadialFunction
}

func NewDoubleSpiralsFunction(function func(x, y float64) float64, box Box, direction Direction) Shape {
	return &DoubleSpiralsFunction{SpiralFunction{Function{function, box, direction}},
				      RadialFunction{Function{function, box, direction}}}
}

func (f *DoubleSpiralsFunction) Paths() Paths {
        return append(f.SpiralFunction.Paths(), f.radial.Paths()...)
}


package main

import (
	"math"
	"math/rand"

	"github.com/fogleman/ln/ln"
)

func main() {
	rand.Seed(1211)
	eye := ln.Vector{8, 8, 8}
	center := ln.Vector{0, 0, 0}
	up := ln.Vector{0, 0, 1}
	scene := ln.Scene{}
	for a := 0; a < 50; a++ {
		n := 200
		xs := LowPassNoise(n, 0.3, 4)
		ys := LowPassNoise(n, 0.3, 4)
		zs := LowPassNoise(n, 0.3, 4)
		ss := LowPassNoise(n, 0.3, 4)
		position := ln.Vector{}
		for i := 0; i < n; i++ {
			sphere := ln.NewOutlineSphere(eye, up, position, 0.1)
			scene.Add(sphere)
			s := (ss[i]+1)/2*0.1 + 0.01
			v := ln.Vector{xs[i], ys[i], zs[i]}.Normalize().MulScalar(s)
			position = position.Add(v)
		}
	}
	width := 380.0 * 5
	height := 315.0 * 5
	fovy := 50.0
	paths := scene.Render(eye, center, up, width, height, fovy, 0.1, 100, 0.01)
	paths.WriteToPNG("out.png", width, height)
	paths.Print()
}

func Normalize(values []float64, a, b float64) []float64 {
	result := make([]float64, len(values))
	lo := values[0]
	hi := values[0]
	for _, x := range values {
		lo = math.Min(lo, x)
		hi = math.Max(hi, x)
	}
	for i, x := range values {
		p := (x - lo) / (hi - lo)
		result[i] = a + p*(b-a)
	}
	return result
}

func LowPass(values []float64, alpha float64) []float64 {
	result := make([]float64, len(values))
	var y float64
	for i, x := range values {
		y -= alpha * (y - x)
		result[i] = y
	}
	return result
}

func LowPassNoise(n int, alpha float64, iterations int) []float64 {
	result := make([]float64, n)
	for i := range result {
		result[i] = rand.Float64()
	}
	for i := 0; i < iterations; i++ {
		result = LowPass(result, alpha)
	}
	result = Normalize(result, -1, 1)
	return result
}

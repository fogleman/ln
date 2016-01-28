package ln

import (
	"math"
	"math/rand"
)

type Vector struct {
	X, Y, Z float64
}

func RandomUnitVector() Vector {
	for {
		x := rand.Float64()*2 - 1
		y := rand.Float64()*2 - 1
		z := rand.Float64()*2 - 1
		if x*x+y*y+z*z > 1 {
			continue
		}
		return Vector{x, y, z}.Normalize()
	}
}

func (a Vector) Length() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

func (a Vector) Distance(b Vector) float64 {
	return a.Sub(b).Length()
}

func (a Vector) LengthSquared() float64 {
	return a.X*a.X + a.Y*a.Y + a.Z*a.Z
}

func (a Vector) DistanceSquared(b Vector) float64 {
	return a.Sub(b).LengthSquared()
}

func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func (a Vector) Cross(b Vector) Vector {
	x := a.Y*b.Z - a.Z*b.Y
	y := a.Z*b.X - a.X*b.Z
	z := a.X*b.Y - a.Y*b.X
	return Vector{x, y, z}
}

func (a Vector) Normalize() Vector {
	d := a.Length()
	return Vector{a.X / d, a.Y / d, a.Z / d}
}

func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func (a Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

func (a Vector) Mul(b Vector) Vector {
	return Vector{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

func (a Vector) Div(b Vector) Vector {
	return Vector{a.X / b.X, a.Y / b.Y, a.Z / b.Z}
}

func (a Vector) AddScalar(b float64) Vector {
	return Vector{a.X + b, a.Y + b, a.Z + b}
}

func (a Vector) SubScalar(b float64) Vector {
	return Vector{a.X - b, a.Y - b, a.Z - b}
}

func (a Vector) MulScalar(b float64) Vector {
	return Vector{a.X * b, a.Y * b, a.Z * b}
}

func (a Vector) DivScalar(b float64) Vector {
	return Vector{a.X / b, a.Y / b, a.Z / b}
}

func (a Vector) Min(b Vector) Vector {
	return Vector{math.Min(a.X, b.X), math.Min(a.Y, b.Y), math.Min(a.Z, b.Z)}
}

func (a Vector) Max(b Vector) Vector {
	return Vector{math.Max(a.X, b.X), math.Max(a.Y, b.Y), math.Max(a.Z, b.Z)}
}

func (a Vector) MinAxis() Vector {
	x, y, z := math.Abs(a.X), math.Abs(a.Y), math.Abs(a.Z)
	switch {
	case x <= y && x <= z:
		return Vector{1, 0, 0}
	case y <= x && y <= z:
		return Vector{0, 1, 0}
	}
	return Vector{0, 0, 1}
}

func (a Vector) MinComponent() float64 {
	return math.Min(math.Min(a.X, a.Y), a.Z)
}

func (p Vector) SegmentDistance(v Vector, w Vector) float64 {
	l2 := v.DistanceSquared(w)
	if l2 == 0 {
		return p.Distance(v)
	}
	t := p.Sub(v).Dot(w.Sub(v)) / l2
	if t < 0 {
		return p.Distance(v)
	}
	if t > 1 {
		return p.Distance(w)
	}
	return v.Add(w.Sub(v).MulScalar(t)).Distance(p)
}

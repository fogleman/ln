package ln

import "math"

type Cube struct {
	Min Vector
	Max Vector
	Box Box
}

func NewCube(min, max Vector) Shape {
	box := Box{min, max}
	return &Cube{min, max, box}
}

func (c *Cube) Compile() {
}

func (c *Cube) BoundingBox() Box {
	return c.Box
}

func (c *Cube) Intersect(r Ray) Hit {
	n := c.Min.Sub(r.Origin).Div(r.Direction)
	f := c.Max.Sub(r.Origin).Div(r.Direction)
	n, f = n.Min(f), n.Max(f)
	t0 := math.Max(math.Max(n.X, n.Y), n.Z)
	t1 := math.Min(math.Min(f.X, f.Y), f.Z)
	if t0 < 0 && t1 > 0 {
		return Hit{c, 0}
	}
	if t0 >= 0 && t0 < t1 {
		return Hit{c, t0}
	}
	return NoHit
}

func (c *Cube) Paths() Paths {
	x1, y1, z1 := c.Min.X, c.Min.Y, c.Min.Z
	x2, y2, z2 := c.Max.X, c.Max.Y, c.Max.Z
	return Paths{
		{Vector{x1, y1, z1}, Vector{x1, y1, z2}},
		{Vector{x1, y1, z1}, Vector{x1, y2, z1}},
		{Vector{x1, y1, z1}, Vector{x2, y1, z1}},
		{Vector{x1, y1, z2}, Vector{x1, y2, z2}},
		{Vector{x1, y1, z2}, Vector{x2, y1, z2}},
		{Vector{x1, y2, z1}, Vector{x1, y2, z2}},
		{Vector{x1, y2, z1}, Vector{x2, y2, z1}},
		{Vector{x1, y2, z2}, Vector{x2, y2, z2}},
		{Vector{x2, y1, z1}, Vector{x2, y1, z2}},
		{Vector{x2, y1, z1}, Vector{x2, y2, z1}},
		{Vector{x2, y1, z2}, Vector{x2, y2, z2}},
		{Vector{x2, y2, z1}, Vector{x2, y2, z2}},
	}
}

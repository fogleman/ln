package ln

import "math"

type Cube struct {
	Min Vector
	Max Vector
	Box Box
}

func NewCube(min, max Vector) *Cube {
	box := Box{min, max}
	return &Cube{min, max, box}
}

func (c *Cube) Compile() {
}

func (c *Cube) BoundingBox() Box {
	return c.Box
}

func (c *Cube) Contains(v Vector, f float64) bool {
	if v.X < c.Min.X-f || v.X > c.Max.X+f {
		return false
	}
	if v.Y < c.Min.Y-f || v.Y > c.Max.Y+f {
		return false
	}
	if v.Z < c.Min.Z-f || v.Z > c.Max.Z+f {
		return false
	}
	return true
}

func (c *Cube) Intersect(r Ray) Hit {
	n := c.Min.Sub(r.Origin).Div(r.Direction)
	f := c.Max.Sub(r.Origin).Div(r.Direction)
	n, f = n.Min(f), n.Max(f)
	t0 := math.Max(math.Max(n.X, n.Y), n.Z)
	t1 := math.Min(math.Min(f.X, f.Y), f.Z)
	if t0 < 1e-3 && t1 > 1e-3 {
		return Hit{c, t1}
	}
	if t0 >= 1e-3 && t0 < t1 {
		return Hit{c, t0}
	}
	return NoHit
}

func (c *Cube) Paths() Paths {
	x1, y1, z1 := c.Min.X, c.Min.Y, c.Min.Z
	x2, y2, z2 := c.Max.X, c.Max.Y, c.Max.Z
	paths := Paths{
		{{x1, y1, z1}, {x1, y1, z2}},
		{{x1, y1, z1}, {x1, y2, z1}},
		{{x1, y1, z1}, {x2, y1, z1}},
		{{x1, y1, z2}, {x1, y2, z2}},
		{{x1, y1, z2}, {x2, y1, z2}},
		{{x1, y2, z1}, {x1, y2, z2}},
		{{x1, y2, z1}, {x2, y2, z1}},
		{{x1, y2, z2}, {x2, y2, z2}},
		{{x2, y1, z1}, {x2, y1, z2}},
		{{x2, y1, z1}, {x2, y2, z1}},
		{{x2, y1, z2}, {x2, y2, z2}},
		{{x2, y2, z1}, {x2, y2, z2}},
	}
	return paths
	paths = paths[:0]
	for i := 0; i <= 10; i++ {
		p := float64(i) / 10
		var x, y float64
		x = x1 + (x2-x1)*p
		y = y1 + (y2-y1)*p
		paths = append(paths, Path{{x, y1, z1}, {x, y1, z2}})
		paths = append(paths, Path{{x, y2, z1}, {x, y2, z2}})
		paths = append(paths, Path{{x1, y, z1}, {x1, y, z2}})
		paths = append(paths, Path{{x2, y, z1}, {x2, y, z2}})
	}
	return paths
}

package ln

import "math"

type Cone struct {
	Radius float64
	Height float64
}

func NewCone(radius, height float64) *Cone {
	return &Cone{radius, height}
}

func (c *Cone) Compile() {
}

func (c *Cone) BoundingBox() Box {
	r := c.Radius
	return Box{Vector{-r, -r, 0}, Vector{r, r, c.Height}}
}

func (c *Cone) Contains(v Vector, f float64) bool {
	return false
}

func (shape *Cone) Intersect(ray Ray) Hit {
	o := ray.Origin
	d := ray.Direction
	r := shape.Radius
	h := shape.Height

	k := r / h
	k = k * k

	a := d.X*d.X + d.Y*d.Y - k*d.Z*d.Z
	b := 2 * (d.X*o.X + d.Y*o.Y - k*d.Z*(o.Z-h))
	c := o.X*o.X + o.Y*o.Y - k*(o.Z-h)*(o.Z-h)
	q := b*b - 4*a*c
	if q <= 0 {
		return NoHit
	}
	s := math.Sqrt(q)
	t0 := (-b + s) / (2 * a)
	t1 := (-b - s) / (2 * a)
	if t0 > t1 {
		t0, t1 = t1, t0
	}
	if t0 > 1e-6 {
		p := ray.Position(t0)
		if p.Z > 0 && p.Z < h {
			return Hit{shape, t0}
		}
	}
	if t1 > 1e-6 {
		p := ray.Position(t1)
		if p.Z > 0 && p.Z < h {
			return Hit{shape, t1}
		}
	}
	return NoHit

}

func (c *Cone) Paths() Paths {
	var result Paths
	for a := 0; a < 360; a += 30 {
		x := c.Radius * math.Cos(Radians(float64(a)))
		y := c.Radius * math.Sin(Radians(float64(a)))
		result = append(result, Path{{x, y, 0}, {0, 0, c.Height}})
	}
	return result
}

type OutlineCone struct {
	Cone
	Eye Vector
	Up  Vector
}

func NewOutlineCone(eye, up Vector, radius, height float64) *OutlineCone {
	cone := NewCone(radius, height)
	return &OutlineCone{*cone, eye, up}
}

func (c *OutlineCone) Paths() Paths {
	center := Vector{0, 0, 0}
	hyp := center.Sub(c.Eye).Length()
	opp := c.Radius
	theta := math.Asin(opp / hyp)
	adj := opp / math.Tan(theta)
	d := math.Cos(theta) * adj
	// r := math.Sin(theta) * adj
	w := center.Sub(c.Eye).Normalize()
	u := w.Cross(c.Up).Normalize()
	c0 := c.Eye.Add(w.MulScalar(d))
	a0 := c0.Add(u.MulScalar(c.Radius * 1.01))
	b0 := c0.Add(u.MulScalar(-c.Radius * 1.01))

	var p0 Path
	for a := 0; a < 360; a++ {
		x := c.Radius * math.Cos(Radians(float64(a)))
		y := c.Radius * math.Sin(Radians(float64(a)))
		p0 = append(p0, Vector{x, y, 0})
	}
	return Paths{
		p0,
		{{a0.X, a0.Y, 0}, {0, 0, c.Height}},
		{{b0.X, b0.Y, 0}, {0, 0, c.Height}},
	}
}

func NewTransformedOutlineCone(eye, up, v0, v1 Vector, radius float64) Shape {
	d := v1.Sub(v0)
	z := d.Length()
	a := math.Acos(d.Normalize().Dot(up))
	m := Translate(v0)
	if a != 0 {
		u := d.Cross(up).Normalize()
		m = Rotate(u, a).Translate(v0)
	}
	c := NewOutlineCone(m.Inverse().MulPosition(eye), up, radius, z)
	return NewTransformedShape(c, m)
}

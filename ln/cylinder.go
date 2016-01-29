package ln

import "math"

type Cylinder struct {
	Radius float64
	Z0, Z1 float64
}

func NewCylinder(radius, z0, z1 float64) *Cylinder {
	return &Cylinder{radius, z0, z1}
}

func (c *Cylinder) Compile() {
}

func (c *Cylinder) BoundingBox() Box {
	r := c.Radius
	return Box{Vector{-r, -r, c.Z0}, Vector{r, r, c.Z1}}
}

func (c *Cylinder) Contains(v Vector, f float64) bool {
	xy := Vector{v.X, v.Y, 0}
	if xy.Length() > c.Radius+f {
		return false
	}
	return v.Z >= c.Z0-f && v.Z <= c.Z1+f
}

func (shape *Cylinder) Intersect(ray Ray) Hit {
	r := shape.Radius
	o := ray.Origin
	d := ray.Direction
	a := d.X*d.X + d.Y*d.Y
	b := 2*o.X*d.X + 2*o.Y*d.Y
	c := o.X*o.X + o.Y*o.Y - r*r
	q := b*b - 4*a*c
	if q < 0 {
		return NoHit
	}
	s := math.Sqrt(q)
	t0 := (-b + s) / (2 * a)
	t1 := (-b - s) / (2 * a)
	if t0 > t1 {
		t0, t1 = t1, t0
	}
	z0 := o.Z + t0*d.Z
	z1 := o.Z + t1*d.Z
	if t0 > 1e-6 && shape.Z0 < z0 && z0 < shape.Z1 {
		return Hit{shape, t0}
	}
	if t1 > 1e-6 && shape.Z0 < z1 && z1 < shape.Z1 {
		return Hit{shape, t1}
	}
	return NoHit

}

func (c *Cylinder) Paths() Paths {
	var result Paths
	for a := 0; a < 360; a += 10 {
		x := c.Radius * math.Cos(Radians(float64(a)))
		y := c.Radius * math.Sin(Radians(float64(a)))
		result = append(result, Path{{x, y, c.Z0}, {x, y, c.Z1}})
	}
	return result
}

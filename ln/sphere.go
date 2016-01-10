package ln

import "math"

type Sphere struct {
	Center Vector
	Radius float64
	Box    Box
}

func NewSphere(center Vector, radius float64) Shape {
	min := Vector{center.X - radius, center.Y - radius, center.Z - radius}
	max := Vector{center.X + radius, center.Y + radius, center.Z + radius}
	box := Box{min, max}
	return &Sphere{center, radius, box}
}

func (s *Sphere) Compile() {
}

func (s *Sphere) BoundingBox() Box {
	return s.Box
}

func (s *Sphere) Intersect(r Ray) Hit {
	to := r.Origin.Sub(s.Center)
	b := to.Dot(r.Direction)
	c := to.Dot(to) - s.Radius*s.Radius
	d := b*b - c
	if d > 0 {
		d = math.Sqrt(d)
		t1 := -b - d
		if t1 > 0 {
			return Hit{s, t1}
		}
		// t2 := -b + d
		// if t2 > 0 {
		// 	return Hit{s, t2}
		// }
	}
	return NoHit
}

func (s *Sphere) Paths() Paths {
	return nil
}

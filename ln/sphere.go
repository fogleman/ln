package ln

import (
	"math"
	"math/rand"
)

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
	radius := s.Radius - 0.001
	to := r.Origin.Sub(s.Center)
	b := to.Dot(r.Direction)
	c := to.Dot(to) - radius*radius
	d := b*b - c
	if d > 0 {
		d = math.Sqrt(d)
		t1 := -b - d
		if t1 > 0 {
			return Hit{s, t1}
		}
		t2 := -b + d
		if t2 > 0 {
			return Hit{s, t2}
		}
	}
	return NoHit
}

func (s *Sphere) Paths3() Paths {
	var paths Paths
	for i := 0; i < 20000; i++ {
		a := RandomUnitVector()
		b := a.Add(RandomUnitVector().MulScalar(0.001)).Normalize()
		a = a.MulScalar(s.Radius).Add(s.Center)
		b = b.MulScalar(s.Radius).Add(s.Center)
		paths = append(paths, Path{a, b})
	}
	return paths
}

func (s *Sphere) Paths2() Paths {
	var equator Path
	for lng := 0; lng <= 360; lng++ {
		v := LatLngToXYZ(0, float64(lng), s.Radius)
		equator = append(equator, v)
	}
	var paths Paths
	for i := 0; i < 100; i++ {
		m := Identity()
		for j := 0; j < 3; j++ {
			v := RandomUnitVector()
			m = m.Rotate(v, rand.Float64()*2*math.Pi)
		}
		m = m.Translate(s.Center)
		paths = append(paths, equator.Transform(m))
	}
	return paths
}

func (s *Sphere) Paths() Paths {
	var paths Paths
	n := 10
	for lat := -90 + n; lat <= 90-n; lat += n {
		var path Path
		for lng := 0; lng <= 360; lng++ {
			v := LatLngToXYZ(float64(lat), float64(lng), s.Radius).Add(s.Center)
			path = append(path, v)
		}
		paths = append(paths, path)
	}
	for lng := 0; lng <= 360; lng += n {
		var path Path
		for lat := -90 + n; lat <= 90-n; lat++ {
			v := LatLngToXYZ(float64(lat), float64(lng), s.Radius).Add(s.Center)
			path = append(path, v)
		}
		paths = append(paths, path)
	}
	return paths
}

func LatLngToXYZ(lat, lng, radius float64) Vector {
	lat, lng = Radians(lat), Radians(lng)
	x := radius * math.Cos(lat) * math.Cos(lng)
	y := radius * math.Cos(lat) * math.Sin(lng)
	z := radius * math.Sin(lat)
	return Vector{x, y, z}
}

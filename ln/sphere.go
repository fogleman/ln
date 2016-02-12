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

func NewSphere(center Vector, radius float64) *Sphere {
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

func (s *Sphere) Contains(v Vector, f float64) bool {
	return v.Sub(s.Center).Length() <= s.Radius+f
}

func (s *Sphere) Intersect(r Ray) Hit {
	radius := s.Radius
	to := r.Origin.Sub(s.Center)
	b := to.Dot(r.Direction)
	c := to.Dot(to) - radius*radius
	d := b*b - c
	if d > 0 {
		d = math.Sqrt(d)
		t1 := -b - d
		if t1 > 1e-2 {
			return Hit{s, t1}
		}
		t2 := -b + d
		if t2 > 1e-2 {
			return Hit{s, t2}
		}
	}
	return NoHit
}

func (s *Sphere) Paths4() Paths {
	var paths Paths
	var seen []Vector
	var radii []float64
	for i := 0; i < 140; i++ {
		var v Vector
		var m float64
		for {
			v = RandomUnitVector()
			m = rand.Float64()*0.25 + 0.05
			ok := true
			for i, other := range seen {
				threshold := m + radii[i] + 0.02
				if other.Sub(v).Length() < threshold {
					ok = false
					break
				}
			}
			if ok {
				seen = append(seen, v)
				radii = append(radii, m)
				break
			}
		}
		p := v.Cross(RandomUnitVector()).Normalize()
		q := p.Cross(v).Normalize()
		n := rand.Intn(4) + 1
		for k := 0; k < n; k++ {
			var path Path
			for j := 0; j <= 360; j += 5 {
				a := Radians(float64(j))
				x := v
				x = x.Add(p.MulScalar(math.Cos(a) * m))
				x = x.Add(q.MulScalar(math.Sin(a) * m))
				x = x.Normalize()
				x = x.MulScalar(s.Radius).Add(s.Center)
				path = append(path, x)
			}
			paths = append(paths, path)
			m *= 0.75
		}
	}
	return paths
}

func (s *Sphere) Paths3() Paths {
	var paths Paths
	for i := 0; i < 20000; i++ {
		v := RandomUnitVector()
		v = v.MulScalar(s.Radius).Add(s.Center)
		paths = append(paths, Path{v, v})
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
	o := 10
	for lat := -90 + o; lat <= 90-o; lat += n {
		var path Path
		for lng := 0; lng <= 360; lng++ {
			v := LatLngToXYZ(float64(lat), float64(lng), s.Radius).Add(s.Center)
			path = append(path, v)
		}
		paths = append(paths, path)
	}
	for lng := 0; lng <= 360; lng += n {
		var path Path
		for lat := -90 + o; lat <= 90-o; lat++ {
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

type OutlineSphere struct {
	Sphere
	Eye Vector
	Up  Vector
}

func NewOutlineSphere(eye, up, center Vector, radius float64) *OutlineSphere {
	sphere := NewSphere(center, radius)
	return &OutlineSphere{*sphere, eye, up}
}

func (s *OutlineSphere) Paths() Paths {
	var path Path
	center := s.Sphere.Center
	radius := s.Sphere.Radius

	hyp := center.Sub(s.Eye).Length()
	opp := radius
	theta := math.Asin(opp / hyp)
	adj := opp / math.Tan(theta)
	d := math.Cos(theta) * adj
	r := math.Sin(theta) * adj

	w := center.Sub(s.Eye).Normalize()
	u := w.Cross(s.Up).Normalize()
	v := w.Cross(u).Normalize()
	c := s.Eye.Add(w.MulScalar(d))
	for i := 0; i <= 360; i++ {
		a := Radians(float64(i))
		p := c
		p = p.Add(u.MulScalar(math.Cos(a) * r))
		p = p.Add(v.MulScalar(math.Sin(a) * r))
		path = append(path, p)
	}
	return Paths{path}
}

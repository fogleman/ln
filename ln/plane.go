package ln

type Plane struct {
	Point  Vector
	Normal Vector
}

func (p *Plane) IntersectSegment(v0, v1 Vector) (Vector, bool) {
	u := v1.Sub(v0)
	w := v0.Sub(p.Point)
	d := p.Normal.Dot(u)
	n := -p.Normal.Dot(w)
	if d > -EPS && d < EPS {
		return Vector{}, false
	}
	t := n / d
	if t < 0 || t > 1 {
		return Vector{}, false
	}
	v := v0.Add(u.MulScalar(t))
	return v, true
}

func (p *Plane) IntersectTriangle(t *Triangle) (Vector, Vector, bool) {
	v1, ok1 := p.IntersectSegment(t.V1, t.V2)
	v2, ok2 := p.IntersectSegment(t.V2, t.V3)
	v3, ok3 := p.IntersectSegment(t.V3, t.V1)
	if ok1 && ok2 {
		return v1, v2, true
	}
	if ok1 && ok3 {
		return v1, v3, true
	}
	if ok2 && ok3 {
		return v2, v3, true
	}
	return Vector{}, Vector{}, false
}

func (p *Plane) IntersectMesh(m *Mesh) Paths {
	var result Paths
	for _, t := range m.Triangles {
		if v1, v2, ok := p.IntersectTriangle(t); ok {
			result = append(result, Path{v1, v2})
		}
	}
	return result
}

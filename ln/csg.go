package ln

type Op int

const (
	Intersection Op = iota
	Difference
	// Union
)

type BooleanShape struct {
	Op   Op
	A, B Shape
}

func NewBooleanShape(op Op, shapes ...Shape) Shape {
	if len(shapes) == 0 {
		return &EmptyShape{}
	}
	shape := shapes[0]
	for i := 1; i < len(shapes); i++ {
		shape = &BooleanShape{op, shape, shapes[i]}
	}
	return shape
}

func NewIntersection(shapes ...Shape) Shape {
	return NewBooleanShape(Intersection, shapes...)
}

func NewDifference(shapes ...Shape) Shape {
	return NewBooleanShape(Difference, shapes...)
}

func (s *BooleanShape) Compile() {
}

func (s *BooleanShape) BoundingBox() Box {
	// TODO: fix this
	a := s.A.BoundingBox()
	b := s.B.BoundingBox()
	return a.Extend(b)
}

func (s *BooleanShape) Contains(v Vector, f float64) bool {
	f = 1e-3
	switch s.Op {
	case Intersection:
		return s.A.Contains(v, f) && s.B.Contains(v, f)
	case Difference:
		return s.A.Contains(v, f) && !s.B.Contains(v, -f)
	}
	return false
}

func (s *BooleanShape) Intersect(r Ray) Hit {
	h1 := s.A.Intersect(r)
	h2 := s.B.Intersect(r)
	h := h1.Min(h2)
	v := r.Position(h.T)
	if !h.Ok() || s.Contains(v, 0) {
		return h
	}
	return s.Intersect(Ray{r.Position(h.T + 0.01), r.Direction})
}

func (s *BooleanShape) Paths() Paths {
	p := s.A.Paths()
	p = append(p, s.B.Paths()...)
	return p.Chop(0.01).Filter(s)
}

func (s *BooleanShape) Filter(v Vector) (Vector, bool) {
	return v, s.Contains(v, 0)
}

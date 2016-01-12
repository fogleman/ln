package ln

type Shape interface {
	Compile()
	BoundingBox() Box
	Contains(Vector) bool
	Intersect(Ray) Hit
	Paths() Paths
}

type EmptyShape struct {
}

func (s *EmptyShape) Compile() {
}

func (s *EmptyShape) BoundingBox() Box {
	return Box{Vector{}, Vector{}}
}

func (s *EmptyShape) Contains(v Vector) bool {
	return false
}

func (s *EmptyShape) Intersect(r Ray) Hit {
	return NoHit
}

func (s *EmptyShape) Paths() Paths {
	return nil
}

type TransformedShape struct {
	Shape
	Matrix  Matrix
	Inverse Matrix
}

func NewTransformedShape(s Shape, m Matrix) Shape {
	return &TransformedShape{s, m, m.Inverse()}
}

func (s *TransformedShape) BoundingBox() Box {
	return s.Matrix.MulBox(s.Shape.BoundingBox())
}

func (s *TransformedShape) Intersect(r Ray) Hit {
	return s.Shape.Intersect(s.Inverse.MulRay(r))
}

func (s *TransformedShape) Paths() Paths {
	return s.Shape.Paths().Transform(s.Matrix)
}

type Intersection struct {
	A, B Shape
}

func NewIntersection(shapes ...Shape) Shape {
	if len(shapes) == 0 {
		return &EmptyShape{}
	}
	shape := shapes[0]
	for i := 1; i < len(shapes); i++ {
		shape = &Intersection{shape, shapes[i]}
	}
	return shape
}

func (s *Intersection) Compile() {
}

func (s *Intersection) BoundingBox() Box {
	// TODO: fix this
	a := s.A.BoundingBox()
	b := s.B.BoundingBox()
	return a.Extend(b)
}

func (s *Intersection) Contains(v Vector) bool {
	return s.A.Contains(v) && s.B.Contains(v)
}

func (s *Intersection) Intersect(r Ray) Hit {
	h1 := s.A.Intersect(r)
	h2 := s.B.Intersect(r)
	h := h1.Min(h2)
	v := r.Position(h.T)
	if !h.Ok() || s.Contains(v) {
		return h
	}
	return s.Intersect(Ray{r.Position(h.T + 0.001), r.Direction})
}

func (s *Intersection) Paths() Paths {
	p := s.A.Paths()
	p = append(p, s.B.Paths()...)
	return p.Filter(s)
}

func (s *Intersection) Filter(v Vector) (Vector, bool) {
	return v, s.Contains(v)
}

type Difference struct {
	A, B Shape
}

func NewDifference(shapes ...Shape) Shape {
	if len(shapes) == 0 {
		return &EmptyShape{}
	}
	shape := shapes[0]
	for i := 1; i < len(shapes); i++ {
		shape = &Difference{shape, shapes[i]}
	}
	return shape
}

func (s *Difference) Compile() {
}

func (s *Difference) BoundingBox() Box {
	// TODO: fix this
	a := s.A.BoundingBox()
	b := s.B.BoundingBox()
	return a.Extend(b)
}

func (s *Difference) Contains(v Vector) bool {
	return s.A.Contains(v) && !s.B.Contains(v)
}

func (s *Difference) Intersect(r Ray) Hit {
	h1 := s.A.Intersect(r)
	h2 := s.B.Intersect(r)
	h := h1.Min(h2)
	v := r.Position(h.T)
	if !h.Ok() || s.Contains(v) {
		return h
	}
	return s.Intersect(Ray{v, r.Direction})
}

func (s *Difference) Paths() Paths {
	p := s.A.Paths()
	p = append(p, s.B.Paths()...)
	return p.Filter(s)
}

func (s *Difference) Filter(v Vector) (Vector, bool) {
	return v, s.Contains(v)
}

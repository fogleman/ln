package ln

type Shape interface {
	Compile()
	BoundingBox() Box
	Intersect(Ray) Hit
	Paths() Paths
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
	hit := s.Shape.Intersect(s.Inverse.MulRay(r))
	if !hit.Ok() {
		return hit
	}
	// if s.Shape is a Mesh, the hit.Shape will be a Triangle in the Mesh
	// we need to transform this Triangle, not the Mesh itself
	shape := &TransformedShape{hit.Shape, s.Matrix, s.Inverse}
	return Hit{shape, hit.T}
}

package ln

type Scene struct {
	Shapes []Shape
	Tree   *Tree
}

func (s *Scene) Compile() {
	for _, shape := range s.Shapes {
		shape.Compile()
	}
	if s.Tree == nil {
		s.Tree = NewTree(s.Shapes)
	}
}

func (s *Scene) Add(shape Shape) {
	s.Shapes = append(s.Shapes, shape)
}

func (s *Scene) Intersect(r Ray) Hit {
	return s.Tree.Intersect(r)
}

func (s *Scene) Visible(eye, point Vector) bool {
	v := eye.Sub(point)
	r := Ray{point, v.Normalize()}
	hit := s.Intersect(r)
	return hit.T >= v.Length()
}

func (s *Scene) Paths() Paths {
	var result Paths
	for _, shape := range s.Shapes {
		result = append(result, shape.Paths()...)
	}
	return result
}

func (s *Scene) Render(eye, center, up Vector, width, height, fovy, near, far, step float64) Paths {
	aspect := width / height
	matrix := LookAt(eye, center, up)
	matrix = matrix.Perspective(fovy, aspect, near, far)
	return s.RenderWithMatrix(matrix, eye, width, height, step)
}

func (s *Scene) RenderWithMatrix(matrix Matrix, eye Vector, width, height, step float64) Paths {
	s.Compile()
	paths := s.Paths()
	if step > 0 {
		paths = paths.Chop(step)
	}
	paths = paths.Filter(&ClipFilter{matrix, eye, s})
	if step > 0 {
		paths = paths.Simplify(1e-6)
	}
	matrix = Translate(Vector{1, 1, 0}).Scale(Vector{width / 2, height / 2, 0})
	paths = paths.Transform(matrix)
	return paths
}

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
		// box := shape.BoundingBox()
		// result = append(result, NewCube(box.Min, box.Max).Paths()...)
	}
	return result
}

func (s *Scene) Render(eye, center, up Vector, fovy, aspect, near, far, step float64) Paths {
	s.Compile()
	paths := s.Paths()
	matrix := LookAt(eye, center, up)
	matrix = matrix.Perspective(fovy, aspect, near, far)
	if step > 0 {
		paths = paths.Chop(step)
	}
	paths = paths.Clip(matrix, eye, s)
	paths = append(paths, Path{{-1, -1, 0}, {1, -1, 0}, {1, 1, 0}, {-1, 1, 0}, {-1, -1, 0}})
	return paths
}

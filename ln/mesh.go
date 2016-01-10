package ln

type Mesh struct {
	Box       Box
	Triangles []*Triangle
	Tree      *Tree
}

func NewMesh(triangles []*Triangle) *Mesh {
	box := BoxForTriangles(triangles)
	return &Mesh{box, triangles, nil}
}

func (m *Mesh) Compile() {
	if m.Tree == nil {
		shapes := make([]Shape, len(m.Triangles))
		for i, triangle := range m.Triangles {
			shapes[i] = triangle
		}
		m.Tree = NewTree(shapes)
	}
}

func (m *Mesh) BoundingBox() Box {
	return m.Box
}

func (m *Mesh) Intersect(r Ray) Hit {
	return m.Tree.Intersect(r)
}

func (m *Mesh) Paths() Paths {
	return nil
}

func (m *Mesh) UpdateBoundingBox() {
	m.Box = BoxForTriangles(m.Triangles)
}

func (m *Mesh) UnitCube() {
	m.FitInside(Box{Vector{}, Vector{1, 1, 1}}, Vector{})
	m.MoveTo(Vector{}, Vector{0.5, 0.5, 0.5})
}

func (m *Mesh) MoveTo(position, anchor Vector) {
	matrix := Translate(position.Sub(m.Box.Anchor(anchor)))
	m.Transform(matrix)
}

func (m *Mesh) FitInside(box Box, anchor Vector) {
	scale := box.Size().Div(m.BoundingBox().Size()).MinComponent()
	extra := box.Size().Sub(m.BoundingBox().Size().MulScalar(scale))
	matrix := Identity()
	matrix = matrix.Translate(m.BoundingBox().Min.MulScalar(-1))
	matrix = matrix.Scale(Vector{scale, scale, scale})
	matrix = matrix.Translate(box.Min.Add(extra.Mul(anchor)))
	m.Transform(matrix)
}

func (m *Mesh) Transform(matrix Matrix) {
	for _, t := range m.Triangles {
		t.V1 = matrix.MulPosition(t.V1)
		t.V2 = matrix.MulPosition(t.V2)
		t.V3 = matrix.MulPosition(t.V3)
		t.UpdateBoundingBox()
	}
	m.UpdateBoundingBox()
	m.Tree = nil // dirty
}

func (m *Mesh) SaveBinarySTL(path string) error {
	return SaveBinarySTL(path, m)
}

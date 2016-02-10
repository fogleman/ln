package ln

import "math"

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

func (m *Mesh) Contains(v Vector, f float64) bool {
	return false
}

func (m *Mesh) Intersect(r Ray) Hit {
	return m.Tree.Intersect(r)
}

func (m *Mesh) Paths() Paths {
	var result Paths
	for _, t := range m.Triangles {
		result = append(result, t.Paths()...)
	}
	return result
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

func (m *Mesh) Voxelize(size float64) []*Cube {
	z1 := m.Box.Min.Z
	z2 := m.Box.Max.Z
	set := make(map[Vector]bool)
	for z := z1; z <= z2; z += size {
		plane := Plane{Vector{0, 0, z}, Vector{0, 0, 1}}
		paths := plane.IntersectMesh(m)
		for _, path := range paths {
			for _, v := range path {
				x := math.Floor(v.X/size+0.5) * size
				y := math.Floor(v.Y/size+0.5) * size
				z := math.Floor(v.Z/size+0.5) * size
				set[Vector{x, y, z}] = true
			}
		}
	}
	var result []*Cube
	for v, _ := range set {
		cube := NewCube(v.SubScalar(size/2), v.AddScalar(size/2))
		result = append(result, cube)
	}
	return result
}

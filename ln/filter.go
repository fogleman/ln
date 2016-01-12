package ln

type Filter interface {
	Filter(v Vector) (Vector, bool)
}

type ClipFilter struct {
	Matrix Matrix
	Eye    Vector
	Scene  *Scene
}

var ClipBox = Box{Vector{-1, -1, -1}, Vector{1, 1, 1}}

func (f *ClipFilter) Filter(v Vector) (Vector, bool) {
	if !f.Scene.Visible(f.Eye, v) {
		return v, false
	}
	v = f.Matrix.MulPositionW(v)
	if !ClipBox.Contains(v) {
		return v, false
	}
	return v, true
}

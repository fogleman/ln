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
	w := f.Matrix.MulPositionW(v)
	if !f.Scene.Visible(f.Eye, v) {
		return w, false
	}
	if !ClipBox.Contains(w) {
		return w, false
	}
	return w, true
}

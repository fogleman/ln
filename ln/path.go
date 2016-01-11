package ln

import "github.com/ungerik/go-cairo"

type Path []Vector

func (p Path) BoundingBox() Box {
	box := Box{p[0], p[0]}
	for _, v := range p {
		box = box.Extend(Box{v, v})
	}
	return box
}

func (p Path) Transform(matrix Matrix) Path {
	var result Path
	for _, v := range p {
		result = append(result, matrix.MulPosition(v))
	}
	return result
}

func (p Path) TransformW(matrix Matrix) Path {
	var result Path
	for _, v := range p {
		v, w := matrix.MulPositionW(v)
		v = v.DivScalar(w)
		result = append(result, v)
	}
	return result
}

func (p Path) Chop(step float64) Path {
	var result Path
	for i := 0; i < len(p)-1; i++ {
		a := p[i]
		b := p[i+1]
		v := b.Sub(a)
		l := v.Length()
		if i == 0 {
			result = append(result, a)
		}
		d := step
		for d < l {
			result = append(result, a.Add(v.MulScalar(d/l)))
			d += step
		}
		result = append(result, b)
	}
	return result
}

func (p Path) Clip(eye Vector, scene *Scene) Paths {
	var result Paths
	var path Path
	for _, v := range p {
		visible := scene.Visible(eye, v)
		if visible {
			path = append(path, v)
		} else {
			if len(path) > 0 {
				result = append(result, path)
				path = nil
			}
		}
	}
	if len(path) > 0 {
		result = append(result, path)
	}
	return result
}

type Paths []Path

func (p Paths) BoundingBox() Box {
	box := p[0].BoundingBox()
	for _, path := range p {
		box = box.Extend(path.BoundingBox())
	}
	return box
}

func (p Paths) Transform(matrix Matrix) Paths {
	var result Paths
	for _, path := range p {
		result = append(result, path.Transform(matrix))
	}
	return result
}

func (p Paths) TransformW(matrix Matrix) Paths {
	var result Paths
	for _, path := range p {
		result = append(result, path.TransformW(matrix))
	}
	return result
}

func (p Paths) Chop(step float64) Paths {
	var result Paths
	for _, path := range p {
		result = append(result, path.Chop(step))
	}
	return result
}

func (p Paths) Clip(eye Vector, scene *Scene) Paths {
	var result Paths
	for _, path := range p {
		result = append(result, path.Clip(eye, scene)...)
	}
	return result
}

func (p Paths) Render(path string, scale float64) {
	pad := 10.0
	box := p.BoundingBox()
	dx := box.Max.X - box.Min.X
	dy := box.Max.Y - box.Min.Y
	width := int(dx*scale + pad*2)
	height := int(dy*scale + pad*2)
	dc := cairo.NewSurface(cairo.FORMAT_ARGB32, width, height)
	dc.Scale(1, -1)
	dc.Translate(0, float64(-height))
	dc.SetSourceRGB(1, 1, 1)
	dc.Paint()
	dc.SetSourceRGB(0, 0, 0)
	for _, path := range p {
		dc.NewSubPath()
		for _, v := range path {
			x := pad + (v.X-box.Min.X)*scale
			y := pad + (v.Y-box.Min.Y)*scale
			dc.LineTo(x, y)
		}
	}
	dc.Stroke()
	dc.WriteToPNG(path)
}

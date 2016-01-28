package ln

import (
	"fmt"

	"github.com/ungerik/go-cairo"
)

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

func (p Path) Filter(f Filter) Paths {
	var result Paths
	var path Path
	for _, v := range p {
		v, ok := f.Filter(v)
		if ok {
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

func (p Path) Print() {
	for _, v := range p {
		fmt.Printf("%g,%g;", v.X, v.Y)
	}
	fmt.Println()
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

func (p Paths) Chop(step float64) Paths {
	var result Paths
	for _, path := range p {
		result = append(result, path.Chop(step))
	}
	return result
}

func (p Paths) Filter(f Filter) Paths {
	var result Paths
	for _, path := range p {
		result = append(result, path.Filter(f)...)
	}
	return result
}

func (p Paths) Print() {
	for _, path := range p {
		path.Print()
	}
}

func (p Paths) Render(path string, width, height, scale float64) {
	dc := cairo.NewSurface(cairo.FORMAT_ARGB32, int(width*scale), int(height*scale))
	dc.SetLineCap(cairo.LINE_CAP_ROUND)
	dc.SetLineJoin(cairo.LINE_JOIN_ROUND)
	dc.SetLineWidth(3)
	dc.Scale(1, -1)
	dc.Translate(0, -height*scale)
	dc.SetSourceRGB(1, 1, 1)
	dc.Paint()
	dc.SetSourceRGB(0, 0, 0)
	for _, path := range p {
		dc.NewSubPath()
		for _, v := range path {
			dc.LineTo(v.X*scale, v.Y*scale)
		}
	}
	dc.Stroke()
	dc.WriteToPNG(path)
}

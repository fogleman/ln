package ln

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/fogleman/gg"
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
		// ok = ok || i%8 < 4 // show hidden lines
		if ok {
			path = append(path, v)
		} else {
			if len(path) > 1 {
				result = append(result, path)
			}
			path = nil
		}
	}
	if len(path) > 1 {
		result = append(result, path)
	}
	return result
}

func (p Path) Simplify(threshold float64) Path {
	if len(p) < 3 {
		return p
	}
	a := p[0]
	b := p[len(p)-1]
	index := -1
	distance := 0.0
	for i := 1; i < len(p)-1; i++ {
		d := p[i].SegmentDistance(a, b)
		if d > distance {
			index = i
			distance = d
		}
	}
	if distance > threshold {
		r1 := p[:index+1].Simplify(threshold)
		r2 := p[index:].Simplify(threshold)
		return append(r1[:len(r1)-1], r2...)
	} else {
		return Path{a, b}
	}
}

func (p Path) Print() {
	for _, v := range p {
		fmt.Printf("%g,%g;", v.X, v.Y)
	}
	fmt.Println()
}

func (p Path) String() string {
	var parts []string
	for _, v := range p {
		parts = append(parts, fmt.Sprintf("%g,%g", v.X, v.Y))
	}
	return strings.Join(parts, ";")
}

func (p Path) ToSVG() string {
	var coords []string
	for _, v := range p {
		coords = append(coords, fmt.Sprintf("%f,%f", v.X, v.Y))
	}
	points := strings.Join(coords, " ")
	return fmt.Sprintf("<polyline stroke=\"black\" fill=\"none\" points=\"%s\" />", points)
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

func (p Paths) Simplify(threshold float64) Paths {
	var result Paths
	for _, path := range p {
		result = append(result, path.Simplify(threshold))
	}
	return result
}

func (p Paths) Print() {
	for _, path := range p {
		path.Print()
	}
}

func (p Paths) String() string {
	var parts []string
	for _, path := range p {
		parts = append(parts, path.String())
	}
	return strings.Join(parts, "\n")
}

func (p Paths) WriteToPNG(path string, width, height float64) {
	scale := 1.0
	w, h := int(width*scale), int(height*scale)
	dc := gg.NewContext(w, h)
	dc.InvertY()
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(3)
	for _, path := range p {
		for _, v := range path {
			dc.LineTo(v.X*scale, v.Y*scale)
		}
		dc.NewSubPath()
	}
	dc.Stroke()
	dc.SavePNG(path)
}

func (p Paths) ToSVG(width, height float64) string {
	var lines []string
	lines = append(lines, fmt.Sprintf("<svg width=\"%f\" height=\"%f\" version=\"1.1\" baseProfile=\"full\" xmlns=\"http://www.w3.org/2000/svg\">", width, height))
	lines = append(lines, fmt.Sprintf("<g transform=\"translate(0,%f) scale(1,-1)\">", height))
	for _, path := range p {
		lines = append(lines, path.ToSVG())
	}
	lines = append(lines, "</g></svg>")
	return strings.Join(lines, "\n")
}

func (p Paths) WriteToSVG(path string, width, height float64) error {
	return ioutil.WriteFile(path, []byte(p.ToSVG(width, height)), 0644)
}

func (p Paths) WriteToTXT(path string) error {
	return ioutil.WriteFile(path, []byte(p.String()), 0644)
}

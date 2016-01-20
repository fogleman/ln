package ln

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseIndex(value string, length int) int {
	parsed, _ := strconv.ParseInt(value, 0, 0)
	n := int(parsed)
	if n < 0 {
		n += length
	}
	return n
}

func LoadOBJ(path string) (*Mesh, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	vs := make([]Vector, 1, 1024) // 1-based indexing
	var triangles []*Triangle
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) == 0 {
			continue
		}
		keyword := fields[0]
		args := fields[1:]
		switch keyword {
		case "v":
			f := ParseFloats(args)
			v := Vector{f[0], f[1], f[2]}
			vs = append(vs, v)
		case "f":
			fvs := make([]int, len(args))
			for i, arg := range args {
				vertex := strings.Split(arg+"//", "/")
				fvs[i] = parseIndex(vertex[0], len(vs))
			}
			for i := 1; i < len(fvs)-1; i++ {
				i1, i2, i3 := 0, i, i+1
				t := NewTriangle(vs[fvs[i1]], vs[fvs[i2]], vs[fvs[i3]])
				triangles = append(triangles, t)
			}
		}
	}
	return NewMesh(triangles), scanner.Err()
}

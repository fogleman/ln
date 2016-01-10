package ln

type Ray struct {
	Origin, Direction Vector
}

func (r Ray) Position(t float64) Vector {
	return r.Origin.Add(r.Direction.MulScalar(t))
}

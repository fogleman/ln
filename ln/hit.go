package ln

type Hit struct {
	Shape Shape
	T     float64
}

var NoHit = Hit{nil, INF}

func (hit Hit) Ok() bool {
	return hit.T < INF
}

func (a Hit) Min(b Hit) Hit {
	if a.T <= b.T {
		return a
	}
	return b
}

func (a Hit) Max(b Hit) Hit {
	if a.T > b.T {
		return a
	}
	return b
}

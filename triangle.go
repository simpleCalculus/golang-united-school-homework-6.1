package golang_united_school_homework

import "math"

// Triangle must satisfy to Shape interface
type Triangle struct {
	Side float64
}

func (t Triangle) CalcPerimeter() float64 {
	return 3 * t.Side
}

func (t Triangle) CalcArea() float64 {
	angle := float64(60)
	s := math.Sin(angle * math.Pi / 180)
	return t.Side * t.Side * s / 2
}

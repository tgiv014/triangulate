package geometry

import "math"

type Point struct {
	X float64
	Y float64
}

func (p *Point) extend(op Point, percent float64) Point {
	dx := op.X - p.X
	dy := op.Y - p.Y
	return Point{p.X + dx*percent, p.Y + dy*percent}
}

func (p *Point) SqDist(op Point) float64 {
	dx := op.X - p.X
	dy := op.Y - p.Y
	return math.Pow(dx, 2) + math.Pow(dy, 2)
}

func (p *Point) Dist(op Point) float64 {
	return math.Sqrt(p.SqDist(op))
}

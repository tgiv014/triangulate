package geometry

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/fogleman/gg"
)

//    c
// A-----B
// |    /
// |   /
//b|  /a
// | /
// |/
// C

type Mode int

const (
	ModeDecide      Mode = 0
	ModeBalance     Mode = 1
	ModeTargetDepth Mode = 2
	ModeSpiral      Mode = 3
	ModeLine        Mode = 4
)

var allModes = []Mode{ModeBalance, ModeTargetDepth, ModeSpiral}

type Triangle struct {
	A, B, C             Point
	Aused, Bused, Cused bool
	Depth               int
}

func NewTriangle(A, B, C Point) Triangle {
	return Triangle{
		A, B, C,
		false, false, false,
		0,
	}
}

func (t *Triangle) Area() float64 {
	// perimeter
	c := t.A.Dist(t.B)
	a := t.B.Dist(t.C)
	b := t.C.Dist(t.A)
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

func (t *Triangle) Subdivide(pt int, percent float64) []Triangle {
	switch pt {
	case 0:
		// Point 0
		midpoint := t.B.extend(t.C, percent)
		return []Triangle{
			{t.A, t.B, midpoint, t.Aused, true, t.Cused, t.Depth + 1},
			{t.A, midpoint, t.C, t.Aused, t.Bused, true, t.Depth + 1},
		}
	case 1:
		midpoint := t.C.extend(t.A, percent)
		return []Triangle{
			{t.B, t.C, midpoint, t.Aused, true, t.Cused, t.Depth + 1},
			{t.B, midpoint, t.A, t.Aused, t.Bused, true, t.Depth + 1},
		}
	case 2:
		midpoint := t.A.extend(t.B, percent)
		return []Triangle{
			{t.C, t.A, midpoint, t.Aused, true, t.Cused, t.Depth + 1},
			{t.C, midpoint, t.B, t.Aused, t.Bused, t.Cused, t.Depth + 1},
		}
	default:
		panic("Invalid triangle point index")
	}
}

func collidesWithAny(pta, ptb Point, triangles []Triangle) bool {
	for _, t := range triangles {
		if t.PtInside(ptb) || t.LineCollidesWithTriangle(pta, ptb) {
			return true
		}
	}
	return false
}

func (t *Triangle) GrowRandomlyWithoutColliding(triangles []Triangle) (Triangle, bool) {
	attempts := 0
	if t.Aused && t.Bused && t.Cused {
		fmt.Println("Skip")
	}
	for (!t.Aused || !t.Bused || !t.Cused) && attempts < 100 {
		theta := rand.Float64()*math.Pi*0.3 + 0.2*math.Pi
		d := rand.Float64()*200. + 100.
		side := rand.Intn(3)
		// side := 0
		switch side {
		case 0:
			if !t.Aused {
				newTri := t.TangentTriangle(0, theta, d)
				if !collidesWithAny(newTri.A, newTri.C, triangles) &&
					!collidesWithAny(newTri.B, newTri.C, triangles) {
					t.Aused = true
					return newTri, true
				}
			}
			break
		case 1:
			if !t.Bused {
				newTri := t.TangentTriangle(1, theta, d)
				if !collidesWithAny(newTri.A, newTri.C, triangles) &&
					!collidesWithAny(newTri.B, newTri.C, triangles) {
					t.Bused = true
					return newTri, true
				}
			}
			break
		case 2:
			if !t.Cused {
				newTri := t.TangentTriangle(2, theta, d)
				if !collidesWithAny(newTri.A, newTri.C, triangles) &&
					!collidesWithAny(newTri.B, newTri.C, triangles) {
					t.Cused = true
					return newTri, true
				}
			}
			break
		}
		attempts++
	}
	return Triangle{}, false
}

func (t *Triangle) TangentTriangle(side int, ang, s float64) Triangle {
	switch side {
	case 0:
		// Side A pt C-B
		dx := t.C.X - t.B.X
		dy := t.C.Y - t.B.Y
		theta := math.Atan2(dy, dx)
		x := t.B.X + math.Cos(theta-ang)*s
		y := t.B.Y + math.Sin(theta-ang)*s
		return Triangle{
			t.C, t.B, Point{x, y},
			false, false, true,
			t.Depth,
		}
	case 1:
		// Side B pt A-C
		dx := t.A.X - t.C.X
		dy := t.A.Y - t.C.Y
		theta := math.Atan2(dy, dx)
		x := t.C.X + math.Cos(theta-ang)*s
		y := t.C.Y + math.Sin(theta-ang)*s
		return Triangle{
			t.A, t.C, Point{x, y},
			false, false, true,
			t.Depth,
		}
	case 2:
		// Side C pt B-A
		dx := t.B.X - t.A.X
		dy := t.B.Y - t.A.Y
		theta := math.Atan2(dy, dx)
		x := t.A.X + math.Cos(theta-ang)*s
		y := t.A.Y + math.Sin(theta-ang)*s
		return Triangle{
			t.B, t.A, Point{x, y},
			false, false, true,
			t.Depth,
		}
	default:
		panic("Invalid triangle point index")
	}
}

func (t *Triangle) Draw(dc *gg.Context) {
	dc.MoveTo(t.A.X, t.A.Y)
	dc.LineTo(t.B.X, t.B.Y)
	dc.LineTo(t.C.X, t.C.Y)
	dc.LineTo(t.A.X, t.A.Y)
}

// func (t *Triangle) Angle(p int) float64 {
// 	c := t.A.Dist(t.B)
// 	a := t.B.Dist(t.C)
// 	b := t.C.Dist(t.A)
// 	switch p {
// 	case 0: // a
// 		n := math.Pow(b, 2) + math.Pow(c, 2) - math.Pow(a, 2)
// 		return math.Acos(n / (2 * b * c))
// 	case 1: // b
// 		n := math.Pow(a, 2) + math.Pow(c, 2) - math.Pow(b, 2)
// 		return math.Acos(n / (2 * a * c))
// 	case 2: // c
// 		n := math.Pow(a, 2) + math.Pow(b, 2) - math.Pow(c, 2)
// 		return math.Acos(n / (2 * a * b))
// 	default:
// 		panic("Invalid pt index")
// 	}
// }

// func (t *Triangle) IdxOfLargestAngle() int {
// 	var l float64 = 0.0
// 	var idx = 0
// 	for i := 0; i < 3; i++ {
// 		ang := t.Angle(i)
// 		if ang > l {
// 			l = ang
// 			idx = i
// 		}
// 	}
// 	return idx
// }

func (t *Triangle) PtOppositeLongestSide() int {
	c := t.A.Dist(t.B)
	a := t.B.Dist(t.C)
	b := t.C.Dist(t.A)
	if a >= b && a >= c {
		return 0
	}
	if b >= a && b >= c {
		return 1
	}
	if c >= a && c >= b {
		return 2
	}
	panic("Inconceivable!")
}

func (t *Triangle) PtInside(pt Point) bool {

	d1 := Sign(pt, t.A, t.B)
	d2 := Sign(pt, t.B, t.C)
	d3 := Sign(pt, t.C, t.A)

	has_neg := (d1 < 0) || (d2 < 0) || (d3 < 0)
	has_pos := (d1 > 0) || (d2 > 0) || (d3 > 0)

	return !(has_neg && has_pos)
}

func (t *Triangle) HighlightUsed(dc *gg.Context) {
	dc.SetRGBA(1, 0, 0, 1)
	dc.SetLineWidth(4)
	if t.Aused {
		// Side A pt C-B
		dc.MoveTo(t.B.X, t.B.Y)
		dc.LineTo(t.C.X, t.C.Y)
		dc.Stroke()
	}

	if t.Bused {
		// Side B pt A-C
		dc.MoveTo(t.C.X, t.C.Y)
		dc.LineTo(t.A.X, t.A.Y)
		dc.Stroke()
	}

	if t.Cused {
		// Side C pt B-A
		dc.MoveTo(t.A.X, t.A.Y)
		dc.LineTo(t.B.X, t.B.Y)
		dc.Stroke()
	}
}

func (t *Triangle) GetSidePoints(side int) (pta, ptb Point) {
	switch side {
	case 0:
		// Side A pt C-B
		return t.C, t.B
	case 1:
		// Side B pt A-C
		return t.A, t.C
	case 2:
		// Side C pt B-A
		return t.B, t.A
	default:
		panic("Unknown side number")
	}
}

func (t *Triangle) LineCollidesWithTriangle(pta, ptb Point) bool {
	for i := 0; i < 3; i++ {
		pta1, ptb1 := t.GetSidePoints(i)
		if LineCollision(pta, ptb, pta1, ptb1) {
			return true
		}
	}
	return false
}

const MIN = 0.001
const MAX = 0.999

func LineCollision(pta0, ptb0, pta1, ptb1 Point) bool {
	uA := ((ptb1.X-pta1.X)*(pta0.Y-pta1.Y) - (ptb1.Y-pta1.Y)*(pta0.X-pta1.X)) / ((ptb1.Y-pta1.Y)*(ptb0.X-pta0.X) - (ptb1.X-pta1.X)*(ptb0.Y-pta0.Y))
	uB := ((ptb0.X-pta0.X)*(pta0.Y-pta1.Y) - (ptb0.Y-pta0.Y)*(pta0.X-pta1.X)) / ((ptb1.Y-pta1.Y)*(ptb0.X-pta0.X) - (ptb1.X-pta1.X)*(ptb0.Y-pta0.Y))
	if uA > 0. && uA < 1. && uB > 0. && uB < 1. {
		return true
	}
	return false
}

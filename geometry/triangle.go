package geometry

import (
	"math"

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
	A     Point
	B     Point
	C     Point
	Depth int
}

func NewTriangle(A, B, C Point) Triangle {
	return Triangle{
		A, B, C,
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
			{t.A, t.B, midpoint, t.Depth + 1},
			{t.A, midpoint, t.C, t.Depth + 1},
		}
	case 1:
		midpoint := t.C.extend(t.A, percent)
		return []Triangle{
			{t.B, t.C, midpoint, t.Depth + 1},
			{t.B, midpoint, t.A, t.Depth + 1},
		}
	case 2:
		midpoint := t.A.extend(t.B, percent)
		return []Triangle{
			{t.C, t.A, midpoint, t.Depth + 1},
			{t.C, midpoint, t.B, t.Depth + 1},
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

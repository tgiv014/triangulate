package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/tgiv014/triangulate/behavior"
	"github.com/tgiv014/triangulate/clevertriangle"
	"github.com/tgiv014/triangulate/geometry"
)

const (
	DPI = 300.
	W   = (8. * DPI)
	H   = (10. * DPI)
	S   = (2. * DPI)
	D   = (400.)
	GAP = (40.)
	D2  = (D + GAP)
)

func UniformTriangleCenteredOn(x, y float64, flip bool) clevertriangle.CleverTriangle {
	theta := make([]float64, 0)
	for i := 0; i < 3; i++ {
		newTheta := float64(i)*math.Pi*(2./3.) - math.Pi/2.
		if flip {
			newTheta += math.Pi
		}
		theta = append(theta, newTheta)
	}

	return clevertriangle.NewCleverTriangle(
		geometry.Point{x + D*math.Cos(theta[0]), y + D*math.Sin(theta[0])},
		geometry.Point{x + D*math.Cos(theta[1]), y + D*math.Sin(theta[1])},
		geometry.Point{x + D*math.Cos(theta[2]), y + D*math.Sin(theta[2])},
		behavior.RandomBehavior(),
	)
}

func main() {
	fmt.Println("Let's party")
	seed := time.Now().UnixNano()
	fmt.Printf("Using seed [%d]\n", seed)
	rand.Seed(seed)

	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	triangleSideLength := math.Cos(math.Pi/6) * 2. * D2
	triangleVerticalSep := math.Sqrt((3. / 4.) * math.Pow(triangleSideLength, 2))
	TriangleHorizSep := triangleSideLength / 2.

	SX := TriangleHorizSep
	SY := triangleVerticalSep

	for y := -1; y <= 2; y++ {
		for x := -2; x <= 2; x++ {
			xp := SX*float64(x) + 0.5*W
			yp := SY*float64(y) + 0.5*H - SY/2
			var triangles []clevertriangle.CleverTriangle
			if (x+y%2)%2 == 0 {
				triangles = []clevertriangle.CleverTriangle{
					UniformTriangleCenteredOn(xp, yp-(SY-D2)/2, true),
				}
			} else {
				triangles = []clevertriangle.CleverTriangle{
					UniformTriangleCenteredOn(xp, yp+(SY-D2)/2, false),
				}
			}

			for {
				newtriangles := make([]clevertriangle.CleverTriangle, 0)
				for _, t := range triangles {
					newtriangles = append(newtriangles, t.Cycle()...)
				}

				if len(newtriangles) == len(triangles) {
					break
				}
				triangles = newtriangles
			}

			maxdepth := 0
			for _, t := range triangles {
				if t.Depth > maxdepth {
					maxdepth = t.Depth
				}
			}

			for _, t := range triangles {
				dc.SetRGBA(1, 1, 1, 1)
				dc.SetLineWidth(2)
				t.Draw(dc)
				dc.Stroke()
				t.Draw(dc)
				// 46, 45, 77
				dc.SetRGBA(46./255., 45./255., 77./255., 0.5+0.5*float64(t.Depth)/float64(maxdepth))
				dc.Fill()
			}

		}
	}
	dc.SavePNG("out.png")
	fmt.Println("Done")
}

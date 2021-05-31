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

func UniformTriangleCenteredOn(x, y, d float64) geometry.Triangle {
	triangleSideLength := math.Cos(math.Pi/6) * 2. * d
	triangleVerticalSep := math.Sqrt((3. / 4.) * math.Pow(triangleSideLength, 2))
	// TriangleHorizSep := triangleSideLength / 2.

	y += (triangleVerticalSep - d) / 2

	theta := make([]float64, 0)
	for i := 0; i < 3; i++ {
		newTheta := float64(i)*math.Pi*(2./3.) - math.Pi/2.
		theta = append(theta, newTheta)
	}

	return geometry.Triangle{
		geometry.Point{x + d*math.Cos(theta[0]), y + d*math.Sin(theta[0])},
		geometry.Point{x + d*math.Cos(theta[1]), y + d*math.Sin(theta[1])},
		geometry.Point{x + d*math.Cos(theta[2]), y + d*math.Sin(theta[2])},
		false, false, false,
		0,
	}
}

func DrawGroup(dc *gg.Context, t geometry.Triangle) {
	b := behavior.RandomBehavior()
	b.Init(t)
	dividableTriangles := []clevertriangle.CleverTriangle{clevertriangle.WrapTriangle(
		t,
		b,
	)}

	for {
		newtriangles := make([]clevertriangle.CleverTriangle, 0)
		for _, t := range dividableTriangles {
			newtriangles = append(newtriangles, t.Cycle()...)
		}

		if len(newtriangles) == len(dividableTriangles) {
			break
		}
		dividableTriangles = newtriangles
	}

	maxdepth := 0
	for _, t := range dividableTriangles {
		if t.Depth > maxdepth {
			maxdepth = t.Depth
		}
	}

	for _, t := range dividableTriangles {
		t.Draw(dc)
		// 46, 45, 77
		dc.SetRGBA(46./255., 45./255., 77./255., 0.5+0.5*float64(t.Depth)/float64(maxdepth))
		dc.Fill()
		dc.SetRGBA(1, 1, 1, 1)
		dc.SetLineWidth(2)
		t.Draw(dc)
		dc.Stroke()
	}
}

func main() {
	fmt.Println("Let's party")
	seed := time.Now().UnixNano()
	fmt.Printf("Using seed [%d]\n", seed)
	rand.Seed(seed)

	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	rootTriangle := UniformTriangleCenteredOn(W/2, H/2, 200)
	triangles := make([]geometry.Triangle, 0)
	triangles = append(triangles, rootTriangle)

	for i := 0; i < 100; i++ {
		// Select and pass by reference
		t := &(triangles[rand.Intn(len(triangles))])
		newTriangle, ok := t.GrowRandomlyWithoutColliding(triangles)
		if ok {
			triangles = append(triangles, newTriangle)
		}
	}
	for _, t := range triangles {
		DrawGroup(dc, t)
	}

	dc.SavePNG("out.png")
	fmt.Println("Done")
}

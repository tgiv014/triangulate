package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/fogleman/gg"
	"github.com/tgiv014/triangulate/geometry"
)

const (
	DPI = 300.
	W   = (8. * DPI)
	H   = (10. * DPI)
)

func main() {
	fmt.Println("Let's party")
	seed := time.Now().UnixNano()
	fmt.Printf("Using seed [%d]\n", seed)
	rand.Seed(seed)

	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	rootTriangle := geometry.UniformTriangleCenteredOn(W/2, H/2, 200)
	triangles := []geometry.Triangle{
		rootTriangle,
	}
	t1 := rootTriangle.TangentTriangle(0, math.Pi/2, 100)
	t2 := rootTriangle.TangentTriangle(1, math.Pi/2, 100)
	t3 := rootTriangle.TangentTriangle(2, math.Pi/2, 100)
	triangles = append(triangles, t1, t2, t3)
	triangles = append(triangles, t1.TangentTriangle(1, math.Pi/2, 100))
	triangles = append(triangles, t2.TangentTriangle(1, math.Pi/2, 100))
	triangles = append(triangles, t3.TangentTriangle(1, math.Pi/2, 100))

	for _, t := range triangles {
		dc.SetRGBA(0, 0, 0, 1)
		dc.SetLineWidth(2)
		t.Draw(dc)
		dc.Stroke()
		t.HighlightUsed(dc)
	}
	dc.SavePNG("out.png")
	fmt.Println("Done")
}

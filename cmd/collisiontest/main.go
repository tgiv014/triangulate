package main

import (
	"fmt"
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
	if rootTriangle.PtInside(geometry.Point{W / 2, H / 2}) {
		fmt.Println("Yup")
	}
	if !rootTriangle.PtInside(geometry.Point{0, 0}) {
		fmt.Println("It works")
	}

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

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
	triangles := make([]geometry.Triangle, 0)
	triangles = append(triangles, rootTriangle)

	for i := 0; i < 100; i++ {
		// Select a random triangle *BY REFERENCE*
		t := &(triangles[rand.Intn(len(triangles))])
		// Attempt to grow a new tangential triangle and
		// This function does modify the original triangle
		// Thou shalt ensure that the OG triangle is passed by reference
		newTriangle, ok := t.GrowRandomlyWithoutColliding(triangles)

		// If it worked, append it!
		if ok {
			triangles = append(triangles, newTriangle)
		}
	}

	for _, t := range triangles {
		dc.SetRGBA(0, 0, 0, 1)
		dc.SetLineWidth(2)
		t.Draw(dc)
		dc.Stroke()
	}
	dc.SavePNG("out.png")
	fmt.Println("Done")
}

package main

import (
	"fmt"
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
	H   = (8. * DPI)
)

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

	rootTriangle := geometry.UniformTriangleCenteredOn(W/2, H/2, 200)
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

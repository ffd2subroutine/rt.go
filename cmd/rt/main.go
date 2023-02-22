package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/ffd2subroutine/rt.go/rt"
)

const (
	// Image
	AspectRation    float64 = 16.0 / 9.0
	ImageWidth      int     = 400
	ImageHeight     int     = int(float64(ImageWidth) / AspectRation)
	SamplesPerPixel int     = 100
)

func main() {
	// World
	world := rt.NewHittableList()
	world.Add(rt.NewSphere(rt.NewVec3(0.0, 0.0, -1.0), 0.5))
	world.Add(rt.NewSphere(rt.NewVec3(0.0, -100.5, -1.0), 100.0))

	// Camera
	cam := rt.NewCamera()

	// Render
	fmt.Fprintf(os.Stdout, "P3\n %d %d\n255\n", ImageWidth, ImageHeight)
	for j := ImageHeight - 1; j >= 0; j-- {
		for i := 0; i < ImageWidth; i++ {
			pixelColor := rt.NewVec3(0.0, 0.0, 0.0)
			for s := 0; s < SamplesPerPixel; s++ {
				u := (float64(i) + rand.Float64()) / float64(ImageWidth-1)
				v := (float64(j) + rand.Float64()) / float64(ImageHeight-1)
				r := cam.GetRay(u, v)
				pixelColor = rt.RayColor(r, world).Add(pixelColor)
			}
			writeColor(os.Stdout, pixelColor, SamplesPerPixel)
		}
	}
}

// TODO: Move to somewhere else, maybe to render.go ?
func writeColor(w io.Writer, pixelColor rt.Vec3, samplesPerPixel int) {
	r := pixelColor.X
	g := pixelColor.Y
	b := pixelColor.Z

	scale := 1.0 / float64(samplesPerPixel)
	r *= scale
	g *= scale
	b *= scale

	fmt.Fprintf(w,
		"%d %d %d\n",
		int(256*rt.Clamp(r, 0.0, 0.999)),
		int(256*rt.Clamp(g, 0.0, 0.999)),
		int(256*rt.Clamp(b, 0.0, 0.999)),
	)
}

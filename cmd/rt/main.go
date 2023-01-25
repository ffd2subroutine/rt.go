package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ffd2subroutine/rt.go/rt"
)

const (
	// Image
	AspectRation float64 = 16.0 / 9.0
	ImageWidth   int     = 400
	ImageHeight  int     = int(float64(ImageWidth) / AspectRation)
)

func main() {
	// World
	world := rt.NewHittableList()
	world.Add(rt.NewSphere(rt.NewVec3(0.0, 0.0, -1.0), 0.5))
	world.Add(rt.NewSphere(rt.NewVec3(0.0, -100.5, -1.0), 100.0))

	// Camera
	viewportHeight := 2.0
	viewportWidth := AspectRation * viewportHeight
	focalLength := 1.0
	origin := rt.Vec3{}
	horizontal := rt.NewVec3(viewportWidth, 0.0, 0.0)
	vertical := rt.NewVec3(0.0, viewportHeight, 0.0)
	lowerLeftCorner := origin.Sub(horizontal.DivS(2.0)).Sub(vertical.DivS(2.0)).Sub(rt.NewVec3(0.0, 0.0, focalLength))

	// Render
	fmt.Fprintf(os.Stdout, "P3\n %d %d\n255\n", ImageWidth, ImageHeight)
	for j := ImageHeight - 1; j >= 0; j-- {
		//fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d", j)
		for i := 0; i < ImageWidth; i++ {
			u := float64(i) / float64(ImageWidth-1)
			v := float64(j) / float64(ImageHeight-1)
			ray := rt.NewRay(origin, lowerLeftCorner.Add(horizontal.MulS(u)).Add(vertical.MulS(v)).Sub(origin))
			pixelColor := rt.RayColor(ray, world)
			writeColor(os.Stdout, pixelColor)
		}
	}
}

// TODO: move to output.go
func writeColor(w io.Writer, pixelColor rt.Vec3) {
	fmt.Fprintf(w,
		"%d %d %d\n",
		int(255.999*pixelColor.X),
		int(255.999*pixelColor.Y),
		int(255.999*pixelColor.Z),
	)
}

package main

import (
	"fmt"
	"io"
	"os"

	"github.com/ffd2subroutine/rt.go/rt"
)

const (
	ImageWidth  int = 256
	ImageHeight int = 256
)

func main() {
	fmt.Fprintf(os.Stdout, "P3\n %d %d\n255\n", ImageWidth, ImageHeight)
	for j := ImageHeight - 1; j >= 0; j-- {
		//fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d", j)
		for i := 0; i < ImageWidth; i++ {
			pixelColor := rt.Vec3{
				X: float64(i) / float64((ImageWidth - 1)),
				Y: float64(j) / float64((ImageHeight - 1)),
				Z: 0.25,
			}
			writeColor(os.Stdout, pixelColor)
		}
	}
}

func writeColor(w io.Writer, pixelColor rt.Vec3) {
	fmt.Fprintf(w,
		"%d %d %d\n",
		int(255.999*pixelColor.X),
		int(255.999*pixelColor.Y),
		int(255.999*pixelColor.Z),
	)
}

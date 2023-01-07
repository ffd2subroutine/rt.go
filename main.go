package main

import (
	"fmt"
	"os"
)

const (
	ImageWidth  int = 256
	ImageHeight int = 256
)

func main() {
	fmt.Fprintf(os.Stdout, "P3\n %d %d\n255\n", ImageWidth, ImageHeight)
	for j := ImageHeight - 1; j >= 0; j-- {
		//fmt.Fprintf(os.Stderr, "\rScanlines remaining: %d", j)
		var ir int
		var ig int
		var ib int
		for i := 0; i < ImageWidth; i++ {
			r := float64(i) / float64((ImageWidth - 1))
			g := float64(j) / float64((ImageHeight - 1))
			b := 0.25

			ir = int(255.999 * r)
			ig = int(255.999 * g)
			ib = int(255.999 * b)

			fmt.Fprintf(os.Stdout, "%d %d %d\n", ir, ig, ib)
		}
	}
}

package main

import (
	"fmt"
	"math"
	// "math/cmplx"
	"strconv"
)

func main() {

	// image_spec := TGAImageSpec {0, 0, 512, 512, 24, 0}
	// color_map := TGAColorMapSpec {0, 0, 0}
	// header := TGAHeader {0, 0, 1, color_map, image_spec}
	const image_w = 1024
	const image_h = 1024
	image_data := make([]byte, 3*image_w*image_h)
	potential := make([]complex128, image_w*image_h)
	px, py := MakePxPy(image_w, image_h)
	// psi := MakeWavePacket(0.5, 0.5, 0.1, 0.1, 30.0, 40.0, image_w, image_h)
	psi := MakeRandom(image_w, image_h)
	psi_norm_factor := math.Sqrt(NormSquared(psi[:]))
	dt := complex(0.5, -0.05)

	for i, j := 0, 0; i < 1000; i++ {
		SpatialStep(psi[:], potential[:], 0.5*dt)
		MomentumStep(psi[:], px[:], py[:], image_w, image_h, dt)
		SpatialStep(psi[:], potential[:], 0.5*dt)
		// Normalize(psi[:], image_w, image_h)
		norm_val := math.Sqrt(NormSquared(psi[:]))
		Scale(psi[:], psi_norm_factor/norm_val)
		if i % 2 == 0 {
			FillImageDataFromComplexData(
				image_data[0:3*image_w*image_h],
				psi[0:image_w*image_h], image_w, image_h)
			filename := "./image"
			if j < 10 {
				filename += "00" + strconv.FormatInt(int64(j), 10) + ".tga"
			} else if j < 100 {
				filename += "0" + strconv.FormatInt(int64(j), 10) + ".tga"
			} else if j < 1000 {
				filename += strconv.FormatInt(int64(j), 10) + ".tga"
			}
			fmt.Println("Saving " + filename + ".")
			WriteTGA(filename,
					image_data[0:3*image_w*image_h],
					image_w, image_h)
			j++
		}
	}
}

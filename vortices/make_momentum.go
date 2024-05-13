package main

import "math"

func MakePxPy(w, h int) (px, py []float64) {
	px = make([]float64, w*h)
	py = make([]float64, w*h)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			i_shift := i
			if i >= h/2 {
				i_shift = -h + i
			}
			j_shift := j
			if j >= w/2 {
				j_shift = -w + j
			}
			ind := i*w + j
			px[ind] = 2.0*math.Pi*float64(j_shift)/float64(w)
			py[ind] = 2.0*math.Pi*float64(i_shift)/float64(h)
		}
	}
	return
}

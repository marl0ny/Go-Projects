package main

func InPlaceTranspose(array[] complex128, w, h int) {
	if (w == h) {
		size := w
		for i := 0; i < size; i++ {
			for j := i+1; j < size; j++ {
				ind1, ind2 := i*size + j, j*size + i
				array[ind1], array[ind2] = array[ind2], array[ind1]
			}
		}
	} else { // Brute force copy to temporary and copy it back
		tmp := make([]complex128, w*h)
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				tmp[j*h + i] = array[i*w + j] 
			}
		}
		for i := 0; i < w*h; i++ {
			array[i] = tmp[i]
		}
	}
}

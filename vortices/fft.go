/* Functions for taking the FFT of array slices.

References:

Wikipedia - Cooley–Tukey FFT algorithm
https://en.wikipedia.org/wiki/Cooley%E2%80%93Tukey_FFT_algorithm

MathWorld Wolfram - Fast Fourier Transform:
http://mathworld.wolfram.com/FastFourierTransform.html

William Press et al.
12.2 Fast Fourier Transform (FFT) - Numerical Recipes
https://websites.pmc.ucsc.edu/~fnimmo/eart290c_17/NumericalRecipesinF77.pdf

*/
package main

import (
	"math"
	"math/cmplx"
)

func ReverseBitSort2(z []complex128) {
	n := len(z)
	var u, d, rev int
	for i := 0; i < n; i++ {
		u = 1
		d = n >> 1
		rev = 0
		for u < n {
			rev += d*((i&u)/u)
			u <<= 1
			d >>= 1
		}
		if rev >= i {
			z[i], z[rev] = z[rev], z[i]
		}
	}
}

func InPlaceRadix2FFT(array[] complex128, is_inverse bool) {
	ReverseBitSort2(array)
	for block_size := 2; block_size <= len(array); block_size *= 2 {
		for j := 0; j < len(array); j += block_size {
			// range block_size/2 did not work
			for i := 0; i < block_size/2; i++ {
				sgn := 1.0 // No such thing as ternary or if else statements
				if is_inverse {
					sgn = -1.0
				}
				var e complex128 = cmplx.Exp(
					complex(0.0,
						sgn*2.0*math.Pi*float64(i)/float64(block_size)))
				even := array[j + i]
				odd := array[j + i + block_size/2]
				if is_inverse && block_size == len(array) {
					array[j + i] = (
						(even + odd*e)/complex(float64(len(array)), 0.0))
					array[j + i + block_size/2] = (
						(even - odd*e)/complex(float64(len(array)), 0.0))
				} else {
					array[j + i] = even + odd*e
					array[j + i + block_size/2] = even - odd*e
				}
			}
		}
	}
}

func InPlaceFFTAlongRows(array[] complex128, w, h int,
				         is_inverse bool) {
	thread_count := GetGlobalNumberOfThreads()
	FFTConcurrent := func(array[] complex128, is_inverse bool,
		n_rows, w int, c chan int) {
		for i := 0; i < n_rows; i++ {
			InPlaceRadix2FFT(array[i*w: (i+1)*w], is_inverse)
		}
		c <- 0
	}
	c := make(chan int, thread_count)
	for i := 0; i < thread_count; i++ {
		n_rows := h/thread_count
		go FFTConcurrent(array[i*n_rows*w: (i+1)*n_rows*w],
					     is_inverse, n_rows, w, c)
	}
	for i := 0; i < thread_count; i++ {
		<-c
	}
}

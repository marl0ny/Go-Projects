/* Implementation of the split operator method.

References:

Split-Operator Method:
James Schloss. The Split Operator Method - Arcane Algorithm Archive.
https://www.algorithm-archive.org/contents/split-operator_method/
 split-operator_method.html

 Xavier Antoine, Weizhu Bao, Christophe Besse.
 Computational methods for the dynamics of the
 nonlinear Schrodinger/Gross-Pitaevskii equations.
 https://arxiv.org/abs/1305.1093

 */
package main

import "math/cmplx"

func Nonlinear(psi complex128) complex128 {
	return 0.1*cmplx.Conj(psi)*psi
}

func SpatialStepConcurrent(psi, potential [] complex128,
	dt complex128, c chan int) {
	for i := 0; i < len(psi); i++ {
		psi[i] *= cmplx.Exp(-(1i*potential[i] + 1i*Nonlinear(psi[i]))*dt)
	}
	c <- 0
}

func SpatialStep(psi, potential [] complex128, dt complex128) {
	n_routines := GetGlobalNumberOfRoutines()
	c := make(chan int, n_routines)
	for i := 0; i < n_routines; i++ {
		ind1, ind2 := (i*len(psi))/n_routines, ((i+1)*len(psi))/n_routines
		go SpatialStepConcurrent(psi[ind1: ind2], potential[ind1: ind2], dt,
			c)
	}
	for i := 0; i < n_routines; i++ {
		<- c
	}
}

func PropagateMomentumConcurrent(psi []complex128,
	px, py []float64, w, h int, dt complex128, c chan int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ind := w*i + j
			p2 := px[ind]*px[ind] + py[ind]*py[ind]
			psi[ind] *= cmplx.Exp(-(0.5i*complex(p2, 0.0))*dt)
		}
	}
	c <- 0
}

func MomentumStep(psi []complex128, px, py []float64,
				  w, h int,
				  dt complex128) {
	InPlaceFFTAlongRows(psi[:], w, h, false)
	InPlaceTranspose(psi[:], w, h)
	InPlaceFFTAlongRows(psi[:], h, w, false)
	InPlaceTranspose(psi[:], h, w)
	n_routines := GetGlobalNumberOfRoutines()
	c := make(chan int, n_routines)
	for i := 0; i < n_routines; i++ {
		row_width, row_count := w, h/n_routines
		ind1, ind2 := i*row_count*row_width, (i + 1)*row_count*row_width
		go PropagateMomentumConcurrent(
			psi[ind1:ind2], px[ind1:ind2], py[ind1:ind2],
			row_width, row_count, dt, c)
	}
	for i := 0; i < n_routines; i++ {
		<- c
	}
	InPlaceFFTAlongRows(psi[:], w, h, true)
	InPlaceTranspose(psi[:], w, h)
	InPlaceFFTAlongRows(psi[:], h, w, true)
	InPlaceTranspose(psi[:], h, w)
}

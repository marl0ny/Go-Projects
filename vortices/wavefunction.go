package main

import (
	"math"
	"math/cmplx"
	"math/rand"
)

func MakeWavePacket(x0, y0, sx, sy, nx, ny float64,
					w, h int) ([] complex128) {
	psi := make([]complex128, w*h)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ind := i*w + j
			x := float64(j)/float64(w)
			y := float64(i)/float64(h)
			xs := x - x0
			ys := y - y0
			psi[ind] = cmplx.Exp(complex(-0.5*(xs*xs + ys*ys)/(sx*sy), 0.0))
			psi[ind] *= cmplx.Exp(complex(0.0, 2.0*math.Pi*(nx*x + ny*y)))
		}
	}
	return psi
}

func MakeRandom(w, h int) []complex128 {
	psi := make([]complex128, w*h)
	for i := 0; i < len(psi); i++ {
		angle := 2.0*math.Pi*rand.Float64()
		psi[i] = complex(math.Cos(angle), math.Sin(angle))
	}
	return psi
}

func ScaleConcurrent(psi []complex128, scale_val float64, c chan int) {
	for i := 0; i < len(psi); i++ {
		psi[i] *= complex(scale_val, 0.0)
	}
	c <- 0
}

func Scale(psi []complex128, scale_val float64) {
	n_routines := GetGlobalNumberOfRoutines()
	c := make(chan int, n_routines)
	for i := 0; i < n_routines; i++ {
		ind1, ind2 := (i*len(psi))/n_routines, ((i+1)*len(psi))/n_routines
		ScaleConcurrent(psi[ind1: ind2], scale_val, c)
	}
	for i := 0; i < n_routines; i++ {
		<- c
	}
}


func NormSquaredConcurrent(psi []complex128, c chan float64) {
	var sum float64 = 0.0
	for i := 0; i < len(psi); i++ {
		sum += real(cmplx.Conj(psi[i])*psi[i])
	}
	c <- sum
}

func NormSquared(psi []complex128) float64 {
	n_routines := GetGlobalNumberOfRoutines()
	var sum float64 = 0.0
	c := make(chan float64, n_routines)
	for i := 0; i < n_routines; i++ {
		ind1, ind2 := (len(psi)*i)/n_routines, (len(psi)*(i+1))/n_routines
		go NormSquaredConcurrent(psi[ind1: ind2], c)
	}
	for i := 0; i < n_routines; i++ {
		sum += <- c
	}
	return sum
}


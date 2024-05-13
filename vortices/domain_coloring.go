/* Domain colour complex data.

References:

Wikipedia - Domain coloring
https://en.wikipedia.org/wiki/Domain_coloring

Wikipedia - Hue
https://en.wikipedia.org/wiki/Hue

https://en.wikipedia.org/wiki/Hue#/media/File:HSV-RGB-comparison.svg

*/
package main

import (
	"math"
	"math/cmplx"
)

type Color struct {
	r, g, b float64
}

func ArgToColor(val float64) Color {
    pi := math.Pi
    max_col := 1.0
    min_col := 50.0/255.0
    col_range := max_col - min_col
    if val <= pi/3.0 && val >= 0.0 {
        return Color {
            max_col,
            min_col + col_range*val/(pi/3.0),
            min_col}
    } else if val > pi/3.0 && val <= 2.0*pi/3.0 {
        return Color {
            max_col - col_range*(val - pi/3.0)/(pi/3.0),
            max_col,
            min_col}
    } else if val > 2.0*pi/3.0 && val <= pi {
        return Color {
            min_col,
            max_col,
            min_col + col_range*(val - 2.0*pi/3.0)/(pi/3.0)};
    } else if val < 0.0 && val > -pi/3.0 {
        return Color {
            max_col,
            min_col,
            min_col - col_range*val/(pi/3.0)};
    } else if val <= -pi/3.0 && val > -2.0*pi/3.0 {
        return Color {
            max_col + (col_range*(val + pi/3.0)/(pi/3.0)),
            min_col,
            max_col}
    } else if val <= -2.0*pi/3.0 && val >= -pi {
        return Color {
            min_col,
            min_col - (col_range*(val + 2.0*pi/3.0)/(pi/3.0)),
            max_col}
    } else {
        return Color {min_col, max_col, max_col}
    }
}

func FillImageDataFromComplexData(image []byte, arr []complex128, w, h int) {
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ind := i*w + j
			arg := cmplx.Phase(arr[ind])
			color := ArgToColor(arg)
			image[3*ind] = byte(min(
                255.0, 255.0*cmplx.Abs(arr[ind])*color.b))
			image[3*ind+1] = byte(min(
                255.0, 255.0*cmplx.Abs(arr[ind])*color.g))
			image[3*ind+2] = byte(min(
                255.0, 255.0*cmplx.Abs(arr[ind])*color.r))
		}
	}
}

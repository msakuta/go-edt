package main

import (
	"image"
	"image/color"
)

func make_image(buffer []float64, shape []int) *image.NRGBA {
	imag := image.NewNRGBA(image.Rect(0, 0, shape[0], shape[1]))

	max_val := 0.
	for i := 0; i < shape[0]*shape[1]; i++ {
		if max_val < buffer[i] {
			max_val = buffer[i]
		}
	}

	for y := 0; y < shape[1]; y++ {
		for x := 0; x < shape[0]; x++ {
			pixel := color.Gray{
				Y: uint8(buffer[x+y*shape[0]] * 255 / max_val),
			}
			imag.Set(x, y, pixel)
		}
	}

	return imag
}

func from_image(img image.Image) (buffer []bool, shape [2]int) {
	shape = [2]int{img.Bounds().Dx(), img.Bounds().Dy()}
	buffer = make([]bool, shape[0]*shape[1])
	for y := 0; y < shape[1]; y++ {
		for x := 0; x < shape[0]; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			buffer[x+y*shape[0]] = r == 0
		}
	}
	return
}

package main

import (
	"image"
	"image/png"
	"log"
	"os"

	// "fmt"
	// "strings"
	"math"
	"time"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if b < a {
		return a
	}
	return b
}

func horizontal_edt(buffer []bool, shape [2]int) []float64 {
	horz_edt := make([]float64, shape[0]*shape[1])

	edt_max := float64(min(shape[0], shape[1]))
	count_true := 0

	for y := 0; y < shape[1]; y++ {
		for x := 0; x < shape[0]; x++ {
			if buffer[x+y*shape[0]] {
				horz_edt[x+y*shape[0]] = edt_max
				count_true += 1
			}
		}
	}

	println("count_true: ", count_true)

	max_val := 0.

	scan := func(x, y int, min_val *float64) {
		f := horz_edt[x+y*shape[0]]
		next := *min_val + 1.
		v := math.Min(f, next)
		horz_edt[x+y*shape[0]] = v
		*min_val = v
		if max_val < v {
			max_val = v
		}
	}

	for y := 0; y < shape[1]; y++ {
		min_val := 0.
		for x := 0; x < shape[0]; x++ {
			scan(x, y, &min_val)
		}
		min_val = 0.
		for x := shape[0] - 1; 0 <= x; x-- {
			scan(x, y, &min_val)
		}
	}

	println("Hello max: ", max_val)

	return horz_edt
}

func edt(buffer []bool, shape [2]int) []float64 {
	horz_edt := horizontal_edt(buffer, shape)
	max_edt := float64(max(shape[0], shape[1]))
	max_edt = max_edt * max_edt

	vertical_scan := func(x, y int) float64 {
		total_edt := max_edt
		for y2 := 0; y2 < shape[1]; y2++ {
			horz_val := horz_edt[x+y2*shape[0]]
			dy := float64(y2) - float64(y)
			val := dy*dy + horz_val*horz_val
			if val < total_edt {
				total_edt = val
			}
		}
		dy := float64(y)
		if dy*dy < total_edt {
			return dy * dy
		}
		dy = float64(shape[1] - y)
		if dy*dy < total_edt {
			return dy * dy
		}
		return total_edt
	}

	ret := make([]float64, shape[0]*shape[1])
	copy(ret, horz_edt)

	for x := 0; x < shape[0]; x++ {
		for y := 0; y < shape[1]; y++ {
			ret[x+y*shape[0]] = vertical_scan(x, y)
		}
	}

	return ret
}

func main() {
	shape := [...]int{512, 512}

	use_shape := ""
	if len(os.Args) == 1 {
		use_shape = "circle"
	} else {
		use_shape = os.Args[1]
	}

	println("Selected shape:", use_shape)

	var img []bool
	switch use_shape {
	case "circle":
		img = get_circle(shape)
	case "cross":
		img = get_cross(shape)
	default:
		f, err := os.Open("Go_Logo.png")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer f.Close()
		png_img, fmt, err := image.Decode(f)
		println("Format: ", fmt)
		if err != nil {
			log.Fatal(err)
			return
		}
		img, shape = from_image(png_img)
	}

	println("Shape: ", shape[0], shape[1])

	start := time.Now()

	horz_edt := edt(img, shape)

	elapsed := time.Since(start)

	log.Printf("time: %g", float64(elapsed)/1e6)

	imag := make_image(horz_edt, shape[:])

	file, err := os.Create("edt.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	png.Encode(file, imag)
}

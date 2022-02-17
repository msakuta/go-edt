package main

func get_circle(shape [2]int) []bool {
	ret := make([]bool, shape[0]*shape[1])

	for y := 0; y < shape[1]; y++ {
		for x := 0; x < shape[0]; x++ {
			dx := x - shape[0]/2
			dy := y - shape[1]/2
			ret[x+y*shape[0]] = dx*dx+dy*dy < shape[0]*shape[0]/4.
		}
	}

	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func get_cross(shape [2]int) []bool {
	ret := make([]bool, shape[0]*shape[1])

	for y := 0; y < shape[1]; y++ {
		for x := 0; x < shape[0]; x++ {
			dx := x - shape[0]/2
			dy := y - shape[1]/2
			ret[x+y*shape[0]] = abs(dx) < shape[0]/4 || abs(dy) < shape[1]/4
		}
	}

	return ret
}

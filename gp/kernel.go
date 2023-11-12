package gp

import "math"

type Kernel = func([]float64, []float64) float64

func RBF_KERNEL(l float64) Kernel {
	return func(xn, xm []float64) float64 {
		var sum float64 = 0
		for i := range xn {
			var diff = xn[i] - xm[i]
			sum += (diff * diff)
		}
		return math.Exp(-sum / (2 * l * l))

	}
}

func CONSTANT_KERNEL(c float64) Kernel {
	return func(xn, xm []float64) float64 {
		return c
	}
}

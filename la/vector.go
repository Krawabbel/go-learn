package la

func RowVec(vec ...float64) Matrix {
	if len(vec) > 0 {
		return Mat([][]float64{vec})
	}
	return nil
}

func ColVec(vec ...float64) Matrix {
	if len(vec) > 0 {
		return Transpose(Mat([][]float64{vec}))
	}
	return nil
}

func Linspace(a, b float64, N int) Matrix {
	if N < 2 {
		return nil
	}
	data := make([]float64, N)
	dx := (b - a) / float64(N-1)
	for i := range data {
		data[i] = a + float64(i)*dx
	}
	return RowVec(data...)
}

package la

import "fmt"

func Mult(M, N Matrix) (Matrix, error) {
	if M == nil || N == nil {
		return nil, fmt.Errorf("cannot multiply nil matrices")
	}

	if M.cols() != N.rows() {
		return nil, fmt.Errorf("cannot multiply %dx%d and %dx%d matrices", M.rows(), M.cols(), N.rows(), N.cols())
	}

	at := func(row, col int) float64 {
		sum := 0.0
		for k := 0; k < M.cols(); k++ {
			sum += M.at(row, k) * N.at(k, col)
		}
		return sum
	}

	return View{M.rows(), N.cols(), at}.store(), nil
}

func Add(M, N Matrix) (Matrix, error) {

	if M == nil || N == nil {
		return nil, fmt.Errorf("cannot add nil matrices")
	}

	if M.cols() != N.cols() || M.rows() != N.rows() {
		return nil, fmt.Errorf("cannot add %dx%d and %dx%d matrices", M.rows(), M.cols(), N.rows(), N.cols())
	}

	at := func(row, col int) float64 {
		return M.at(row, col) + N.at(row, col)
	}

	return View{M.rows(), N.cols(), at}.store(), nil
}

func Scale(c float64, M Matrix) (Matrix, error) {

	if M == nil {
		return nil, fmt.Errorf("cannot scale nil matrix")
	}

	at := func(row, col int) float64 {
		return c * M.at(row, col)
	}

	return View{M.rows(), M.cols(), at}.store(), nil
}

package matrix

import "fmt"

func Lower(M Matrix) Matrix {
	at := func(row, col int) float64 {
		if col <= row {
			return M.at(row, col)
		}
		return 0
	}
	return View{nRows: M.rows(), nCols: M.cols(), at_fun: at}
}

func Upper(M Matrix) Matrix {
	at := func(row, col int) float64 {
		if col >= row {
			return M.at(row, col)
		}
		return 0
	}
	return View{nRows: M.rows(), nCols: M.cols(), at_fun: at}
}

func lower_solve(L, Y Matrix) (Matrix, error) {

	if L == nil || Y == nil {
		return nil, fmt.Errorf("lower solve error: unexpected nil matrix")
	}

	if !IsSquare(L) {
		return nil, fmt.Errorf("lower solve error: lower matrix L must be square")
	}

	dim := L.rows()

	if Y.rows() != dim {
		return nil, fmt.Errorf("lower solve error: L has %d rows and Y has %d rows (must be equal)", dim, Y.rows())
	}

	num := Y.cols()

	x := make([][]float64, dim)
	for row := range x {
		x[row] = make([]float64, num)
		for col := range x[row] {
			var sum float64 = 0

			for k := 0; k < row; k++ {
				sum += L.at(row, k) * x[k][col]
			}
			x[row][col] = (Y.at(row, col) - sum) / L.at(row, row)

		}
	}
	return Dense{data: x}, nil
}

func upper_solve(U, Y Matrix) (Matrix, error) {

	if U == nil || Y == nil {
		return nil, fmt.Errorf("upper solve error: unexpected nil matrix")
	}

	if !IsSquare(U) {
		return nil, fmt.Errorf("upper solve error: upper matrix U must be square")
	}

	dim := U.rows()

	if Y.rows() != dim {
		return nil, fmt.Errorf("upper solve error: U has %d rows and Y has %d rows (must be equal)", dim, Y.rows())
	}

	num := Y.cols()

	x := make([][]float64, dim)

	for i := range x {
		row := dim - 1 - i
		x[row] = make([]float64, num)
		for j := range x[row] {
			col := num - 1 - j
			var sum float64 = 0

			for k := dim - 1; k > row; k-- {
				sum += U.at(row, k) * x[k][col]
			}
			x[row][col] = (Y.at(row, col) - sum) / U.at(row, row)
		}
	}

	return Dense{data: x}, nil

}

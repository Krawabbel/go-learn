package la

import "fmt"

const (
	CHOLESKY = iota
	LEAST_SQUARES
	LOWER
	UPPER
)

func Solve(A, Y Matrix, solver int) (X Matrix, err error) {
	switch solver {
	case CHOLESKY:
		X, err = chol_solve(A, Y)
	case LEAST_SQUARES:
		X, err = ls_solve(A, Y)
	case LOWER:
		X, err = lower_solve(A, Y)
	case UPPER:
		X, err = upper_solve(A, Y)
	default:
		return nil, fmt.Errorf("unknown solver")
	}
	return
}

func ls_solve(A, Y Matrix) (Matrix, error) {

	AT := Transpose(A)

	PHI, err := Mult(AT, A)
	if err != nil {
		return nil, err
	}

	PSI, err := Mult(AT, Y)
	if err != nil {
		return nil, err
	}

	X, err := chol_solve(PHI, PSI)
	return X, err
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

package la

import (
	"fmt"
	"math"
)

func chol_dec(M Matrix) (Matrix, error) {

	if M == nil {
		return nil, fmt.Errorf("cannot decompose nil matrix")
	}

	if !IsSymmetric(M) {
		return nil, fmt.Errorf("matrix is not symmetric")
	}

	L := make([][]float64, M.rows())
	for i := range L {
		L[i] = make([]float64, i+1)
		for j := range L[i] {

			var sum float64 = 0
			for k := 0; k < j; k++ {
				sum += L[i][k] * L[j][k]
			}

			if M.at(i, i) < sum {
				return nil, fmt.Errorf("matrix is not positive definite (delta = %v)", M.at(i, i)-sum)
			}

			if i == j {
				L[i][j] = math.Sqrt(M.at(i, i) - sum)
			} else {
				L[i][j] = (M.at(i, j) - sum) / L[j][j]
			}
		}
	}

	dim := M.rows()

	at := func(row, col int) float64 {
		if col <= row {
			return L[row][col]
		}
		return 0
	}

	return View{nRows: dim, nCols: dim, at_fun: at}, nil
}

func chol_solve(A, Z Matrix) (Matrix, error) {

	if Z == nil || A == nil {
		return nil, fmt.Errorf("invalid nil argument")
	}

	L, err := chol_dec(A)
	if err != nil {
		return nil, err
	}

	Y, err := lower_solve(L, Z)
	if err != nil {
		return nil, err
	}

	U := Transp(L)

	X, err := upper_solve(U, Y)
	if err != nil {
		return nil, err // probably unreachable
	}

	return X, nil
}

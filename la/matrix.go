package la

import (
	"fmt"
	"strings"
)

type Matrix interface {
	rows() int
	cols() int
	at(row int, col int) float64
}

func Transp(M Matrix) Matrix {
	if M == nil {
		return nil
	}
	at := func(row, col int) float64 {
		return M.at(col, row)
	}
	return View{M.cols(), M.rows(), at}
}

func IsSquare(M Matrix) bool {
	return M != nil && (M.cols() == M.rows())
}

func IsSymmetric(M Matrix) bool {

	if M == nil {
		return false
	}

	if !IsSquare(M) {
		return false
	}

	for i := 0; i < M.rows(); i++ {
		for j := 0; j < i; j++ {
			if M.at(i, j) != M.at(j, i) {
				return false
			}
		}
	}

	return true
}

func String(M Matrix) string {
	if M == nil {
		return "{}"
	}
	s := make([]string, M.rows())
	for row := range s {
		t := make([]string, M.cols())
		for col := range t {
			t[col] = fmt.Sprintf("%v", M.at(row, col))
		}
		s[row] = "{" + strings.Join(t, ", ") + "}"
	}
	return "{" + strings.Join(s, ", ") + "}"
}

func Data(M Matrix) [][]float64 {
	if M == nil {
		return nil
	}
	data := make([][]float64, M.rows())
	for row := range data {
		data[row] = make([]float64, M.cols())
		for col := range data[row] {
			data[row][col] = M.at(row, col)
		}
	}
	return data
}

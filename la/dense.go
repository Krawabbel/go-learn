package la

import "math/rand"

type Dense struct {
	data [][]float64
}

func (d Dense) rows() int {
	return len(d.data)
}

func (d Dense) cols() int {
	if d.rows() > 0 {
		return len(d.data[0])
	}
	return 0
}

func (d Dense) at(row, col int) float64 {
	return d.data[row][col]
}

func Empty() Matrix {
	return Dense{}
}

func Rand(rows, cols int) Matrix {
	data := make([][]float64, rows)
	for row := range data {
		data[row] = make([]float64, cols)
		for col := range data[row] {
			data[row][col] = rand.Float64()
		}
	}
	return Dense{data}
}

func RowVec(vec ...float64) Matrix {
	if len(vec) > 0 {
		return Dense{data: [][]float64{vec}}
	}
	return nil
}

func ColVec(vec ...float64) Matrix {
	if len(vec) > 0 {
		return Transpose(Dense{data: [][]float64{vec}})
	}
	return nil
}

func Scalar(scalar float64) Matrix {
	return Dense{data: [][]float64{{scalar}}}
}

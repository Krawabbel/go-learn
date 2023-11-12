package la

type Const struct {
	nRows, nCols int
	val          float64
}

func (c Const) rows() int {
	return c.nRows
}

func (c Const) cols() int {
	return c.nCols
}

func (c Const) at(row, col int) float64 {
	return c.val
}

func Zeros(rows, cols int) Matrix {
	return Const{nRows: rows, nCols: cols, val: 0}
}

func Ones(rows, cols int) Matrix {
	return Const{nRows: rows, nCols: cols, val: 1}
}

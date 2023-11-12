package matrix

type View struct {
	nRows, nCols int
	at_fun       func(row, col int) float64
}

func (v View) rows() int { return v.nRows }

func (v View) cols() int { return v.nCols }

func (v View) at(row int, col int) float64 {
	return v.at_fun(row, col)
}

func (v View) store() Matrix {
	return Dense{data: Data(v)}
}

package matrix

type LazyView struct {
	view  View
	known [][]bool
	data  [][]float64
}

func NewLazyView(rows, cols int, at func(row, col int) float64) *LazyView {
	view := View{nRows: rows, nCols: cols, at_fun: at}
	known := make([][]bool, view.rows())
	data := make([][]float64, view.rows())
	for i := 0; i < view.rows(); i++ {
		known[i] = make([]bool, view.cols())
		data[i] = make([]float64, view.cols())
	}
	return &LazyView{view: view, known: known, data: data}
}

func (l LazyView) rows() int { return l.view.rows() }

func (l LazyView) cols() int { return l.view.cols() }

func (l *LazyView) at(row int, col int) float64 {
	if !l.known[row][col] {
		l.data[row][col] = l.view.at(row, col)
		l.known[row][col] = true
	}
	return l.data[row][col]
}

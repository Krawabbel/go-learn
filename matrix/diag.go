package matrix

type Diagonal struct {
	diag []float64
}

func (d Diagonal) dim() int { return len(d.diag) }

func (d Diagonal) rows() int { return d.dim() }

func (d Diagonal) cols() int { return d.dim() }

func (d Diagonal) at(row, col int) float64 {
	if row == col {
		return d.diag[row]
	}
	return 0
}

func Diag(vals ...float64) Matrix {
	if len(vals) > 0 {
		return Diagonal{diag: vals}
	}
	return nil
}

func Eye(dim int) Matrix {
	diag := make([]float64, dim)
	for i := range diag {
		diag[i] = 1
	}
	return Diagonal{diag: diag}
}

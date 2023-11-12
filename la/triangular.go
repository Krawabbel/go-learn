package la

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

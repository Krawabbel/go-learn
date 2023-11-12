package gp

import "github.com/Krawabbel/go-learn/la"

func cov(X1, X2 [][]float64, kernel Kernel) la.Matrix {
	var K = make([][]float64, len(X1))
	for i := range K {
		K[i] = make([]float64, len(X2))
		for j := range K[i] {
			K[i][j] = kernel(X1[i], X2[j])
		}
	}
	return la.Mat(K)
}

func Predict(x_train [][]float64, y_train []float64, x_pred [][]float64, kernel Kernel, s float64) (Y_pred la.Matrix) {

	dim := len(y_train)

	S, err := la.Scale(s, la.Eye(dim))
	if err != nil {
		panic(err)
	}

	K_tt := cov(x_train, x_train, kernel)

	C, err := la.Add(K_tt, S)
	if err != nil {
		panic(err)
	}

	Y_train := la.ColVec(y_train...)

	W, err := la.Solve(C, Y_train, la.CHOLESKY_SOLVER)
	if err != nil {
		panic(err)
	}

	K_tp := cov(x_train, x_pred, kernel)

	Y_pred, err = la.Mult(la.Transp(K_tp), W)
	if err != nil {
		panic(err)
	}

	return
}

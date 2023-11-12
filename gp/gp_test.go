package gp

import (
	"testing"

	"github.com/Krawabbel/go-learn/la"
)

func TestGP(t *testing.T) {

	f_train := func(x []float64) float64 {
		return (x[0] - 5.0) * (x[0] - 5.0)
	}

	X_train := [][]float64{
		{1.0},
		{3.0},
		{5.0},
		{7.0},
		{9.0},
	}

	Y_train := make([]float64, len(X_train))
	for i, x := range X_train {
		Y_train[i] = f_train(x)
	}

	X_pred := [][]float64{
		{5.5},
	}

	l := 1.0
	s := 0.0

	have := Predict(X_train, Y_train, X_pred, RBF_KERNEL(l), s)

	want_data := make([]float64, len(X_pred))
	for i, x := range X_pred {
		want_data[i] = f_train(x)
	}
	want := la.ColVec(want_data...)

	la.EXPECT_EQ(t, have, want, 0.03)
}

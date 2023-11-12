package la

import (
	"math"
	"testing"
)

const (
	FLOAT_EQ_TOL = 1e-5
)

func EXPECT_EQ(t *testing.T, M, N Matrix, tol float64) {

	if M == nil && N == nil {
		return
	}

	if M == nil || N == nil {
		t.Errorf("unmatched nil matrix: %v != %v", M, N)
		return
	}

	if M.cols() != N.cols() || M.rows() != N.rows() {
		t.Errorf("dimension mismatch (%dx%d) != (%dx%d)", M.rows(), M.cols(), N.rows(), N.cols())
		return
	}

	for row := 0; row < M.rows(); row++ {
		for col := 0; col < M.cols(); col++ {
			if math.Abs(M.at(row, col)-N.at(row, col)) > tol {
				t.Errorf("entry mismatch with tolerance %v: %s != %s", tol, String(M), String(N))
			}
		}
	}
}

package la

import (
	"math"
	"reflect"
	"testing"
)

const (
	float_eq_tol = 1e-5
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

func TestIsSymmetric(t *testing.T) {
	type args struct {
		M Matrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nil", args{nil}, false},
		{"empty", args{Empty()}, true},
		{"scalar", args{Scalar(1)}, true},
		{"row vector", args{RowVec(1, 2, 3)}, false},
		{"col vector", args{ColVec(1, 2, 3)}, false},
		{"symmetric", args{Dense{[][]float64{{1, 2}, {2, 1}}}}, true},
		{"non-symmetric", args{Dense{[][]float64{{1, 2}, {-2, 1}}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSymmetric(tt.args.M); got != tt.want {
				t.Errorf("IsSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestData(t *testing.T) {
	type args struct {
		M Matrix
	}
	tests := []struct {
		name string
		args args
		want [][]float64
	}{
		{"nil", args{nil}, nil},
		{"empty", args{Empty()}, [][]float64{}},
		{"scalar", args{Scalar(0)}, [][]float64{{0}}},
		{"vector", args{RowVec(1, 2, 3)}, [][]float64{{1, 2, 3}}},
		{"2x2", args{NewLazyView(2, 2, func(row, col int) float64 { return float64(row + col) })}, [][]float64{{0, 1}, {1, 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Data(tt.args.M); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Data() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		M Matrix
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{nil}, "{}"},
		{"", args{Empty()}, "{}"},
		{"", args{Scalar(1)}, "{{1}}"},
		{"", args{RowVec()}, "{}"},
		{"", args{RowVec(-0.00000001)}, "{{-1e-08}}"},
		{"", args{RowVec(1, 2)}, "{{1, 2}}"},
		{"", args{RowVec([]float64{1, 2, 3}...)}, "{{1, 2, 3}}"},
		{"", args{ColVec()}, "{}"},
		{"", args{ColVec(-0.00000001)}, "{{-1e-08}}"},
		{"", args{ColVec(1, 2)}, "{{1}, {2}}"},
		{"", args{ColVec([]float64{1, 2, 3}...)}, "{{1}, {2}, {3}}"},
		{"", args{Dense{[][]float64{{1, 2}, {3, 4}}}}, "{{1, 2}, {3, 4}}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := String(tt.args.M); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranspose(t *testing.T) {
	type args struct {
		M Matrix
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{"nil", args{M: nil}, nil},
		{"empty", args{M: Empty()}, Empty()},
		{"scalar", args{M: Scalar(1)}, Scalar(1)},
		{"row vec", args{M: RowVec(1, 2, 3)}, ColVec(1, 2, 3)},
		{"col vec", args{M: ColVec(1, 2, 3)}, RowVec(1, 2, 3)},
		{"diag", args{M: Diag(1, 2, 3)}, Diag(1, 2, 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Transpose(tt.args.M)
			EXPECT_EQ(t, got, tt.want, float_eq_tol)
		})
	}
}

func TestIsSquare(t *testing.T) {
	type args struct {
		M Matrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"nil", args{nil}, false},
		{"empty", args{Empty()}, true},
		{"scalar", args{Scalar(1)}, true},
		{"row vec", args{RowVec(1, 2, 3)}, false},
		{"col vec", args{ColVec(1, 2, 3)}, false},
		{"row vec", args{RowVec(1)}, true},
		{"col vec", args{ColVec(1)}, true},
		{"2x2", args{Rand(2, 2)}, true},
		{"2x3", args{Rand(2, 3)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSquare(tt.args.M); got != tt.want {
				t.Errorf("IsSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}

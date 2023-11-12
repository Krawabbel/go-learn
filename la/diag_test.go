package la

import (
	"testing"
)

func TestDiag(t *testing.T) {
	type args struct {
		vals []float64
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{"nil", args{nil}, nil},
		{"success", args{[]float64{1, 2}}, Dense{[][]float64{{1, 0}, {0, 2}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Diag(tt.args.vals...)
			EXPECT_EQ(t, got, tt.want, float_eq_tol)
		})
	}
}

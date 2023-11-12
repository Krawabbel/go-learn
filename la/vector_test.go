package la

import (
	"testing"
)

func TestLinspace(t *testing.T) {
	type args struct {
		a float64
		b float64
		N int
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{"empty", args{0, 2, 0}, nil},
		{"single", args{0, 2, 1}, nil},
		{"default", args{0, 2, 3}, RowVec(0, 1, 2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Linspace(tt.args.a, tt.args.b, tt.args.N)
			EXPECT_EQ(t, got, tt.want, FLOAT_EQ_TOL)
		})
	}
}

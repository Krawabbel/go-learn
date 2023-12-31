package la

import (
	"testing"
)

func TestMult(t *testing.T) {
	type args struct {
		M Matrix
		N Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    Matrix
		wantErr bool
	}{
		{"nil", args{nil, Scalar(0)}, nil, true},
		{"empty * 1x1", args{Rand(0, 0), Scalar(0)}, nil, true},
		{"2x2 * 1x1", args{Rand(2, 2), Scalar(0)}, nil, true},
		{"empty", args{Empty(), Empty()}, Empty(), false},
		{"scalar", args{Scalar(1), Scalar(-1)}, Scalar(-1), false},
		{"mat-vec", args{Mat([][]float64{{1, 2}, {3, 4}}), ColVec(5, 6)}, ColVec(17, 39), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Mult(tt.args.M, tt.args.N)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mult() error = %v, wantErr %v", err, tt.wantErr)
			}
			EXPECT_EQ(t, got, tt.want, FLOAT_EQ_TOL)
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		M Matrix
		N Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    Matrix
		wantErr bool
	}{
		{"nil", args{nil, Scalar(0)}, nil, true},
		{"empty + 1x1", args{Empty(), Scalar(0)}, nil, true},
		{"2x2 + 1x1", args{Rand(2, 2), Scalar(0)}, nil, true},
		{"empty", args{Empty(), Empty()}, Empty(), false},
		{"scalar", args{Scalar(1), Scalar(-1)}, Scalar(0), false},
		{"col vec", args{ColVec(1, 2, 3), ColVec(4, 5, 6)}, ColVec(5, 7, 9), false},
		{"row vec", args{RowVec(1, 2, 3), RowVec(4, 5, 6)}, RowVec(5, 7, 9), false},
		{"dense", args{Mat([][]float64{{1, 2}, {3, 4}}), Mat([][]float64{{1, 2}, {3, 4}})}, Mat([][]float64{{2, 4}, {6, 8}}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.M, tt.args.N)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
			}
			EXPECT_EQ(t, got, tt.want, FLOAT_EQ_TOL)
		})
	}
}

func TestScale(t *testing.T) {
	type args struct {
		c float64
		M Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    Matrix
		wantErr bool
	}{
		{"nil", args{c: 0, M: nil}, nil, true},
		{"empty", args{c: 0, M: Empty()}, Empty(), false},
		{"zero", args{c: 0, M: Ones(2, 2)}, Zeros(2, 2), false},
		{"double", args{c: 2, M: Diag(1, 2, 3)}, Diag(2, 4, 6), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Scale(tt.args.c, tt.args.M)
			if (err != nil) != tt.wantErr {
				t.Errorf("Scale() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			EXPECT_EQ(t, got, tt.want, FLOAT_EQ_TOL)
		})
	}
}

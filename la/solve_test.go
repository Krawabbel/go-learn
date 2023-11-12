package la

import (
	"testing"
)

func Test_ls_solve(t *testing.T) {
	type args struct {
		A Matrix
		Y Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    Matrix
		wantErr bool
	}{
		{"nil nil", args{nil, nil}, nil, true},
		{"nil empty", args{nil, Empty()}, nil, true},
		{"empty nil", args{Empty(), nil}, nil, true},
		{"empty empty", args{Empty(), Empty()}, Empty(), false},
		{"scalar scalar", args{Scalar(-2), Scalar(4)}, Scalar(-2), false},
		{"vec vec", args{ColVec(1, 1), ColVec(2, 3)}, Scalar(2.5), false},
		{"linear", args{ColVec(1, 2, 3), ColVec(2, 4, 6)}, Scalar(2), false},
		{"interpolate 1+2x+3x^2 at x=(0,1,2)", args{Mat([][]float64{{1, 0, 0}, {1, 1, 1}, {1, 2, 4}}), ColVec(1, 6, 17)}, ColVec(1, 2, 3), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ls_solve(tt.args.A, tt.args.Y)
			if (err != nil) != tt.wantErr {
				t.Errorf("ls_solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			EXPECT_EQ(t, got, tt.want, FLOAT_EQ_TOL)
		})
	}
}

func Test_lower_solve(t *testing.T) {
	type args struct {
		L Matrix
		Y Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    Matrix
		wantErr bool
	}{
		{"nil nil", args{nil, nil}, nil, true},
		{"nil empty", args{nil, Empty()}, nil, true},
		{"empty nil", args{Empty(), nil}, nil, true},
		{"empty scalar", args{Empty(), Scalar(1)}, nil, true},
		{"scalar empty", args{Scalar(1), Empty()}, nil, true},
		{"empty empty", args{Empty(), Empty()}, Empty(), false},
		{"scalar scalar", args{Scalar(2), Scalar(6)}, Scalar(3), false},
		{"2x2 2x1", args{Diag(1, 2), ColVec(2, 4)}, ColVec(2, 2), false},
		{"2x1 2x2", args{ColVec(1, 2), Diag(2, 4)}, nil, true},
		{"2x2 2x2", args{Mat([][]float64{{1, 0}, {2, 3}}), Mat([][]float64{{4, 0}, {5, 3}})}, Mat([][]float64{{4, 0}, {-1, 1}}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lower_solve(tt.args.L, tt.args.Y)
			if (err != nil) != tt.wantErr {
				t.Errorf("lower_solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			EXPECT_EQ(t, got, tt.want, FLOAT_EQ_TOL)
		})
	}
}

func TestLower_Solve_2x2(t *testing.T) {
	Ls := []Matrix{
		Mat([][]float64{{1, 0}, {2, 3}}),
	}
	Ys := []Matrix{
		ColVec(4, 5),
		ColVec(6, 7),
		Mat([][]float64{{4, 5}, {6, 7}}),
	}
	for _, L := range Ls {
		for _, Y := range Ys {
			X, err := lower_solve(L, Y)
			if err != nil {
				t.Error(err)
			}
			Z, err := Mult(L, X)
			if err != nil {
				t.Error(err)
			}
			EXPECT_EQ(t, Y, Z, FLOAT_EQ_TOL)
		}
	}
}

func TestLower_Solve_Rand(t *testing.T) {
	for dim := 1; dim < 10; dim++ {
		for num := 1; num < 3; num++ {
			L := Lower(Rand(dim, dim))
			Y := Rand(dim, num)
			X, err := lower_solve(L, Y)
			if err != nil {
				t.Error(err)
			}
			Z, err := Mult(L, X)
			if err != nil {
				t.Error(err)
			}
			EXPECT_EQ(t, Y, Z, FLOAT_EQ_TOL)
		}
	}
}

func Test_upper_solve(t *testing.T) {
	type args struct {
		L Matrix
		Y Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    Matrix
		wantErr bool
	}{
		{"nil nil", args{nil, nil}, nil, true},
		{"nil empty", args{nil, Empty()}, nil, true},
		{"empty nil", args{Empty(), nil}, nil, true},
		{"empty scalar", args{Empty(), Scalar(1)}, nil, true},
		{"scalar empty", args{Scalar(1), Empty()}, nil, true},
		{"empty empty", args{Empty(), Empty()}, Empty(), false},
		{"scalar scalar", args{Scalar(2), Scalar(6)}, Scalar(3), false},
		{"2x2 2x1", args{Diag(1, 2), ColVec(2, 4)}, ColVec(2, 2), false},
		{"2x1 2x2", args{ColVec(1, 2), Diag(2, 4)}, nil, true},
		{"2x2 2x2", args{Mat([][]float64{{1, 1}, {0, 3}}), Mat([][]float64{{4, 0}, {5, 3}})}, Mat([][]float64{{7.0 / 3, -1}, {5.0 / 3, 1}}), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := upper_solve(tt.args.L, tt.args.Y)
			if (err != nil) != tt.wantErr {
				t.Errorf("upper_solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			EXPECT_EQ(t, got, tt.want, FLOAT_EQ_TOL)
		})
	}
}

func TestUpper_Solve_Rand(t *testing.T) {
	for dim := 1; dim < 10; dim++ {
		for num := 1; num < 3; num++ {
			L := Upper(Rand(dim, dim))
			Y := Rand(dim, num)
			X, err := upper_solve(L, Y)
			if err != nil {
				t.Error(err)
			}
			Z, err := Mult(L, X)
			if err != nil {
				t.Error(err)
			}
			EXPECT_EQ(t, Y, Z, FLOAT_EQ_TOL)
		}
	}
}

package matrix

import (
	"testing"
)

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
		{"2x2 2x2", args{Dense{[][]float64{{1, 0}, {2, 3}}}, Dense{[][]float64{{4, 0}, {5, 3}}}}, Dense{[][]float64{{4, 0}, {-1, 1}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lower_solve(tt.args.L, tt.args.Y)
			if (err != nil) != tt.wantErr {
				t.Errorf("lower_solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			EXPECT_EQ(t, got, tt.want, float_eq_tol)
		})
	}
}

func TestLower_Solve_2x2(t *testing.T) {
	Ls := []Dense{
		{[][]float64{{1, 0}, {2, 3}}},
	}
	Ys := []Matrix{
		ColVec(4, 5),
		ColVec(6, 7),
		Dense{[][]float64{{4, 5}, {6, 7}}},
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
			EXPECT_EQ(t, Y, Z, float_eq_tol)
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
			EXPECT_EQ(t, Y, Z, float_eq_tol)
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
		{"2x2 2x2", args{Dense{[][]float64{{1, 1}, {0, 3}}}, Dense{[][]float64{{4, 0}, {5, 3}}}}, Dense{[][]float64{{7.0 / 3, -1}, {5.0 / 3, 1}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := upper_solve(tt.args.L, tt.args.Y)
			if (err != nil) != tt.wantErr {
				t.Errorf("upper_solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			EXPECT_EQ(t, got, tt.want, float_eq_tol)
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
			EXPECT_EQ(t, Y, Z, float_eq_tol)
		}
	}
}

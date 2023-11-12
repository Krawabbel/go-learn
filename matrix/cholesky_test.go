package matrix

import (
	"math"
	"reflect"
	"testing"
)

func Test_chol_dec(t *testing.T) {
	type args struct {
		M Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    Matrix
		wantErr bool
	}{
		{"nil", args{M: nil}, nil, true},
		{"empty", args{M: Empty()}, Empty(), false},
		{"[1]", args{M: Scalar(1)}, Scalar(1), false},
		{"[4]", args{M: Scalar(4)}, Scalar(2), false},
		{"[-4]", args{M: Scalar(-4)}, nil, true},
		{"diag(4,9)", args{M: Diag(4, 9)}, Diag(2, 3), false},
		{"diag(-4,9)", args{M: Diag(-4, 9)}, nil, true},
		{"1x2", args{M: RowVec(1, 2)}, nil, true},
		{"[4,0;1,9]", args{M: Dense{[][]float64{{4, 0}, {1, 9}}}}, nil, true},
		{"[2,1;1,3]", args{M: Dense{[][]float64{{2, 1}, {1, 3}}}}, Dense{[][]float64{{math.Sqrt(2), 0}, {math.Sqrt(0.5), math.Sqrt(2.5)}}}, false},
		{"[4,12,-16;12,37,-43;-16,-43,98]", args{M: Dense{[][]float64{{4, 12, -16}, {12, 37, -43}, {-16, -43, 98}}}}, Dense{[][]float64{{2, 0, 0}, {6, 1, 0}, {-8, 5, 3}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			have, err := chol_dec(tt.args.M)
			if (err != nil) != tt.wantErr {
				t.Errorf("Cholesky() error = %v, wantErr %v", err, tt.wantErr)
			}
			EXPECT_EQ(t, have, tt.want, float_eq_tol)

			if err == nil {
				reverse, err := Mult(have, Transpose(have))
				if err != nil {
					t.Error(err)
				}
				EXPECT_EQ(t, reverse, tt.args.M, float_eq_tol)
			}
		})
	}
}

func Test_chol_solve(t *testing.T) {
	type args struct {
		A Matrix
		Z Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    Matrix
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := chol_solve(tt.args.A, tt.args.Z)
			if (err != nil) != tt.wantErr {
				t.Errorf("chol_solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chol_solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChol_Solve_Rand(t *testing.T) {
	for dim := 1; dim < 10; dim++ {
		for num := 1; num < 3; num++ {
			M := Rand(dim, dim)
			N, err := Mult(Transpose(M), M)
			if err != nil {
				t.Fatal(err)
			}
			A, err := Add(N, Eye(dim))
			if err != nil {
				t.Fatal(err)
			}
			Y := Rand(dim, num)
			X, err := chol_solve(A, Y)
			if err != nil {
				t.Fatal(err)
			}
			Z, err := Mult(A, X)
			if err != nil {
				t.Fatal(err)
			}
			EXPECT_EQ(t, Y, Z, float_eq_tol)
		}
	}
}

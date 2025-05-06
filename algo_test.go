package main

import "testing"

func generateBigMatrix() [][]int64 {
	m := make([][]int64, 100)
	for i := range len(m) {
		m[i] = make([]int64, 100)
	}

	for i := range len(m) {
		for j := range len(m[i]) {
			m[i][j] = 1000000000
		}
	}

	return m
}

func Test_doAlgorithm(t *testing.T) {
	type args struct {
		matrix [][]int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{matrix: [][]int64{
				{1, 2},
				{1, 2},
			}},
			want: false,
		},
		{
			name: "example 2",
			args: args{matrix: [][]int64{
				{10, 20, 30},
				{1, 1, 1},
				{0, 0, 1},
			}},
			want: false,
		},
		{
			name: "simple true",
			args: args{matrix: [][]int64{
				{0, 5, 5},
				{15, 0, 5},
				{15, 15, 0},
			}},
			want: true,
		},
		{
			name: "zero matrix",
			args: args{matrix: [][]int64{
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0},
			}},
			want: true,
		},
		{
			name: "zero container without zero color",
			args: args{matrix: [][]int64{
				{0, 0, 0},
				{1, 1, 1},
				{2, 2, 2},
			}},
			want: false,
		},
		{
			name: "zero container with zero color",
			args: args{matrix: [][]int64{
				{0, 0, 0},
				{0, 1, 2},
				{0, 2, 1},
			}},
			want: true,
		},
		{
			name: "the biggest matrix",
			args: args{matrix: generateBigMatrix()},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doAlgorithm(tt.args.matrix); got != tt.want {
				t.Errorf("doAlgorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import "testing"

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := doAlgorithm(tt.args.matrix); got != tt.want {
				t.Errorf("doAlgorithm() = %v, want %v", got, tt.want)
			}
		})
	}
}

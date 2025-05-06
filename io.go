package main

import (
	"fmt"
	"io"
)

const (
	YES string = "yes"
	NO  string = "no"
)

func readMatrix(reader io.Reader) ([][]int64, error) {
	var n int
	if _, err := fmt.Fscan(reader, &n); err != nil {
		return nil, fmt.Errorf("invalid data: %s", err.Error())
	}
	if n < 1 || n > 100 {
		return nil, fmt.Errorf("invalid data")
	}

	matrix := make([][]int64, n)
	for i := range matrix {
		matrix[i] = make([]int64, n)
	}

	for i := range n {
		for j := range n {
			if _, err := fmt.Fscan(reader, &matrix[i][j]); err != nil {
				return nil, fmt.Errorf("invalid data: %s", err.Error())
			}
		}
	}

	return matrix, nil
}

func printAnswer(answer bool) {
	if answer {
		fmt.Println(YES)
	} else {
		fmt.Println(NO)
	}
}

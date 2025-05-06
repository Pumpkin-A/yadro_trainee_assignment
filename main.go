package main

import (
	"os"
)

func main() {
	var answer bool

	defer func() {
		printAnswer(answer)
	}()

	matrix, err := readMatrix(os.Stdin)
	if err != nil {
		return
	}

	answer = doAlgorithm(matrix)
}

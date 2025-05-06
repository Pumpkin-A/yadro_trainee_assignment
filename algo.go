package main

import "slices"

func countItems(matrix [][]int64) ([]int64, []int64) {
	сontainersCapacity := make([]int64, len(matrix))
	colorsCapacity := make([]int64, len(matrix))
	for i, row := range matrix {
		for j, elem := range row {
			сontainersCapacity[i] += elem
			colorsCapacity[j] += elem
		}
	}

	return colorsCapacity, сontainersCapacity
}

func doAlgorithm(matrix [][]int64) bool {
	colorsItem, containersItem := countItems(matrix)

	slices.Sort(containersItem)
	slices.Sort(colorsItem)

	for i := range len(matrix) {
		if colorsItem[i] != containersItem[i] {
			return false
		}
	}

	return true
}

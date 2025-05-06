package main

import "slices"

type Item struct {
	row []int64
	sum int64
}

func countItems(matrix [][]int64) (Item, Item) {
	сontainersItem := Item{
		row: make([]int64, len(matrix)),
	}
	colorsItem := Item{
		row: make([]int64, len(matrix)),
	}
	for i, row := range matrix {
		for j, elem := range row {
			сontainersItem.row[i] += elem
			colorsItem.row[j] += elem

			сontainersItem.sum += elem
			colorsItem.sum += elem
		}
	}

	return colorsItem, сontainersItem
}

func doAlgorithm(matrix [][]int64) bool {
	colorsItem, containersItem := countItems(matrix)
	if colorsItem.sum != containersItem.sum {
		return false
	}

	slices.Sort(containersItem.row)
	slices.Sort(colorsItem.row)

	for i := range len(matrix) {
		if colorsItem.row[i] != containersItem.row[i] {
			return false
		}
	}

	return true
}

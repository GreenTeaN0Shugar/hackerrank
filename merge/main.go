package main

import "fmt"

func main() {
	s := []int{9, 2, 3, 1, 4, 7, 5, 6, 0, 8}
	res := make([]int, len(s))
	mergeSortRecursive(s, res, 0, len(s)-1)
}

func mergeSortRecursive(s, result []int, left, right int) {
	if left == right {
		return
	}

	middle := (left + right) / 2

	mergeSortRecursive(s, result, left, middle)
	mergeSortRecursive(s, result, middle+1, right)

	merge()
}

func merge(s, result, leftBegin, rightBegin, rightEnd): {
	fmt.Println("MERGE NOT IMPLEMENTED")
}

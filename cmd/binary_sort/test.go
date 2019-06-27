package main

import (
	"fmt"
)

func main() {
	slice := []int32{1, 4, 5, 12, 3, 343, 5, 12312, 444, 555, 11, 34, 55, 656, 888, 996, 12}

	sliceSort := binarySort(slice)
	fmt.Println(sliceSort)
}

func binarySort(slice []int32) []int32 {
	len := len(slice)
	if len == 1 || len == 0 {
		return slice
	}
	mid := len / 2
	midNumber := slice[mid]
	var left []int32
	var right []int32
	var middle []int32
	for i := 0; i < len; i++ {
		if slice[i] == midNumber {
			middle = append(middle, slice[i])
		} else if slice[i] < midNumber {
			left = append(left, slice[i])
		} else {
			right = append(right, slice[i])
		}

	}

	result := merge(middle, binarySort(left), binarySort(right))

	return result
}

func merge(middle []int32, left []int32, right []int32) []int32 {
	for m := 0; m < len(middle); m++ {
		left = append(left, middle[m])
	}
	for i := 0; i < len(right); i++ {
		left = append(left, right[i])
	}

	return left
}

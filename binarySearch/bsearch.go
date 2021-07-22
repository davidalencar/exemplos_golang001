package main

import "fmt"

func binarySerach(list []int, item int) (bool, int) {
	low := 0
	high := len(list) - 1

	for low <= high {
		median := (low + high) / 2

		if list[median] == item {
			return true, median
		} else if list[median] < item {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	return false, -1
}

func main() {
	list := []int{2, 3, 5, 7, 11, 13, 17, 19}
	exist, index := binarySerach(list, 7)

	fmt.Println(exist, index)
}

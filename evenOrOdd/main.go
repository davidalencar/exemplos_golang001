package main

import "fmt"

func createSliceOfInt(size int) []int {
	vals := make([]int, size)
	for i := range vals {
		vals[i] = i + 1
	}
	return vals
}

func checkEvenOrOdd(val int) string {
	if val%2 == 0 {
		return "even"
	}

	return "odd"
}

func main() {

	vals := createSliceOfInt(10)

	for _, v := range vals {
		fmt.Println(v, checkEvenOrOdd(v))
	}
}

package main

import (
	"fmt"
)

func sol(arr []int, req int) int {
	start := 0
	end := len(arr)

	for start < end {
		pivot := (start + end) / 2 // assume there is no overflow
		if req < arr[pivot] {
			end = pivot
		} else {
			start = pivot + 1
		}
	}
	return start
}

func main() {
	len := 0
	fmt.Scan(&len)

	arr := make([]int, 0, len)

	for i := 0; i < len; i++ {
		i := 0
		fmt.Scan(&i)
		arr = append(arr, i)
	}

	req := 0
	fmt.Scan(&req)
	fmt.Print(sol(arr, req))
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}

	sum := 0
	max_sum := 0
	temp := 0
	left := 0
	right := 0

	for i := 0; i < len(a); i++ {
		sum += a[i]

		if sum > max_sum {
			max_sum = sum
			left = temp
			right = i
		}

		if sum < 0 {
			sum = 0
			temp = i + 1
		}
	}

	fmt.Println("Maximum subarray: [" + strconv.Itoa(left) + ", " + strconv.Itoa(right) + "] with sum= " + strconv.Itoa(max_sum))
}

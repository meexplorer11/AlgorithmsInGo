package main

import (
	"fmt"
	"strconv"
)

type tuple struct {
	left  int
	right int
	sum   int
}

func main() {
	a := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}

	t := find_maximum_subarray(a, 0, len(a)-1)

	fmt.Println("Maximum subarray: [" + strconv.Itoa(t.left) + ", " + strconv.Itoa(t.right) + "] with sum= " + strconv.Itoa(t.sum))
}

func find_maximum_subarray(a []int, low int, high int) tuple {
	if low == high {
		return tuple{low, high, a[low]}
	} else {
		mid := low + (high-low)/2

		left_tuple := find_maximum_subarray(a, low, mid)

		righ_tuple := find_maximum_subarray(a, mid+1, high)

		cross_tuple := find_cross_maximum_subarray(a, low, mid, high)

		if left_tuple.sum >= righ_tuple.sum && left_tuple.sum >= cross_tuple.sum {
			return left_tuple
		} else if righ_tuple.sum >= left_tuple.sum && righ_tuple.sum >= cross_tuple.sum {
			return righ_tuple
		} else {
			return cross_tuple
		}
	}
}

func find_cross_maximum_subarray(a []int, low int, mid int, high int) tuple {
	sum := 0
	left_sum := 0
	left_idx := 0

	for i := mid; i >= low; i-- {
		sum += a[i]
		if sum > left_sum {
			left_sum = sum
			left_idx = i
		}
	}

	sum = 0
	right_sum := 0
	right_idx := 0

	for i := mid + 1; i < high; i++ {
		sum += a[i]
		if sum > right_sum {
			right_sum = sum
			right_idx = i
		}
	}

	return tuple{left_idx, right_idx, left_sum + right_sum}
}

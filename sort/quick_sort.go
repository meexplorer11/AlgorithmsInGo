package main

import "fmt"

func main() {
	a := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

	fmt.Println("Before sorting", a)

	sort(&a, 0, len(a)-1)

	fmt.Println("After sorting", a)
}

func sort(a *[]int, low int, high int) {
	if low < high {
		pivot := partition(a, low, high)

		sort(a, low, pivot-1)
		sort(a, pivot+1, high)
	}
}

func partition(a *[]int, low int, high int) int {
	pivot := (*a)[high]

	i := low - 1

	for j := low; j < high; j++ {
		if (*a)[j] < pivot {
			i++

			(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
		}
	}

	(*a)[i+1], (*a)[high] = (*a)[high], (*a)[i+1]

	return i + 1
}

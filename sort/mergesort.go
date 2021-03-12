package main

import "fmt"

func main() {
	a := []int{5, 2, 7, 1, 4, 3, 6, 2}

	fmt.Println("Before sort", a)

	sort(&a, 0, len(a)-1)

	fmt.Println("After sort", a)
}

func sort(a *[]int, start int, end int) {
	if start < end {
		mid := (start + end - 1) / 2

		sort(a, start, mid)
		sort(a, mid+1, end)
		merge(a, start, mid, end)
	}
}

func merge(a *[]int, start int, mid int, end int) {
	n1 := mid - start + 1
	n2 := end - mid

	left := make([]int, n1)
	right := make([]int, n2)

	for i := 0; i < n1; i++ {
		left[i] = (*a)[start+i]
	}

	for i := 0; i < n2; i++ {
		right[i] = (*a)[mid+1+i]
	}

	i := 0
	j := 0
	k := start

	for i < n1 && j < n2 {
		if left[i] < right[j] {
			(*a)[k] = left[i]
			i++
		} else {
			(*a)[k] = right[j]
			j++
		}
		k++
	}

	for i < n1 {
		(*a)[k] = left[i]
		i++
		k++
	}

	for j < n2 {
		(*a)[k] = right[j]
		j++
		k++
	}
}

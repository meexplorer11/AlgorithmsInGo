package main

import "fmt"

func main() {
	a := []int{5, 2, 4, 7, 1, 3, 2, 6}

	fmt.Println("Before sorting", a)

	sort(&a)

	fmt.Println("After sorting", a)
}

func sort(a *[]int) {
	for i := 0; i < len(*a); i++ {
		key := (*a)[i]
		j := i - 1
		for j >= 0 && (*a)[j] > key {
			(*a)[j+1] = (*a)[j]
			j = j - 1
		}
		(*a)[j+1] = key
	}
}

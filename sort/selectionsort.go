package main

import "fmt"

func main() {
	a := []int{5, 2, 7, 4, 1, 3, 6, 2}

	fmt.Println("Before sort", a)

	sort(&a)

	fmt.Println("After sort", a)
}

func sort(a *[]int) {
	for index := range *a {
		minKey := index

		for j := index + 1; j < len(*a); j++ {
			if (*a)[j] < (*a)[minKey] {
				minKey = j
			}
		}

		(*a)[index], (*a)[minKey] = (*a)[minKey], (*a)[index]
	}
}

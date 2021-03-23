package main

import "fmt"

func main() {
	a := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}

	fmt.Println("Before sorting", a)

	sort(a)
}

func sort(a []int) {
	buildMaxHeap(&a)

	for i := len(a) - 1; i >= 0; i-- {
		swap(&a, i, 0)
		maxHeapify(&a, i, 0)
	}

	fmt.Println("After sorting", a)
}

func buildMaxHeap(a *[]int) {
	heapSize := len(*a)

	for i := heapSize/2 - 1; i >= 0; i-- {
		maxHeapify(a, heapSize, i)
	}

}

func maxHeapify(a *[]int, heapSize int, i int) {
	largest := i
	left := left(i)
	right := right(i)

	if left < heapSize && (*a)[left] > (*a)[largest] {
		largest = left
	}

	if right < heapSize && (*a)[right] > (*a)[largest] {
		largest = right
	}

	if largest != i {
		swap(a, largest, i)
		maxHeapify(a, heapSize, largest)
	}
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func swap(a *[]int, i int, j int) {
	(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
}

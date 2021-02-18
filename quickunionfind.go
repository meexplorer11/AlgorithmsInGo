package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	size, _ := strconv.Atoi(os.Args[1])

	items := make([]int, size)

	for index := range items {
		items[index] = index
	}
	fmt.Println(items)

	logConnected(isConnected(&items, 2, 5))
	union(&items, 2, 5)
	logConnected(isConnected(&items, 2, 5))
}

func root(arr *[]int, p int) int {
	for p != (*arr)[p] {
		p = (*arr)[p]
	}
	fmt.Println("Root =" + strconv.Itoa(p))
	return p
}

func isConnected(arr *[]int, p int, q int) bool {
	return root(arr, p) == root(arr, q)
}

func logConnected(flag bool) {
	if flag {
		fmt.Println("Connected")
	} else {
		fmt.Println("Not Connected")
	}
}

func union(arr *[]int, p int, q int) {
	proot := root(arr, p)
	qroot := root(arr, q)

	(*arr)[qroot] = proot
}

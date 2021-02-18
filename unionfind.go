package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	size := os.Args[1]
	fmt.Println(size)

	s, _ := strconv.Atoi(size)
	arr := make([]int, s)
	fmt.Println(arr)

	for index := range arr {
		arr[index] = index
	}
	fmt.Println(arr)

	logConnected(connected(&arr, 2, 5))

	union(&arr, 2, 5)
	fmt.Println(arr)
	logConnected(connected(&arr, 2, 5))

	union(&arr, 2, 3)
	fmt.Println(arr)
	logConnected(connected(&arr, 2, 3))
}

func logConnected(flag bool) {
	if flag {
		fmt.Println("Connected")
	} else {
		fmt.Println("Not Connected")
	}
}

func connected(arr *[]int, p int, q int) bool {
	fmt.Println((*arr)[p], (*arr)[q])
	return (*arr)[p] == (*arr)[q]
}

func union(arr *[]int, p int, q int) {
	pVal := (*arr)[p]
	qVal := (*arr)[q]

	for index, value := range *arr {
		if value == pVal {
			(*arr)[index] = qVal
		}
	}
}

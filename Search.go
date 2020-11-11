package main

import (
	"fmt"
	"math"
)

func main() {
	xs := []int64{1, 2, 3, 4, 5} // Ascending order:
	pos := linearSearch(xs, 4)
	fmt.Println(pos)

	fmt.Println(binarySearch(xs, 1, 0, len(xs)-1))
	fmt.Println(binarySearch(xs, 2, 0, len(xs)-1))
	fmt.Println(binarySearch(xs, 3, 0, len(xs)-1))
	fmt.Println(binarySearch(xs, 4, 0, len(xs)-1))
	fmt.Println(binarySearch(xs, 5, 0, len(xs)-1))

	isSorted := isSortedAsc(xs)
	fmt.Println(isSorted)

	xs2 := []int64{1, 2, 4, 3, 5}
	isSorted2 := isSortedAsc(xs2)
	fmt.Println(isSorted2)

	xs3 := []int64{5, 4, 3, 2, 1}
	isSorted3 := isSortedDec(xs3)
	fmt.Println(isSorted3)

	fmt.Println(binarySearchDec(xs3, 1, 0, len(xs)-1))
	fmt.Println(binarySearchDec(xs3, 2, 0, len(xs)-1))
	fmt.Println(binarySearchDec(xs3, 3, 0, len(xs)-1))
	fmt.Println(binarySearchDec(xs3, 4, 0, len(xs)-1))
	fmt.Println(binarySearchDec(xs3, 5, 0, len(xs)-1))

	xs4 := []int64{5, 3, 4, 2, 1}
	isSorted4 := isSortedDec(xs4)
	fmt.Println(isSorted4)
}

func linearSearch(xs []int64, x int64) int {
	for i := 0; i < len(xs); i++ {
		if xs[i] == x {
			return i
		}
	}
	return -1
}

func isSortedAsc(xs []int64) bool {
	var prev int64 = -math.MaxInt64 - 1

	for i := 0; i < len(xs); i++ {
		if xs[i] < prev {
			return false
		}
		prev = xs[i]
	}
	return true
}

func isSortedDec(xs []int64) bool {
	var prev int64 = math.MaxInt64

	for i := 0; i < len(xs); i++ {
		if xs[i] > prev {
			return false
		}
		prev = xs[i]
	}
	return true
}

func binarySearch(xs []int64, x int64, l int, h int) int {
	if h < l {
		return -1
	}

	mid := (l + h) / 2

	if xs[mid] == x {
		return mid
	} else if x < xs[mid] {
		return binarySearch(xs, x, l, mid-1)
	} else {
		return binarySearch(xs, x, mid+1, h)
	}
}

func binarySearchDec(xs []int64, x int64, l int, h int) int {
	if h < l {
		return -1
	}

	mid := (l + h) / 2

	if xs[mid] == x {
		return mid
	} else if x > xs[mid] {
		return binarySearchDec(xs, x, l, mid-1)
	} else {
		return binarySearchDec(xs, x, mid+1, h)
	}
}

func funnyMsg() string {
	return "Funny"
}

func funnyMsg2() string {
	return "funny 2"
}

func funnyMsg3() string {
	return "funny 3333"
}

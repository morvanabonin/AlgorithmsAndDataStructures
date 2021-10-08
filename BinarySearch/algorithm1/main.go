package main

import (
	"fmt"
	"log"
)

// Binary-search

// Input: Sorted array a[1..n], a[1] < a[2] < · · · < a[n], element x
// Output: { m if there is an 1 ≤ m ≤ n with a[m] = x
//         { −1 otherwise
var (
	sortedSlice = []int{2, 3, 5, 7, 11, 13, 16, 20, 21, 27}
	x           = 13
	left        int
	right       int
)

func main() {
	ret := BinarySearch(sortedSlice, x)
	fmt.Println(ret)
}

func BinarySearch(ss []int, x int) int {
	left = -1
	right = len(ss) + 1
	for left+1 < right { /* 0 ≤ left < right ≤ n + 1 and a[left] < x < a[right] */
		middle := (left + right) / 2

		if !found(ss, x) {
			log.Fatal("index not found")
		}

		if ss[middle] == x {
			return middle
		} else if ss[middle] < x {
			left = middle
		} else {
			right = middle
		}
	}

	return -1
}

func found(ss []int, x int) bool  {
	for _, v := range ss {
		if v == x {
			return true
		}
	}
	return false
}

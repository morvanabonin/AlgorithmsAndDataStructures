package main

import (
	"fmt"
)

// Binary-search (recursive)

// Input: Sorted array a[1..n], a[1] < a[2] < · · · < a[n], element x
// Output: { m if there is an 1 ≤ m ≤ n with a[m] = x
//         { −1 otherwise

// Algorithm based on article
// https://levelup.gitconnected.com/binary-search-with-go-727b1943fd64
// https://brilliant.org/wiki/binary-search/
var sortedSlice = []int{2, 3, 5, 7, 11, 13, 16, 20, 21, 27}
var x = 27
var start = 0
var end = len(sortedSlice)

func main() {
	ret := binarySearch(sortedSlice, x, start, end)
	fmt.Println(ret)
}

func binarySearch(ss []int, x, s, e int) int {

	// checking if s or e are lower than 0, case true
	//  return -1
	if e < 0 || s < 0 {
		return -1
	}

	// checking the slice itself, if the slice has just one element, check if the element is the one
	if e - s == 0 {
		if x == ss[e] {
			return end
		}
		return -1
	}

	// create de pivot, this simply and tricky approach helps to divide the slice and two subslice
	// if the number is equal to the pivot = return true
	// if the number is smaller to the pivot = start searching in the left subslice
	// if the number is bigger to the pivot = start searching in the right subslice
	pivot := (e + s) /2

	if x == ss[pivot] {
		return pivot
	} else if x > ss[pivot] {
		return binarySearch(ss, x, pivot+1, e)
	}
	return binarySearch(ss, x, s, pivot)
}




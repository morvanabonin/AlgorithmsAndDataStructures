package main

import "fmt"

// Selection-sort
// Input: array a[1..n]
// Output: afterwards, a[1] ≤ a[2] ≤ · · · ≤ a[n] holds

// 0, 4, 7, 10, 14, 23, 45, 47, 53
var unsortedSlice = []int{14, 10, 23, 7, 45, 4, 47, 0, 53}

func main()  {
	min := GetMinimum(unsortedSlice)
	ret := SelectionSort(unsortedSlice, min)
	fmt.Println(ret)
}

func GetMinimum(us []int) int {
	var (
		j int
		i int
		aux = len(us)
	)

	for i = 0; i < len(us); i++ {
		for j = 1; j < len(us); j++  {
			if us[i] < us[j] {
				if aux > us[i] {
					aux = us[i]
				}
			}
		}
	}
	return aux
}

func SelectionSort(us []int, min int) []int {
	var s []int

	s = append(s, min)

	for i := 0; i < len(us); i++ {
		for j := 1; j < len(us); j++  {
			if min < us[i] && us[i] > us[j] {
				min = us[i]
				fmt.Println(min)
			}
		}
	}

	return s
}
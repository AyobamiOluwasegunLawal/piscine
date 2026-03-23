package main

import (
	"fmt"
)

func ConcatAlternate(slice1, slice2 []int) []int {
	newSlice := []int{}
	startWithFirst := true
	maxLen := len(slice1)

	if len(slice2) > len(slice1) {
		startWithFirst  = false
		maxLen = len(slice2)
	}

	for i := 0; i < maxLen; i++ {
		if startWithFirst {
			if i < len(slice1){
				newSlice = append(newSlice, slice1[i])
			}
			if i < len(slice2){
				newSlice = append(newSlice, slice2[i])
			}
		} else {
			if i < len(slice2){
				newSlice = append(newSlice, slice2[i])
			}

			if i < len(slice1) {
				newSlice = append(newSlice, slice1[i])
			}
		}
	}

	return newSlice
	
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
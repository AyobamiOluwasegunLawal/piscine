package main
import (
	"fmt"
)

func ConcatSlice(slice1, slice2 []int) []int {
	arr := []int{}
	i := 0
	for _, n := range slice1{
		if !check(arr, n) {
			arr = append(arr, n)
			i++
		}
	}

	for _, n := range slice2 {
		if !check(arr, n) {
			arr = append(arr, n)
			i++
		}
	}

	return arr
}

func check(s []int, c int) bool {
	for _, n := range s {
		if n == c {
			return true
		}
	}

	return false
} 

func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))
}
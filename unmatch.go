package main

func Unmatch(a []int) int {
	result := 0

	for _, v := range a {
		result ^= v
	}

	if result == 0 {
		return -1
	}

	return result
}
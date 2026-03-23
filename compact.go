package main
import "fmt"

func Compact(ptr *[]string) int{
	result := []string{}

	for _, v := range *ptr {
		if v != "" {
			result = append(result, v)
		}
	}

	*ptr = result
	return len(result)

}



const N = 6

func main() {
	a := make([]string, N)
	a[0] = "a"
	a[2] = "b"
	a[4] = "c"

	for _, v := range a {
		fmt.Println(v)
	}

	fmt.Println("Size after compacting:", Compact(&a))

	for _, v := range a {
		fmt.Println(v)
	}
}
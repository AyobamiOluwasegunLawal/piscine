package main
import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	if len(args) != 3{
		return
	}

	s1 := args[1]
	s2 := args[2]

	result := ""

	for _, c := range s1{
		if contains(s2, c) && !contains(result, c){
			z01.PrintRune(c)
			result += string(c)
		}
	}
	z01.PrintRune('\n')
}

func contains(s string, r rune) bool{
	for _, c := range s {
		if c == r {
			return true
		}
	}

	return false
}
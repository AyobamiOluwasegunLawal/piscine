package main
import (
	"os"
	"github.com/01-edu/z01"
)

func main () {
	if os.Args != 3 {
		z01.PrintRune('\n')
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	seen := ""

	for _, c := range s1 {
		if !contains(seen, c){
			z01.PrintRune(c)
			seen += string(c)
		}
	}

	for _, c := range s2 {
		if !contains(seen, c) {
			z01.PrintRune(c)
			seen += string(c)
		}
 	}
}

func contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}

	return false
}
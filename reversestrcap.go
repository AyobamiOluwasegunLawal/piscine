package main 
import (
	"os"
	"github.com/01-edu/z01"
)


func main () {
	args := os.Args
	if len(args) < 2 {
		return
	}

	for i := 1; i < len(args); i++ {
		s := args[1]

		for j := 0; j < len(s); j++{
			c := s[j]

			if c >= 'A' && c <= 'Z' {
				c += 32
			}

			if (j == len(s)-1 || s[j+1] == ' ') && (c >= 'a' && c <= 'z'){
				c -= 32
			}

			z01.PrintRune(rune(c))
		}

		z01.PrintRune('\n')
	}
}
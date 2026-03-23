package main
import (
	"os"
	"github.com/01-edu/z01"
)
func main () {
	args := os.Args

	if len(args) != 2 {
		return
	}

	n := atoi(args[1])

	if n <= 1 {
		return
	}

	first := true

	for i := 2; i <= n; i++ {
		for n%i == 0 {
			if !first{
				z01.PrintRune('*')
			}
			printNbr(i)

			n /= i
			first = false
		}
	}
	z01.PrintRune('\n')
}

func atoi(s string) int {
	n := 0

	for _, c := range s {
		if c < '0' || c > '9'{
			return 0
		}

		n = n * 10 + int(c - '0')
	}

	return n
}

func printNbr(n int) {
	if n >= 10 {
		printNbr(n/10)
	}

	z01.PrintRune(rune(n%10 + '0'))
}
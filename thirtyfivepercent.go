package main 
import (
	"github.com/01-edu/z01"
)

func isPrime(n int) bool {
	for i:= 2; i * i <= n; i++ {
		if n % i == 0 {
			return false
		}
	}

	return true
}

func FindPrevPrime(nb int) int {
	for nb >= 2 {
		if isPrime(nb) {
			return nb
		}

		nb--
	}

	return 0
}

func FromTo(from int, to int) string{
	if from > 99 || to > 99 || from < 0 || to < 0 {
		return "Invalid\n"
	}

	result := ""

	step := 1
	if from > to {
		step = -1
	}

	for i := from; ; i+=step {
		result += string('0' + rune(i/10))
		result += string('0' + rune(i%10))

		if i == to {
			break
		}

		result += ", "
	}

	return result + "\n"
}

func IsCapitalized(s string) bool{
	if s == "" {
		return false 
	}
	
	for i, _ := range s {
		if i == 0 || s[i-1] == ' '{
			if s[i] >= 'a' && s[i] <= 'z' {
				return false
			}
		}
	}
	return true
}

func Itoa(n int) string{
	if n == 0 {
		return "0"
	}
	
	result := ""
	if n < 0 {
		result += "-"
		n = -n
	}

	for n > 0 {
		result = string('0' + n % 10) + result
		n /= 10
	}

	return result
}

func PrintMemory(arr [10]byte) {
	for i:=0; i <len(arr); i++ {
		//convert byte to hex
		high := arr[i] / 16
		low := arr[i] % 16

		PrintHex(high)
		PrintHex(low)

		if i != len(arr)-1 && (i+1)%4 != 0 {
			z01.PrintRune(' ')
		}

		if (i+1)%4 == 0 || i == len(arr)-1 {
			z01.PrintRune('\n')
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			z01.PrintRune(rune(arr[i]))
		} else{
			z01.PrintRune('.')
		}
	}

	z01.PrintRune('\n')
}

func PrintHex(n byte) {
	if n < 10 {
		z01.PrintRune(rune('0' + n))
	} else{
		z01.PrintRune(rune('a' + (n - 10)))
	}
}
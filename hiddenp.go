package main
import (
	"os"
)


func main () {
	args := os.Args

	if len(args) != 3 {
		return
	}

	s1:= args[1]
	s2:= args[2]

	if len(s1) == 0 {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}

	j := 0

	for i:=0;, i<len(s2) && j < len(s1); i++ {
		if s2[i] == s1[j]{
			j++
		}
	}

	if j == len(s1){
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}

	z01.PrintRune('\n')
}
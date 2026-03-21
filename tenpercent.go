package main

func CheckNumber(arg string) bool {
	for _, r := range arg {
		if r >= '0' && r <= '9'{
			return true
		}
	}

	return false
}

func CountAlpha (s string) int {
	count := 0
	for _, r := range s {
		if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z' ) {
			continue
		} else {
			count++
		}
	}

	return count
}

func CountChar(str string, c rune) int {
	count := 0
	for _, r := range str {
		if r == c {
			count++
		} else {
			continue
		}
 	}

	return count
}

func PrintIf(str string) string {
	// if str == "" {
	// 	return "G\n"
	// }

	// if len(str) >= 3 {
	// 	return "G\n"
	// }

	if str == "" || len(str) >= 3{
		return "G\n"
	}

	return "Invalid Input\n"
}

func PrintIfNot(str string) string{
	if len(str) < 3 || str == "" {
		return "G\n"
	} else{
		return "Invalid Input\n"
	}
}


func RectPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}

	return 2 * (w + h)
}

func RetainFirstHalf(str string) string{
	if str == "" {
		return ""
	}

	if len(str) == 1 {
		return str
	}

	half := len(str) / 2
	return str[:half]
}
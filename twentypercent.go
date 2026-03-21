package main

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	for i, r := range s {
		if !(r >= 'A' && r <= 'Z'|| r >= 'a' && r <= 'z'){
			return s
		}

		if i > 0 && r >= 'A' && r <= 'Z' && rune(s[i-1]) >= 'A' && rune(s[i-1]) <= 'Z'{
			return s
		}
	}

	if s[len(s) -1] >= 'A' && s[len(s) - 1] <= 'Z' {
		return s
	}

	result := ""

	for i, r := range s {
		if r >= 'A' && r <= 'Z'{
			if i != 0 {
				result += "_"
			}

			result += string(r + 32)
		}else {
			result += string(r)
		}
	}
	
	return result

}

func DigitLen(n, base int) int {
	count := 0

	if !(base >= 2 && base <= 36){
		return -1
	}

	if n < 0 {
		n = -n
	}

	for n != 0{
		count++
		n /= base
	}

	return count
}

func FirstWord(s string) string{
	end := 0
	for _, r := range s {
		if r == ' ' {
			break
		}

		end++
	}

	return s[:end] + "\n"
}

func FishAndChips(n int) string {
	if n < 0{
		return "error: number is negative"
	} else if n % 2 == 0 &&  n % 3 == 0 {
		return "fish and chips"
	} else if n % 2 == 0 {
		return "fish"
	} else if n % 3 == 0 {
		return "chips"
	} else {
		return "error: non divisible"
	}
}

func Gcd(a, b uint) uint{
	if a == 0 || b == 0 {
		return 0
	}

	for b != 0 {
		a, b = b, a % b
	}

	return a
}

func HashCode(dec string) string{
	str := ""
	// for _, r := range dec {
	// 	hash := (int(r) + len(dec)) % 127

	// 	if rune(hash) < '!'{
	// 		hash += 33
	// 	}

	// 	str += string(rune(hash))
	// }

	for i:=0; i < len(dec); i++ {
		hash := (int(dec[i]) + len(dec)) % 127

		if byte(hash) < 32 {
			hash += 33
		}

		str += string(hash)
	}

	return str
}

func LastWord(s string) string{
	i := len(s) -1

	for i >= 0 && s[i] == ' '{
		i--
	}

	if i < 0{
		return "\n"
	}

	end := i

	for i >= 0 && s[i] != ' '{
		i--
	}
 
	return s[i+1:end+1] + "\n"
}

func RepeatAlpha(s string)string{
	result := ""
	for _, c := range s{
		if c >= 'a' && c <= 'z' {
			n := int(c - 'a') + 1

			for i := 0; i < n; i++ {
				result += string(c)
			}
		} else if c >= 'A' && c <= 'Z' {
			n:= int(c - 'A') + 1

			for i:=0; i < n; i++ {
				result += string(c)
			} 
		}else{
				result += string(c)
		}
	}
	return result
}
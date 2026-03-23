package main

func ROt14(s string) string{
	result := ""
	for _, c := range s {
		if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'c'{
			result = result + string(c + 14)

			if c > 'Z' || c > 'z' {
				result = result + string(c+14 -26)
			}
		}
	}

	return result
}
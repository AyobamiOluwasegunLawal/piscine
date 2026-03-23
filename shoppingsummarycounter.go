package main
import "fmt"

func ShoppingSummaryCounter(str string) map[string]int {
	result := make(map[string]int)
	word := ""

	for i, _ := range str {
		if str[i] != ' '{
			word += string(str[i])
		} else{
			if word != ""{
				result[word]++
				word = ""
			}
		}
	}

	if word != ""{
		result[word]++
	}

	return result
}

func main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	for index, element := range ShoppingSummaryCounter(summary) {
		fmt.Println(index, "=>", element)
	}
}
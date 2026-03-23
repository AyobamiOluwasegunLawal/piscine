package main
import "fmt"

func DealAPackOfCards (deck []int){
	cardsPerPlayer := len(deck) / 4
	for i:= 0; i < 4; i++ {
		fmt.Print("Player ")
		fmt.Print(i + 1)
		fmt.Print(": ")

		start := i * cardsPerPlayer
		end := start + cardsPerPlayer


		for j := start; j < end; j++ {
			fmt.Print(deck[j])
			if j != end - 1{
				fmt.Print(", ")
			}
		}

		fmt.Print("\n")
	}
	
}

func main() {
	deck := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	DealAPackOfCards(deck)
}
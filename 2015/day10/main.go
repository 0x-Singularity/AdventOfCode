package main

import (
	"fmt"
)

func calculate(input string) string {
	var result string
	count := 1
	for i := 0; i < len(input); i++ {
		// checks if i+1 is within the bounds of input and if they are the same character
		if i+1 < len(input) && input[i] == input[i+1] {
			count++
		} else {
			result += fmt.Sprintf("%d%c", count, input[i])
			count = 1
		}
	}
	return result
}

func main() {
	input := "1321131112"
	for i := 0; i < 50; i++ {
		fmt.Println("Run ", i)
		input = calculate(input)
	}
	fmt.Println(len(input))
}

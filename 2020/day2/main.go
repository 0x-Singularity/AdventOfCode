package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the file", err)
	}
	defer file.Close()

	validPasswords := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		character := parts[1][0]
		numberPositions := strings.Split(parts[0], "-")
		firstNumberPosition, _ := strconv.Atoi(numberPositions[0])
		firstNumberPosition = firstNumberPosition - 1
		secondNumberPosition, _ := strconv.Atoi(numberPositions[1])
		secondNumberPosition = secondNumberPosition - 1

		passwordToSearch := parts[2]

		if (passwordToSearch[firstNumberPosition] == character) != (passwordToSearch[secondNumberPosition] == character) {
			validPasswords++
		}
	}
	fmt.Println(validPasswords)
}

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
		numberOfInstances := strings.Split(parts[0], "-")
		minNumberOfInstances, _ := strconv.Atoi(numberOfInstances[0])
		maxNumberOfInstances, _ := strconv.Atoi(numberOfInstances[1])

		passwordToSearch := parts[2]
		numberOfOccurences := 0

		for i := 0; i < len(passwordToSearch); i++ {
			if passwordToSearch[i] == character {
				numberOfOccurences++
			}
		}
		if numberOfOccurences >= minNumberOfInstances && numberOfOccurences <= maxNumberOfInstances {
			validPasswords++
		}
		fmt.Println(validPasswords)
	}
}

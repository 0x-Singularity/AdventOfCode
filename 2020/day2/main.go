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

	validPasswords := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		passwordSlice := strings.Split(line, " ")

		minNumberOfInstances := 0
		maxNumberOfInstances := 0
		character := passwordSlice[1]
		characterToSearchFor := strings.TrimRight(character, ":")[0]

		numberOfInstances := strings.Split(passwordSlice[0], "-")

		minNumberOfInstances, err := strconv.Atoi(numberOfInstances[0])
		if err != nil {
			fmt.Println("Error converting minNumberOfInstances to int")
		}

		maxNumberOfInstances, err2 := strconv.Atoi(numberOfInstances[1])
		if err2 != nil {
			fmt.Println("Error convering maxNumberOfInstances to int")
		}

		passwordToSearch := passwordSlice[2]
		numberOfOccurences := 0

		for i := 0; i < len(passwordToSearch); i++ {
			if passwordToSearch[i] == characterToSearchFor {
				numberOfOccurences++
			}
		}
		if numberOfOccurences >= minNumberOfInstances && numberOfOccurences <= maxNumberOfInstances {
			validPasswords++
		}

		fmt.Println(validPasswords)

	}
}

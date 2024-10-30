package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the file", err)
	}

	scanner := bufio.NewScanner(file)

	target := 2020
	lines := []int{}

	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error scanning the information from the file", err)
		}
		lines = append(lines, line)
	}

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if lines[i]+lines[j] == target {
				fmt.Println(lines[i] * lines[j])
			}
		}
	}
}

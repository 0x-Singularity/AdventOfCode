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
		fmt.Println("Error opening file", err)
	}
	scanner := bufio.NewScanner(file)

	depthsSlice := []int{}
	depthIncreaseNumber := 0

	for scanner.Scan() {
		line := scanner.Text()
		depth, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error converting the depths to integers", err)
			continue
		}
		depthsSlice = append(depthsSlice, depth)
	}

	for i := 1; i < len(depthsSlice); i++ {
		if depthsSlice[i] > depthsSlice[i-1] {
			depthIncreaseNumber++
		}
	}
	fmt.Println(depthIncreaseNumber)
}

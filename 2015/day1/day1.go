package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading the data", err)
	}

	content := string(data)

	calculateBasementEntrance(content)
}

func calculateBasementEntrance(content string) {
	position, currentFloor := 0, 0

	for _, char := range content {
		switch char {
		case '(':
			currentFloor++
		case ')':
			currentFloor--
		}
		position++
		if currentFloor == -1 {
			fmt.Println(position)
			break
		}
	}
}

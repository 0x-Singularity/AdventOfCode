package main

import (
	"fmt"
	"io"
	"os"
)
//comment to test git
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
		fmt.Println("Error reading file: ", err)
		return
	}

	floorLevel, upAFloor, downAFloor, position := 0, 0, 0, 0

	content := string(data)

	for _, char := range content {
		switch char {
		case '(':
			upAFloor++
		case ')':
			downAFloor++
		}
		floorLevel = upAFloor - downAFloor
		position++

		if floorLevel == -1 {
			fmt.Println(position)
			break
		}
	}

}

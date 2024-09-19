package main

import (
	"fmt"
	"io"
	"os"
)
<<<<<<< HEAD
//comment to test git
=======

// comment
>>>>>>> c8265810a55a436e485988d07d166f83fc585a57
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

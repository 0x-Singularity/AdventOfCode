package main

import (
  "fmt"
  "os"
  "io"
)

<<<<<<< HEAD
func main() {
	//Open the file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return
	}
	defer file.Close()
=======
func main () {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error opening the file", err)
  }
>>>>>>> 3c09dc343d05e50b596a5ce5b2fab37c473fc588

  data, err := io.ReadAll(file)
  if err != nil {
    fmt.Println("Error reading the data", err)
  }

  content := string(data)

  calculateBasementEntrance(content)
}

func calculateBasementEntrance (content string) {
  position, currentFloor := 0,0

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
      break;
    }
  } 
}

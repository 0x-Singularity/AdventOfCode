package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  file,err := os.Open("input.txt")
  if err != nil{
    fmt.Println("Error opening the file", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)

  currentSum := 0
  maxCalories := 0

  for scanner.Scan() {
    line := scanner.Text()
    if line == ""{
      if currentSum > maxCalories {
        maxCalories = currentSum
      }
      currentSum = 0
    } else {
      calories, err := strconv.Atoi(line)
      if err != nil{
        fmt.Println("Error converting data to integers", err)
      }
      currentSum += calories
    }
  }

  if currentSum > maxCalories {
    maxCalories = currentSum
  }
  fmt.Printf("The elf carrying the most calories was carrying %d calories.\n", maxCalories)
}



// read the data line by line, if there's an empty line that is the end for that elf
// if I run into an empty line, I should sum the total of the previous lines, and add
// that total to a slice

// At the end I should iterate through the slice and print the index + 1 with the highest sum

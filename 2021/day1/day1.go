package main

import (
  "fmt"
  "os"
  "strconv"
  "bufio"
)

func main () {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error opening file", err)
  }

  scanner := bufio.NewScanner(file)

  depthSlice := []int {}


  for scanner.Scan(){
    line := scanner.Text()
    depth,err := strconv.Atoi(line)
    if err != nil {
      fmt.Println("Error converting data to integers", err)
    }
      depthSlice = append(depthSlice, depth)
  }

  fmt.Println(calculateDepthIncreaseNumber(depthSlice))
}

func calculateDepthIncreaseNumber(depthSlice []int) int{
  numberOfIncrease := 0
  for i := 1; i < len(depthSlice); i++ {
    if depthSlice[i] > depthSlice[i - 1]{
      numberOfIncrease++
    }
  } 
  return numberOfIncrease
}


package main

import (
 "os"
 "bufio"
)
func turnOn()

func main() {
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error opening file:", err)
  }

  defer file.Close()

  var input []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    input = append(input,scanner.Text())
  }

  var lightGrid [1000][1000]bool

  
}

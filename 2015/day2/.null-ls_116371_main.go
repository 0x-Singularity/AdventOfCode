package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {

  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Print("Error opening file", err)
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan(){
    line := scanner.Text()
    fmt.Println(line)
  }
}

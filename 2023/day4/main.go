package main

import (
  "fmt"
  "os"
  "bufio"
)


func main () {

  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error opening the file", err)
  }

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    //line := scanner.Text()

  }


}

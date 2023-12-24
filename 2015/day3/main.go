package main

import (
  "fmt"
  "os"
  "io"
)

func main () {
  file, err := os.OpenFile("input.txt");
  if err != nil {
    fmt.Println("Error opening file", err)
    return
  }
  


}



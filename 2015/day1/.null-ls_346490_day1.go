package main

import (
  "fmt"
  "os"
  "bufio"
)


func main () {
  //Open the file
  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("Error opening file: ", err)
    return
  }
  defer file.Close()
  
  data,err := ioutil.ReadAll(file)
  if err != nil {
    fmt.Println("Error reading file: ", err)
    return
  }

  fmt.Println(string(data))






}

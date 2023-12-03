package main

import (
	"bufio"
	"fmt"
	"os"
  "log"
)

// find games that would be possible with 12 red cubes, 13 green cubes, and 14 blue cubes.
func main () {

  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  
  var sumOfIDs int
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
  }

  
}

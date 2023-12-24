package main

import (
  "fmt"
  "os"
  "io"
)

type Position struct {
  x,y int
}

func main () {
  file, err := os.Open("input.txt");
  if err != nil {
    fmt.Println("Error opening file", err)
    return
  }
  
  directions, err := io.ReadAll(file)
  if err != nil {
    fmt.Println("Error reading file", err)
    return
  }

  visitedHouses := make(map[string]bool)
  santaPosition := Position{0,0}
  visitedHouses[santaPosition.toKey()] = true
  
  for _, direction := range directions {
    switch direction {
    case '^':
      santaPosition.y++
    case 'v':
      santaPosition.y--
    case '>':
      santaPosition.x++
    case '<':
      santaPosition.x--
    }
    visitedHouses[santaPosition.toKey()] = true
  }

  numberofHouses := len(visitedHouses)
  fmt.Printf("Number of houses that were visited at least once %v", numberofHouses)
  
}

func (p Position) toKey() string {
  return fmt.Sprintf("%d,%d", p.x, p.y)
}



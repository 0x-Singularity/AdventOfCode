package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findSurfaceArea(length, width, height float64) float64 {
  surfaceArea := 2*length*width + 2*width*height + 2*height*length
  return surfaceArea
}

func findSmallestArea(length, width, height float64) float64 {
  area1 := length * width
  area2 := width * height
  area3 := height * length

  smallestArea := area1
  if area2 < smallestArea{
    smallestArea = area2
  }

  if area3 < smallestArea{
    smallestArea = area3
  }

  return smallestArea
}


func main() {

  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Print("Error opening file", err)
  }

  total :=  0.0
  scanner := bufio.NewScanner(file)
  for scanner.Scan(){
    line := scanner.Text()
    seperated := strings.Split(line, "x")

    length,err := strconv.ParseFloat(seperated[0], 64)
    if err != nil {
      fmt.Println("Error parsing length", err)
    }
    width, err := strconv.ParseFloat(seperated[1], 64)
    if err != nil{
      fmt.Println("Error parsing width", err)
    }
    height, err := strconv.ParseFloat(seperated[2], 64)
    if err != nil{
      fmt.Println("Error parsing the height", err)
    }

    surfaceArea := findSurfaceArea(length, width, height)

    smallestArea := findSmallestArea(length, width, height)
   
    total += surfaceArea + smallestArea
  }
  fmt.Printf("total %f", total)
}

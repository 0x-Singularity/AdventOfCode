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

func findSmallestPerimeter(length, width, height float64) float64{
  perimeter1 := 2 * (length + width)
  perimeter2 := 2 * (width + height)
  perimeter3 := 2 * (height + length)

  smallestPerimeter := perimeter1
  if perimeter2 < smallestPerimeter{
    smallestPerimeter = perimeter2
  }
  if perimeter3 < smallestPerimeter {
    smallestPerimeter = perimeter3
  }
  return smallestPerimeter
}

func findVolume(length, width, height float64) float64 {
  return length * width * height
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

    // Needed for part 1 gold star
    //surfaceArea := findSurfaceArea(length, width, height)
    //smallestArea := findSmallestArea(length, width, height)
    //total += surfaceArea + smallestArea

    smallestPerimeter := findSmallestPerimeter(length, width, height)
    volume := findVolume(length, width, height)
    total += smallestPerimeter + volume
    
  }
  fmt.Printf("total %f", total)
}

package main

import (
  "fmt"
  "bufio"
  "unicode"
  "strconv"
  "log"
  "os"
)

/*
As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?
*/

func main() {
   sum:=0

  // open the file using the Open() function from the os library
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  
  // read the file line by line using a scanner
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    var numberStr string

    // Filter out non-numeric characters
    for _, char := range line {
      if unicode.IsDigit(char) {
        numberStr += string(char)
      }
    }

    //Convert the string to a Integer
    if numberStr != "" {
      number, err := strconv.Atoi(numberStr)
      if err != nil {
        fmt.PrintLn("Error converting string to number", err)
        continue
      }
      fmt.PrintLn("Number found: ", number)
    }
    
  }

  if err := scanner.Err(); err != nil {
    fmt.PrintLn("Error reading from file", err)
  }


  file.Close()


}

package main

import (
  "fmt"
  "bufio"
  "unicode"
  "strconv"
  "log"
  "os"
  "strings"
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

  // open the file using the Open() function from the os library
  file, err := os.Open("input.txt")
  if err != nil {
    log.Fatal(err)
  }
  
  var totalSum int
  // read the file line by line using a scanner
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    firstDigit := ""
    lastDigit := ""

    // Filter out non-numeric characters and find first and last digits
    for _, char := range line {
      if unicode.IsDigit(char) {
        if firstDigit == "" {
          firstDigit = string(char)
        }
        lastDigit = string(char)
      }
    }
    // Concatenate the digits and convert to integer
        if firstDigit != "" && lastDigit != "" {
            concatenatedDigits := firstDigit + lastDigit
            digitInt, err := strconv.Atoi(concatenatedDigits)
            if err != nil {
                fmt.Println("Error converting concatenated digits to integer:", err)
                continue
            }

            totalSum += digitInt
            fmt.Printf("Concatenated digits for line '%s': %s\n", line, concatenatedDigits)
        }
    }
  

if err := scanner.Err(); err != nil {
        fmt.Println("Error reading from file:", err)
    }

    fmt.Println("Total sum:", totalSum)
  
}

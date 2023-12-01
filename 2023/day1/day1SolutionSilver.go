package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    // Mapping of spelled-out digits to their numerical counterparts
    digitMap := map[string]string{
        "one": "1", "two": "2", "three": "3", "four": "4",
        "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
    }

    // Open the file
    file, err := os.Open("input2.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var totalSum int
    scanner := bufio.NewScanner(file)

    // Read line by line
    for scanner.Scan() {
        line := scanner.Text()
        words := strings.Fields(line)
        firstDigit := ""
        lastDigit := ""

        // Find first and last digit
        for _, word := range words {
            if num, exists := digitMap[word]; exists {
                if firstDigit == "" {
                    firstDigit = num
                }
                lastDigit = num
            }
        }

        // Convert to integers and sum them
        if firstDigit != "" && lastDigit != "" {
            digitInt, err := strconv.Atoi(firstDigit + lastDigit)
            if err != nil {
                fmt.Println("Error converting digits to integer:", err)
                continue
            }

            totalSum += digitInt
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading from file:", err)
    }

    fmt.Println("Total sum:", totalSum)
}



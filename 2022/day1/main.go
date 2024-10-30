package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Error opening the file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    currentSum := 0
    sumOfCalories := []int{}

    for scanner.Scan() {
        line := scanner.Text()

        if line == "" {
            sumOfCalories = append(sumOfCalories, currentSum)
            currentSum = 0
        } else {
            calories, err := strconv.Atoi(line)
            if err != nil {
                fmt.Println("Error converting data to integers:", err)
                return
            }
            currentSum += calories
        }
    }

    if currentSum > 0 {
        sumOfCalories = append(sumOfCalories, currentSum)
    }

    sort.Sort(sort.Reverse(sort.IntSlice(sumOfCalories)))

    totalSum := sumOfCalories[0] + sumOfCalories[1] + sumOfCalories[2]
    fmt.Printf("The total calories carried by the top 3 Elves is: %d\n", totalSum)
}

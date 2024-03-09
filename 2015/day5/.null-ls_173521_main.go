package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)
func countVowels(word string) int {
  vowels := "aeiou"
  count := 0

  for _, c := range word {
    for _, v := range vowels {
      if c == v{
        count++
        break
      }
    }
  }
  return count
}

func doubleLetter(word string) bool {
  foundDoubleLetter := false
  for i := 1; i < len(word); i++{
    if line[i] == line[i - 1]{
      foundDoubleLetter = true
      break
    }
  }
  return foundDoubleLetter
}

func doesNotContain(word string) bool {
  for _, sub := range []string{"ab", "cd", "pq", "xy"} {
  if strings.Contains(str, sub){
    return true
   }
 }
 return false
}

func main () {

  file, err := os.Open("input.txt")
  if err != nil {
    fmt.Println("error opening the file")
    return
  }
  defer file.Close()
  
  goodWords := 0 

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text() 
    vowelCount := countVowels(line)
    doubleLet := doubleLetter(line)
    containsSub := doesNotContain(line)
    if vowelCount >= 3 && doubleLet && !containsSub{
     goodWords++
    }
    }
  }

fmt.Println(goodWords)
  

}

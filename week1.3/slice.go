package main

import (
  "fmt"
  "sort"
  "strconv"
  "strings"
)

func main() {
  var input string
  slice := make([]int, 0, 3)

  for {
    fmt.Println("Enter an integer to add it to the sorted slice, or 'X' to exit.")
    fmt.Scan(&input)
    if strings.Contains(input, "X") {
      fmt.Println("Exiting")
      break
    }
    userInput, err := strconv.Atoi(input)
    if err != nil {
			fmt.Println(err)
			continue
		}
    slice = append(slice, userInput)
    if !sort.SliceIsSorted(slice, func(i int, j int) bool {return slice[i] < slice[j] }) {
      sort.SliceStable(slice, func(i int, j int) bool {return slice[i] < slice[j] })
    }
    fmt.Println("Sorted slice: ", slice)
  }
}

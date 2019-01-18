package main

import (
  "fmt"
)

func main() {
  var input int
  slice := make([]int, 0, 10)

  fmt.Println("Enter up to 10 integers to sort: ")
  for len(slice) < 10 {
    fmt.Scan(&input)
    slice = append(slice, input)
  }
  BubbleSort(slice)
  fmt.Print("After sorting: ")
  fmt.Println(slice)
}

func BubbleSort(input []int) {
  var (
    n = len(input)
    sorted = false
  )

  for !sorted {
    swapped := false
    for i := 0; i < n-1; i++ {
      if input[i] > input[i+1] {
        Swap(input, i)
        swapped = true
      }
    }
    if !swapped {
      sorted = true
    }
    n = n - 1
  }
}

func Swap(input []int, i int) {
  input[i], input[i+1] = input[i+1], input[i]
}

package main

import (
  "fmt"
  "strconv"
)

func main() {
  var a float64
  fmt.Printf("Enter a floating point number: ")
  fmt.Scan(&a)
  b := int(a)
  fmt.Printf(strconv.Itoa(b))
}

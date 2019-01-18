package main

import (
  "fmt"
  "math"
)

func main() {
  var a, v, s, t float64
  fmt.Print("Enter acceleration: ")
  fmt.Scan(&a)
  fmt.Print("Enter velocity: ")
  fmt.Scan(&v)
  fmt.Print("Enter displacement: ")
  fmt.Scan(&s)
  fmt.Print("Enter time: ")
  fmt.Scan(&t)

  fn := GetDisplaceFn(a, v, s)

  fmt.Println("Dispacement: ")
  fmt.Println(fn(t))
}

func GetDisplaceFn(a, v, s float64) func(t float64) float64 {
  fn := func(t float64) float64 {
    return (0.5*a) * math.Pow(t, 2) + v * t + s
  }
  return fn
}

package main

import (
  "fmt"
)

func main() {
  fmt.Printf("Enter a list of integers:\n")
  var input int
  slice := make([]int, 0, 12)
  for len(slice) < 12 {
    fmt.Scan(&input)
    slice = append(slice, input)
  }

  c := make(chan []int)

  go MergeSort(slice, c)
  cData := <- c

  close(c)

  fmt.Printf("%v\n%v\n", slice, cData)
}

func Merge(l, r []int) []int {
	ret := make([]int, 0, len(l)+len(r))
	for len(l) > 0 || len(r) > 0 {
		if len(l) == 0 {
			return append(ret, r...)
		}
		if len(r) == 0 {
			return append(ret, l...)
		}
		if l[0] <= r[0] {
			ret = append(ret, l[0])
			l = l[1:]
		} else {
			ret = append(ret, r[0])
			r = r[1:]
		}
	}
	return ret
}

func MergeSort(s []int, c chan []int) {
  fmt.Printf("Sorting slice: %v\n", s)
  if len(s) <= 1 {
    c <- s
    return
  }

  lChan := make(chan []int)
  l2Chan := make(chan []int)
  rChan := make(chan []int)
  r2Chan := make(chan []int)

  n := len(s) / 2
  q := n / 2

  go MergeSort(s[:q], lChan)
  go MergeSort(s[q:n], l2Chan)
  go MergeSort(s[n:n+q], rChan)
  go MergeSort(s[n+q:], r2Chan)

  lData := <- lChan
  l2Data := <- l2Chan
  rData := <- rChan
  r2Data := <- r2Chan

  close(lChan)
  close(l2Chan)
  close(rChan)
  close(r2Chan)

  h1Data := Merge(lData, l2Data)
  h2Data := Merge(rData, r2Data)

  c <- Merge(h1Data, h2Data)

  return
}

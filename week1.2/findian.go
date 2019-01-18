package main

import (
  "bufio"
  "fmt"
  "strings"
  "os"
)

func main() {
  fmt.Printf("Enter a string:\n")
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  s := scanner.Text()
  s = strings.ToLower(s)

  if strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
    fmt.Println("Found!")
  } else {
    fmt.Println("Not found!")
  }
}

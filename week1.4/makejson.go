package main

import (
  "bufio"
  "encoding/json"
  "fmt"
  "os"
  "strings"
)

func main() {
  personAddress := make(map[string]string)
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter your name: ")
  name, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Print("Enter your address: ")
  address, err := reader.ReadString('\n')
  if err != nil {
    fmt.Println(err)
    return
  }
  personAddress["name"] = strings.Trim(name, "\n")
  personAddress["address"] = strings.Trim(address, "\n")
  b, err := json.Marshal(personAddress)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(b))
}

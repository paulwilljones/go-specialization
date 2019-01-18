package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

type Name struct {
  fname string
  lname string
}

func main() {
  var nameList []Name
  var fileName string
  fmt.Printf("Enter filename: ")
  if _, err := fmt.Scan(&fileName); err != nil {
    fmt.Println(err)
    return
  }
  file, err := os.Open(fileName)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()
  scanner := bufio.NewScanner(bufio.NewReader(file))
  for scanner.Scan() {
    line := scanner.Text()
    splitLine := strings.SplitN(line, " ", 2)
    name := Name{fname: splitLine[0], lname: splitLine[1]}
    nameList = append(nameList, name)
  }
  for _, name := range nameList {
    fmt.Printf("\n%s %s\n", name.fname, name.lname)
  }
}

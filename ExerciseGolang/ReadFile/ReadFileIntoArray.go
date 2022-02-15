package main

import (
  "os"
  "strings"
  "fmt"
)
func main() {
  byte, err := os.ReadFile("file.txt")
  if err != nil {
    return
  }
  text := string(byte)
  arr := strings.Split(text,"\n")
  fmt.Println(arr)
}

  
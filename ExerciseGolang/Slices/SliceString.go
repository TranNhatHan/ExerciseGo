package main

import (
  "fmt"
)

func main() {
  z := "Hello world"
  a := z[0:6]
  b := z[6:]
  fmt.Println(a)
  fmt.Println(b)
}
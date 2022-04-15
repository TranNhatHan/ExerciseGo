package main

import (
  "fmt"
)

func show(x... string) {
  for _, name := range x {
    fmt.Println(name)
  }
}

func main() {
  show("Tran", "Nhat", "Han")
}
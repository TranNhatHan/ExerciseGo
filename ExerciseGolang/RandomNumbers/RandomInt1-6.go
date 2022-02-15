package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  ranint := rand.Intn(6)+1
  fmt.Println(ranint)
}
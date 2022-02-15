package main

import (
  "fmt"
)

type House struct {
  roomnumber int
  price int
  place string
}

func main() {
  var a House
  a.roomnumber = 5
  a.price = 50000
  a.place = "Hue"
  fmt.Println("Price of this house:",a.price)
}
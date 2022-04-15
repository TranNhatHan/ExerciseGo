package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scanln(&number)
	if number > 0 {
		fmt.Println(float32(number)/2)
	}
 }
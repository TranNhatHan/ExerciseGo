package main

import (
    "fmt"
    "strconv" 
    "os"
)
 func main() {
 	var sum int
 	for i:=1; i<=5; i++ {
 		var val string
 		fmt.Print(">>")
 		fmt.Scanln(&val)
        weight, err := strconv.Atoi(val)
        if err != nil {
            fmt.Println("Input must an integer!")
            os.Exit(0)
        }
 		sum+=weight
 	}
 	fmt.Println("Average Weight:", sum/5)
 }
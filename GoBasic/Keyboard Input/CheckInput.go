package main
import (
    "fmt"
    "strconv" 
    "os" 
)

func main() {
    fmt.Print("Enter a number: ")
    var str string
    fmt.Scanln(&str)
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("Input not an integer!")
        os.Exit(0)
    }
    if num>=0 && num<=10 {
        fmt.Println("Correct")
    } else {
        fmt.Println("Incorrect")
    }
}
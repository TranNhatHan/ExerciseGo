package main
  
import (
	"fmt"
	"time"
	"os"
)
  
// Main function
func main() {
	var age int
	var mob int
	var dob int
	fmt.Print("Your age: ")
	fmt.Scanln(&age)
	fmt.Print("Month of birth: ")
	fmt.Scanln(&mob)
	fmt.Print("Day of birth: ")
	fmt.Scanln(&dob)
	if dob > 31 && mob > 12 {
		fmt.Println("Wrong")
		os.Exit(0)
	}
    today := time.Now()
    birthday := time.Date(today.Year(), time.Month(mob), dob, 0, 0, 0, 0, time.UTC)
	if today.After(birthday) {
		fmt.Println("Your birthyear:",today.Year()-age)
	} else {
		fmt.Println("Your birthyear:",today.Year()-age-1)
	}
}
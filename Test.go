// Golang program to compare times
package main
  
import "fmt"
  
// importing time module
import "time"
  
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
    today := time.Now()
    birthday := time.Date(today.Year(), time.Month(mob), dob, 0, 0, 0, 0, time.UTC)
	fmt.Println(birthday)
	if today.After(birthday) {
		fmt.Println("Your birthyear:",today.Year()-age)
	} else {
		fmt.Println("Your birthyear:",today.Year()-age-1)
	}
}
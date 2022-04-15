package main
import "fmt"

func main() {
	var name string
	fmt.Print("Your name is:")
	fmt.Scanln(&name)
	fmt.Println("Hello",name)
}
// constants cannot be declared using the := syntax//
//constants are declared like variables but with the const keyword..


package main
import "fmt"


const Pi = 3.14

func main() {
	const World = "People"
	fmt.Println("Hello", World)
	fmt.Println("Happy" , Pi, "Day")

	const Truth = "true"
	fmt.Println("Go rules?", Truth)
}

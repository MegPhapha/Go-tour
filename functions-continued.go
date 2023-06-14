/* in this example, x int , y int is shortened to 
x , y int */

package main
import ("fmt")

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
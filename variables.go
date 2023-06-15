// a var staement can e at package and functions levels.
// both is displayed in this example

package main
import ("fmt")

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)
}
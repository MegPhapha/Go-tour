// the variable will take the type of the initializer.
// the type can be ommitted if an initializer is present.
// a var declaration can include an initializer per a variable

package main
import ("fmt")

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}



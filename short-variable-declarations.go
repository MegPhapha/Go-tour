/* the := short assignment statement can be used
instead of a var declaration
inside a function with implicit type */

package main

import (
	"fmt"
)

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

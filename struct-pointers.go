// struct fields can be accessed through a struct pointer.
/* cumbersome notation= (*p).X.
write it this way instead = p.X
without the explicit derefence. */

package main
import ("fmt")

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(p.X)
	fmt.Println(v)
}
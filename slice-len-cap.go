/* A slice has both lenth and capacity
the length of a slice is the number of elements it contains
the capacity is thenumber of elements in the underlying
arraym counting from the first element in the slice. */

package main
import ("fmt")

func main() {
	s := []int{2, 3, 5, 7, 11, 13 }
	fmt.PrintSlice(s) 
}
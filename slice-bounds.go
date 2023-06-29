/* You may omit the high or low bounds to use their defaults
instead when slicing.
The default is zero for low bound and the lenth of the slice 
for high bound. */

package main
import ("fmt")

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

}
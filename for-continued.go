// the init and post statements are optional
//the init and post statements are optional

package main
import ("fmt")

func main() {
	sum := 1
	for ; sum < 1000; {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println(sum)
}
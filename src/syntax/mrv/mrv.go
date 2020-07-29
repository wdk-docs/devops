package main
import "fmt"

func mrv() (int,int) {
	return 3,7
}
func main()  {
	a,b := mrv()
	fmt.Println("a:",a)
	fmt.Println("b:",b)
	_,c := mrv()
	fmt.Println("c:",c)
}

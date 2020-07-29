package main
import "fmt"
func main() {
		fmt.Println("最基本的类型，有一个条件。")
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i++
    }
		fmt.Println("一个典型的初始/条件/ for循环之后。")
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }
		fmt.Println("对于没有条件将循环反复，直到你打破从封闭功能的环路或返回了。")
    for {
        fmt.Println("loop")
        break
    }

		fmt.Println("您也可以继续循环的下一次迭代。")
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}

package main
import "fmt"
// 下面是将整数作为参数的任意数的函数。

func sum(nums ...int) {
  fmt.Print(nums, " ")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}
func main() {
    // 参数可变型函数可以与单个参数的常用方法被调用。

    sum(1, 2)
    sum(1, 2, 3)

    // 如果你已经有一个多片ARGS, 它们适用于使用FUNC（切片...）这样的可变参数函数。

    nums := []int{1, 2, 3, 4}
    sum(nums...)

}

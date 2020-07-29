package main
import "fmt"
func main() {
    // 这里有一个基本的例子。
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }
    // 你可以有一个if语句没有别的。
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }
    // 声明可以先条件语句;在此声明中声明的任何变量在各分公司提供。
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

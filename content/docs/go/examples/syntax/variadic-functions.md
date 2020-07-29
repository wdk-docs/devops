---
title: '变数函数 - func sum(nums ...int){}'
linkTitle: '变数函数'
weight: 13
description: >
  变参函数可调用任意数量的尾参。
type: 'docs'
---

例如，`fmt.Println`是一种常见的可变参数函数。

```go
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

    // 如果你已经有多个ARGS的切片，它们适用于使用FUNC（切片...）这样的可变参数函数。

    nums := []int{1, 2, 3, 4}
    sum(nums...)

}
```

```sh
$ go run variadic-functions.go
[1 2] 3
[1 2 3] 6
[1 2 3 4] 10
```

在 go 函数的另一个重要方面是它们的形式关闭，我们将看看未来的能力。

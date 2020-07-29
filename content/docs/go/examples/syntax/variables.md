---
title: '变量 - variable'
linkTitle: ''
weight: 2
description: >
  在`Go`，变量显式声明的编译器，以用于例如检查的函数调用类型的正确性。
type: 'docs'
---

VAR 声明 1 个或多个变量。

您可以一次声明多个变量。

`Go`将推断初始化的变量的类型。

没有相应的初始化声明的变量是零值。 例如，对于一个无值 `int` 是 `0`。

`:=`语法简写用于声明和初始化的变量, 例如 在本例中为 `var f string = "apple"` 语法简写.

```go
package main
import "fmt"
func main() {
    var a = "initial"
    fmt.Println(a)
    var b, c int = 1, 2
    fmt.Println(b, c)
    var d = true
    fmt.Println(d)
    var e int
    fmt.Println(e)
    f := "apple"
    fmt.Println(f)
}
```

```sh
$ go run variables.go
initial
1 2
true
0
apple
```

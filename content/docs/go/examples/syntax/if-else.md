---
title: '条件语句 - if/else'
linkTitle: ''
weight: 6
description: >
  分枝如果和其他人在go是直截了当的。
type: 'docs'
---

```go
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
```

请注意，您不需要在周围条件去括号，但所需的括号。

```sh
$ go run if-else.go
7 is odd
8 is divisible by 4
9 has 1 digit
```

没有三元如果`Go`，所以你需要使用一个完整的 if 语句即使是基本的条件。

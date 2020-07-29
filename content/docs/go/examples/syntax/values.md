---
title: '值 - value'
weight: 1
description: >
  Go 有不同的值类型，包括字符串，整数，浮点数，布尔值等。这里有几个基本的例子。
type: 'docs'
---

串，其可以连同+加入。

整数和浮点数。

布尔值，用逻辑运算符如你所期望。

```go
package main
import "fmt"
func main() {
    fmt.Println("go" + "lang")
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
```

```shell
$ go run values.go
golang
1+1 = 2
7.0/3.0 = 2.3333333333333335
false
true
false
```

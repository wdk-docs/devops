---
title: '多返回值 - func vals() (int, int){}'
linkTitle: '多返回值'
weight: 12
description: >
  Go有内置的多个返回值的支持。
type: 'docs'
---

该功能在惯用的'Go`经常使用, 例如用于从函数返回两个结果和错误值。

```go
package main
import "fmt"
// The (int, int) in this function signature shows that the function returns 2 ints.

func vals() (int, int) {
    return 3, 7
}
func main() {
    // Here we use the 2 different return values from the call with multiple assignment.

    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // If you only want a subset of the returned values, use the blank identifier \_.

    _, c := vals()
    fmt.Println(c)

}
```

```sh
$ go run multiple-return-values.go
3
7
7
```

Accepting a variable number of arguments is another nice feature of Go functions; we’ll look at this next.

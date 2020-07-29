---
title: '递归 - Recursion'
linkTitle: ''
weight: 15
description: >
  `Go`支持递归函数。
type: 'docs'
---

下面是一个典型的阶乘的例子。

直到它到达的`func(0)`基情况下，这实际上函数调用自身。

```go
package main
import "fmt"

func fact(n int) int {
  if n == 0 {
    return 1
  }
  return n * fact(n-1)
}
func main() {
  fmt.Println(fact(7))
}
```

```sh
$ go run recursion.go
5040
```

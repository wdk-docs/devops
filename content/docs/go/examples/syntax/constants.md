---
title: '常量 - constant'
linkTitle: ''
weight: 3
description: >
  `Go`支持字符，字符串，布尔和数值的常量。
type: 'docs'
---

一个 const 语句可以出现在任何 var 语句出现地方。

常量表达式执行任意精度算术。

一个数字`constant`没有类型，直到它给一个, 例如通过显式转换.

一些可以通过在需要一个如变量赋值或函数调用的上下文使用它被给予一个类型。
例如, 这里 math.Sin 期望一个 float64。

```go
package main
import (
"fmt"
"math"
)
// 常量声明一个恒定值。

const s string = "constant"
func main() {
  fmt.Println(s)
    const n = 500000000
    const d = 3e20 / n
    fmt.Println(d)
    fmt.Println(int64(d))
    fmt.Println(math.Sin(n))
}
```

```sh
$ go run constant.go
constant
6e+11
600000000000
-0.28470407323754404
```

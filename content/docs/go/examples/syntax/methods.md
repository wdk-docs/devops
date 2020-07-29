---
title: '方法 - Methods'
linkTitle: ''
weight: 18
description: >
  `Go`支持在结构类型定义的方法。
type: 'docs'
---

```go
package main
import "fmt"
type rect struct {
    width, height int
}
// 此区域方法具有`* rect`接收器类型。

func (r *rect) area() int {
  return r.width * r.height
}
// 方法可以用于无论是指针或值的接收器类型来定义。
// 这里有一个值接收机的例子。

func (r rect) perim() int {
  return 2*r.width + 2*r.height
}
func main() {
    r := rect{width: 10, height: 5}
    // 在这里，我们呼吁我们的结构定义的2种方法。
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())
    // `Go`自动处理值和指针为方法调用之间的转换。
    // 您可能需要使用一个指针接收器类型，以避免复制的方法调用或允许的方法来突变接收结构。
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
```

```sh
$ go run methods.go
area: 50
perim: 30
area: 50
perim: 30
```

接下来我们来看看`Go`的机制，分组和命名相关的方法集: [interfaces](interfaces.md).

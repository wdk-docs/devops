---
title: '接口 - Interfaces'
linkTitle: ''
weight: 19
description: >
  接口被命名方法签名的集合。
type: 'docs'
---

```go
package main
import (
"fmt"
"math"
)
// 下面是几何形状基本接口。

type geometry interface {
  area() float64
  perim() float64
}
// 在我们的例子中，我们将实现在`rect`和`circle`类型此接口。
type rect struct {
  width, height float64
}
type circle struct {
  radius float64
}
// 为了实现在Go接口，我们只需要实现该接口的所有方法。
// 在这里，我们实现几何上rects。

func (r rect) area() float64 {
  return r.width * r.height
}
func (r rect) perim() float64 {
  return 2*r.width + 2*r.height
}
// The implementation for circles.

func (c circle) area() float64 {
  return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
  return 2 * math.Pi * c.radius
}
// If a variable has an interface type, then we can call methods that are in the named interface. Here’s a generic measure function taking advantage of this to work on any geometry.

func measure(g geometry) {
  fmt.Println(g)
  fmt.Println(g.area())
  fmt.Println(g.perim())
}
func main() {
  r := rect{width: 3, height: 4}
  c := circle{radius: 5}
  // The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
  measure(r)
  measure(c)

}
```

```sh
$ go run interfaces.go
{3 4}
12
14
{5}
78.53981633974483
31.41592653589793
```

要了解更多关于`Go`的接口，检查出这个伟大的博客文章。

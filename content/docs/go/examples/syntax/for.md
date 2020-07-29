---
title: '循环 - for'
linkTitle: ''
weight: 5
description: >
  `for`是Go唯一的循环结构。
type: 'docs'
---

下面是一些基本类型的 for 循环。

```go
package main
import "fmt"
func main() {
    // 最基本的类型，有一个条件。
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
    // 一个典型的初始/条件/ for循环之后。
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }
    // 对于没有条件将循环反复，直到你打破从封闭功能的环路或返回了。
    for {
        fmt.Println("loop")
        break
    }
    // 您也可以继续循环的下一次迭代。
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
}
```

```sh
$ go run for.go
1
2
3
7
8
9
loop
1
3
5
```

稍后我们会看到一些其他的形式，当我们看报表范围，渠道和其他数据结构。

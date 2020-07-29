---
title: '闭包 - Closures'
linkTitle: ''
weight: 14
description: >
  `Go`支持匿名的功能，其可以形成封闭。
type: 'docs'
---

当你要定义一个函数内联，而无需将其命名为匿名函数是有用的。

```go
package main
import "fmt"
// 这个函数返回intSeq另一个函数，这是我们在intSeq的身体匿名定义。
// 返回的函数关闭在变量i以形成闭合。

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
func main() {
    // 我们带哦用intSeq，分配结果（函数）来nextInt。
    // 这个函数值撷取自己的I值，这将是每一个我们称之为nextInt时间更新。

    nextInt := intSeq()

    // 通过调用nextInt几次看到封闭的效果。

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // 要确认状态是唯一的特定功能，创建并测试一个新的。

    newInts := intSeq()
    fmt.Println(newInts())

}
```

```sh
$ go run closures.go
1
2
3
1
```

我们来看看现在的函数的最后一项功能是递归。

---
title: '数组 - [no]type'
linkTitle: ''
weight: 7
description: >
  在`Go`，阵列是具有特定长度的元件的编号序列。
type: 'docs'
---

```go
package main
import "fmt"
func main() {
    // 在这里，我们创建一个将举行整整5个整数组成的数组。
    // 类型元素和长度的是阵列的类型的两个部分。
    // 通过默认的阵列是零值，这对于整数手段0。
    var a [5]int
    fmt.Println("emp:", a)
    // 我们可以使用`阵列[索引] = value`语法索引在设定的值，并获得与`阵列[索引]`的值。
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
    // 内建len个返回的数组的长度。
    fmt.Println("len:", len(a))
    // Use this syntax to declare and initialize an array in one line.
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)
    // Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
```

Note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println.

```sh
$ go run arrays.go
emp: [0 0 0 0 0]
set: [0 0 0 0 100]
get: 100
len: 5
dcl: [1 2 3 4 5]
2d: [[0 1 2][1 2 3]]
```

You’ll see slices much more often than arrays in typical Go. We’ll look at slices next.

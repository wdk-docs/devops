---
title: '切片 - slices'
linkTitle: ''
weight: 8
description: >
  切片是在Go一个密钥数据类型，得到更强大的界面比阵列的序列。
type: 'docs'
---

```go
package main
import "fmt"
func main() {
    // 不同于阵列，切片仅通过它们所包含的元件（未元素的数目）键入。
    // 要创建具有非零长度的空片，可以使用内建的make。
    // 在这里，我们做一个长度为3的字符串的切片（最初零值）。

    s := make([]string, 3)
    fmt.Println("emp:", s)

    // 我们可以设置和获取，就像使用数组。

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // LEN返回切片的长度为预期的。

    fmt.Println("len:", len(s))

    // 除了这些基本操作，切片支持几声，使它们比数组更丰富。
    // 一个是内置的追加，返回包含一个或多个新值的切片。
    // 需要注意的是，我们需要接受来自追加一个返回值，我们可以得到一个新的切片值。

    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // 片也可以copy'd。在这里，我们创建相同的长度为s的空片c和复制到c中从s。

    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // 片支持“切片”操作的语法片[low:high].
    // 例如，该获取的元素的切片 s[2], s[3], and s[4].

    l := s[2:5]
    fmt.Println("sl1:", l)

    // 该切片达（但不包括） s[5].

    l = s[:5]
    fmt.Println("sl2:", l)

    // 而这个切片了从（包括） s[2].

    l = s[2:]
    fmt.Println("sl3:", l)

    // 我们可以声明和在一行初始化切片的变量也是如此。

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // 片可以由成多维数据结构。
    // 内片的长度可以变化，不像多维数组。

    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

}
```

请注意，虽然片不同类型的阵列相比，它们也同样由 fmt.Println 呈现。

```sh
$ go run slices.go
emp: [ ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d: [[0][1 2] [2 3 4]]
```

看看由 Go 队这个伟大的博客文章对切片在 Go 的设计和实现的更多细节。

现在，我们已经看到了数组和切片，我们将看看 Go 的其他关键数据内置结构: maps.

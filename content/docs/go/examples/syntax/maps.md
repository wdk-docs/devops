---
title: '映射 - map'
linkTitle: ''
weight: 9
description: >
  **map**是Go的内置关联的数据类型(有时在其他语言中被称为 _hash_ 或 _dicts_ ).
type: 'docs'
---

```go
package main
import "fmt"
func main() {
    // 要创建一个空的map, 使用内置make: make(map[key-type]val-type).
    m := make(map[string]int)
    // 使用典型的 `name[key] = val` 语法组键/值对。
    m["k1"] = 7
    m["k2"] = 13
    // 打印具有例如一个map`fmt.Println`将展示其所有键/值对。
    fmt.Println("map:", m)
    // 获取的值与`name[key]`的关键。
    v1 := m["k1"]
    fmt.Println("v1: ", v1)
    // 内建len个在map上调用时返回键/值对的数量。
    fmt.Println("len:", len(m))
    // 从map内置delete删除键/值对。
    delete(m, "k2")
    fmt.Println("map:", m)
    // 获得从映射中的值时，所述任选的第二返回值表示如果键存在于map。
    // 这可以用象0或“”零个值被用来消除丢失的钥匙和钥匙之间。
    // 在这里，我们并不需要本身的价值，所以我们用空白标识 _ 忽略了它。
    _, prs := m["k2"]
    fmt.Println("prs:", prs)
    // 您也可以在同一行使用此语法声明并初始化新map。
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

}
```

请注意，当用`fmt.Println`打印时`map`显示在样式`map[k:v k:v]`。

```sh
$ go run maps.go
map: map[k1:7 k2:13]
v1: 7
len: 2
map: map[k1:7]
prs: false
map: map[bar:2 foo:1]
```

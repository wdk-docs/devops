---
title: '结构体 - Structs'
linkTitle: ''
weight: 17
description: >
  Go 的结构体(struct) 是带类型的字段(fields)集合。
type: 'docs'
---

这在组织数据时非常有用。

```go
package main
import "fmt"
// 这里的 person 结构体包含了 name 和 age 两个字段。
type person struct {
    name string
    age int
}
// newPerson构建一个新的人结构用的名字。
func newPerson(name string) *person {
    // 您可以将指针安全返回局部变量作为一个局部变量将生存函数的范围。
    p := person{name: name}
    p.age = 42
    return &p
}
func main() {
    // 这句法创建一个新的结构。
    fmt.Println(person{"Bob", 20})
    // 初始化一个结构时，您可以命名的字段。
    fmt.Println(person{name: "Alice", age: 30})
    // 省略字段将是零值。
    fmt.Println(person{name: "Fred"})
    // 一个＆前缀产生的指针结构。
    fmt.Println(&person{name: "Ann", age: 40})
    // 这是惯用的封装在构造函数新的结构生成
    fmt.Println(newPerson("Jon"))
    // 访问结构以点领域。
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)
    // 您还可以使用点使用结构指针 - 指针会自动取消引用。
    sp := &s
    fmt.Println(sp.age)
    // 结构是可变的。
    sp.age = 51
    fmt.Println(sp.age)
}
```

```sh
$ go run structs.go
{Bob 20}
{Alice 30}
{Fred 0}
&{Ann 40}
&{Jon 42}
Sean
50
51
```

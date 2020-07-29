---
title: '开关 - switch'
linkTitle: ''
weight: 7
description: >
  switch语句表达在许多分支条件语句。
type: 'docs'
---

```go
package main
import (
"fmt"
"time"
)
func main() {
    // 这是一个基本的开关。
    i := 2
    fmt.Print("Write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }
    // 您可以使用逗号多个表达式在同一案例中的语句分开。 我们在这个例子中使用可选默认情况下也是如此。
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
    // 没有表达式开关是为了表达的if / else逻辑的另一种方法。 在这里，我们还显示的情况下表现如何可以是非常数。
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
    // A型开关进行比较而不是值类型。 您可以使用此发现的接口值的类型。 在此示例中，变量t将具有对应于它的条款的类型。
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
}
```

```sh
$ go run switch.go
Write 2 as two
It's a weekday
It's after noon
I'm a bool
I'm an int
Don't know type string
```

---
title: "Hello World"
linkTitle: "Hello World"
weight: 1
description: >
  我们的第一个程序将打印经典的“Hello World”的消息。下面是完整的源代码。
type: "docs"
---

```go
package main
import "fmt"
func main() {
    fmt.Println("hello world")
}
```

要运行程序，把代码中的 HELLO-world.go 和使用 Go 运行。

```sh
$ go run hello-world.go
hello world
```

有时我们会想建立我们的节目成二进制文件。我们可以做到这一点使用去构建。

```sh
$ go build hello-world.go
$ ls
hello-world    hello-world.go
```

然后，我们可以直接执行构建的二进制。

```sh
$ ./hello-world
hello world
```

现在，我们可以运行和构建基本围棋程序，让我们了解更多的语言。

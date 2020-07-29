---
title: "Panic"
linkTitle: ""
weight: 100
description: >
  一个`panic`通常意味着出事意外错误。 主要是我们用它来在正常运行期间不应该出现的错误快速失败了，或者说，我们不准备妥善处理。
type: "docs"
---

```go
package main
import "os"
func main() {
We’ll use panic throughout this site to check for unexpected errors. This is the only program on the site designed to panic.

    panic("a problem")

A common use of panic is to abort if a function returns an error value that we don’t know how to (or want to) handle. Here’s an example of panicking if we get an unexpected error when creating a new file.

    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }

}
Running this program will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status.

\$ go run panic.go
panic: a problem
goroutine 1 [running]:
main.main()
/.../panic.go:12 +0x47
...
exit status 2
Note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible.
```

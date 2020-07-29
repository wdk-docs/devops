---
title: "Defer"
linkTitle: ""
weight: 100
description: >
  `Defer`被用来确保一个函数调用以程序的执行之后进行，通常是为了清理的目的。 `defer`通常用于其中例如保证最后会在其他语言中使用。
type: "docs"
---

```go
package main
import (
"fmt"
"os"
)
Suppose we wanted to create a file, write to it, and then close when we’re done. Here’s how we could do that with defer.

func main() {
Immediately after getting a file object with createFile, we defer the closing of that file with closeFile. This will be executed at the end of the enclosing function (main), after writeFile has finished.

    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)

}
func createFile(p string) *os.File {
fmt.Println("creating")
f, err := os.Create(p)
if err != nil {
panic(err)
}
return f
}
func writeFile(f *os.File) {
fmt.Println("writing")
fmt.Fprintln(f, "data")
}
It’s important to check for errors when closing a file, even in a deferred function.

func closeFile(f \*os.File) {
fmt.Println("closing")
err := f.Close()
if err != nil {
fmt.Fprintf(os.Stderr, "error: %v\n", err)
os.Exit(1)
}
}
Running the program confirms that the file is closed after being written.

\$ go run defer.go
creating
writing
closing
```

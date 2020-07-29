---
title: "Exit"
linkTitle: ""
weight: 100
description: >
  使用`os.Exit`立即退出与给定的状态。
type: "docs"
---

defers will not be run when using os.Exit, so this fmt.Println will never be called.

Note that unlike e.g. C, Go does not use an integer return value from main to indicate exit status. If you’d like to exit with a non-zero status you should use os.Exit.

```go
package main
import (
"fmt"
"os"
)
func main() {

    defer fmt.Println("!")

Exit with status 3.

    os.Exit(3)

}
```

If you run exit.go using go run, the exit will be picked up by go and printed.

\$ go run exit.go
exit status 3
By building and executing a binary you can see the status in the terminal.

$ go build exit.go
$ ./exit
$ echo $?
3
Note that the ! from our program never got printed.

```

```

---
title: "Spawning Processes"
linkTitle: ""
weight: 100
description: >

type: "docs"
---

Sometimes our Go programs need to spawn other, non-Go processes. For example, the syntax highlighting on this site is implemented by spawning a pygmentize process from a Go program. Let’s look at a few examples of spawning processes from Go.

```go
package main
import (
"fmt"
"io/ioutil"
"os/exec"
)
func main() {
We’ll start with a simple command that takes no arguments or input and just prints something to stdout. The exec.Command helper creates an object to represent this external process.

    dateCmd := exec.Command("date")

.Output is another helper that handles the common case of running a command, waiting for it to finish, and collecting its output. If there were no errors, dateOut will hold bytes with the date info.

    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

Next we’ll look at a slightly more involved case where we pipe data to the external process on its stdin and collect the results from its stdout.

    grepCmd := exec.Command("grep", "hello")

Here we explicitly grab input/output pipes, start the process, write some input to it, read the resulting output, and finally wait for the process to exit.

    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("hello grep\ngoodbye grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

We omitted error checks in the above example, but you could use the usual if err != nil pattern for all of them. We also only collect the StdoutPipe results, but you could collect the StderrPipe in exactly the same way.

    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

Note that when spawning commands we need to provide an explicitly delineated command and argument array, vs. being able to just pass in one command-line string. If you want to spawn a full command with a string, you can use bash’s -c option:

    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))

}
The spawned programs return output that is the same as if we had run them directly from the command-line.

\$ go run spawning-processes.go

> date
> Wed Oct 10 09:53:11 PDT 2012
> grep hello
> hello grep
> ls -a -l -h
> drwxr-xr-x 4 mark 136B Oct 3 16:29 .
> drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
> -rw-r--r-- 1 mark 1.3K Oct 3 16:28 spawning-processes.go
```

---
title: "Context"
linkTitle: ""
weight: 100
description: >
  在前面的例子中我们看到了建立一个简单的HTTP服务器。 HTTP servers are useful for demonstrating the usage of context.Context for controlling cancellation. A Context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines.
type: "docs"
---

```go
package main
import (
"fmt"
"net/http"
"time"
)
func hello(w http.ResponseWriter, req \*http.Request) {
A context.Context is created for each request by the net/http machinery, and is available with the Context() method.

    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

Wait for a few seconds before sending a reply to the client. This could simulate some work the server is doing. While working, keep an eye on the context’s Done() channel for a signal that we should cancel the work and return as soon as possible.

    select {
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done():

The context’s Err() method returns an error that explains why the Done() channel was closed.

        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }

}
func main() {
As before, we register our handler on the “/hello” route, and start serving.

    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":8090", nil)

}
Run the server in the background.

\$ go run context-in-http-servers.go &
Simulate a client request to /hello, hitting Ctrl+C shortly after starting to signal cancellation.

\$ curl localhost:8090/hello
server: hello handler started
^C
server: context canceled
server: hello handler ended
```

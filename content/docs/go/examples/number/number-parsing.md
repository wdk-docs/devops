---
title: "Number Parsing"
linkTitle: ""
weight: 100
description: >

type: "docs"
---

Parsing numbers from strings is a basic but common task in many programs; here’s how to do it in Go.

```go
package main
The built-in package strconv provides the number parsing.

import (
"fmt"
"strconv"
)
func main() {
With ParseFloat, this 64 tells how many bits of precision to parse.

    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

For ParseInt, the 0 means infer the base from the string. 64 requires that the result fit in 64 bits.

    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

ParseInt will recognize hex-formatted numbers.

    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

A ParseUint is also available.

    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

Atoi is a convenience function for basic base-10 int parsing.

    k, _ := strconv.Atoi("135")
    fmt.Println(k)

Parse functions return an error on bad input.

    _, e := strconv.Atoi("wat")
    fmt.Println(e)

}
\$ go run number-parsing.go
1.234
123
456
789
135
strconv.ParseInt: parsing "wat": invalid syntax
Next we’ll look at another common parsing task: URLs.
```
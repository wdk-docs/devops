---
title: "Sorting"
linkTitle: ""
weight: 100
description: >
  Go的排序包实现分拣内建和用户自定义类型。 我们将在第一分拣建宏。
type: "docs"
---

```go
package main
import (
"fmt"
"sort"
)
func main() {
Sort methods are specific to the builtin type; here’s an example for strings. Note that sorting is in-place, so it changes the given slice and doesn’t return a new one.

    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

An example of sorting ints.

    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

We can also use sort to check if a slice is already in sorted order.

    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)

}
Running our program prints the sorted string and int slices and true as the result of our AreSorted test.

\$ go run sorting.go
Strings: [a b c]
Ints: [2 4 7]
Sorted: true
```

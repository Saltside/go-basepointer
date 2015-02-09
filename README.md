# Pointers

This package provides functions that return a pointer to their argument. This
is useful when a pointer is required but you don't have a variable defined etc.

E.g.

```go
package main

import (
	. "antiklimax.se/go-basepointer"
)

var x *string = StringPtr("point to me")
var y *float64 = Float64Ptr(3.141592)
```

# Example

```go
package main

import (
	"github.com/lbb4511/code/qrencode"
)

func main() {
	str := "I <3 Github."
	path := "example.png"

	err := qrencode.Encode([]byte(str), path, qrencode.ECLevelQ)
	if err != nil {
		return
	}
}
```
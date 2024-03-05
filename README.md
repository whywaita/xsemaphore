# xsemaphore

xsemaphore provide semaphore.Weighted by keys.

## Usage

```go
package main

import (
	"context"
	"fmt"

	"github.com/whywaita/xsemaphore"
)

func main() {
	ctx := context.Background()
	keys := []string{"key1", "key2"}

	for _, key := range keys {
		func() {
			sem := xsemaphore.Get(key, 1)
			if err := sem.Acquire(ctx, 1); err != nil {
				fmt.Println(err)
				return
			}
			defer sem.Release(1)

			// do something
		}()
	}
}
```
// context carries deadlines, cancellation signals, and request-scoped
// values across API boundaries and goroutines. context.WithTimeout
// produces a context whose Done() channel closes when it expires or is
// cancelled.
package main

import (
	"context"
	"fmt"
	"time"
)

func work(ctx context.Context) {
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("work completed")
	case <-ctx.Done():
		fmt.Println("work cancelled:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	work(ctx)
}

package main

import (
	"context"
	"fmt"
)

func main() {
	var ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	var count = 0
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println(i, "Done")
			// break
		default:
			count++
			fmt.Println(i, count)
		}
		if i == 1 {
			cancel()
		}
	}
	_ = count
}

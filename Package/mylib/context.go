package mylib

import (
	"context"
	"fmt"
	"time"
)

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(2 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func Context() {
	ch := make(chan string)
	ctx := context.Background()
	// 3秒以上処理がかかったらキャンセルする
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	// まだ処理が残っている場合に
	// ctx := context.TODO()
	go longProcess(ctx, ch)

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
}

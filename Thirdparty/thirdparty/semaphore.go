package thirdparty

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	// このgoroutineの処理が行ってる際に他の処理をキャンセルしたい場合にtryAquireを使用する
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Coud not get lock")
		return
	}
	// if err := s.Acquire(ctx, 1); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func Semaphore() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}

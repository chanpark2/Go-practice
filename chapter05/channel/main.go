package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d canceled\n", id)
			return
		default:
			fmt.Printf("Worker %d is working\n", id)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 3*time.Second)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	time.Sleep(5 * time.Second) // 5초 대기 후 main 함수가 종료되면서 취소 요청

	fmt.Println("Main function is done")
}

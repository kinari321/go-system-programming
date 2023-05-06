package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()
	singctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	toctx, cancel2 := context.WithTimeout(ctx, time.Second*5)
	defer cancel2()

	select {
	case <-singctx.Done():
		fmt.Println("signal")
	case <-toctx.Done():
		fmt.Println("timeout")
	}

}

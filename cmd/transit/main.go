package main

import (
	"context"
	"dora/modules/api/transit"
	"dora/pkg/utils/ginutil"
	"os/signal"
	"syscall"
)

// dora cmd transit
// 数据接收服务
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	serve := transit.Serve()
	ginutil.GracefulShutdown(ctx, serve)
}

package main

import (
	"context"
	"dora/modules/api/manage"
	"dora/pkg/utils/ginutil"
	"os/signal"
	"syscall"
)

// dora manage
// 后台管理服务
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	serve := manage.Serve()
	ginutil.GracefulShutdown(ctx, serve)
}

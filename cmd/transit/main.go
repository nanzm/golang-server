package main

import (
	"dora/modules/api/transit"
)

// dora cmd transit
// 数据接收服务
func main() {
	transit.Serve()
}

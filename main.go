package main

import (
	_ "bamboo-dashboard/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"bamboo-dashboard/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

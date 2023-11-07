package main

import (
	_ "NTNU_APP_2023_Final_Backend/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"NTNU_APP_2023_Final_Backend/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}

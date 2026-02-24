package main

import (
	"context"
	"log"

	"go-zero-ai-kit/mcp/internal/server"
	"go-zero-ai-kit/mcp/internal/tools"
)

func main() {
	srv := server.New()
	srv.Register(tools.InitProjectTool())
	srv.Register(tools.GenApiSkeletonTool())

	if err := server.RunDefault(context.Background(), srv); err != nil {
		log.Fatal(err)
	}
}

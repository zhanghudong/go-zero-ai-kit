#!/bin/sh
set -e

echo "Building gozero-ai-mcp..."
go build -o gozero-ai-mcp ./mcp/cmd/gozero-ai-mcp

echo "Done: ./gozero-ai-mcp"

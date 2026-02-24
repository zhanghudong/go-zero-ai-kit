package server

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type Tool struct {
	Name        string
	Description string
	InputSchema map[string]interface{}
	Handler     func(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error)
}

type Server struct {
	tools map[string]Tool
}

func New() *Server {
	return &Server{tools: make(map[string]Tool)}
}

func (s *Server) Register(tool Tool) {
	s.tools[tool.Name] = tool
}

func (s *Server) Serve(ctx context.Context, in io.Reader, out io.Writer) error {
	reader := bufio.NewReader(in)
	writer := bufio.NewWriter(out)
	defer writer.Flush()

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if len(line) == 0 {
			continue
		}

		var req rpcRequest
		if err := json.Unmarshal(line, &req); err != nil {
			_ = writeError(writer, req.ID, -32700, "parse error")
			continue
		}

		if req.Method == "initialize" {
			_ = writeResult(writer, req.ID, map[string]interface{}{
				"protocolVersion": "2024-11-05",
				"serverInfo": map[string]interface{}{
					"name":    "gozero-ai-mcp",
					"version": "0.1.0",
				},
				"capabilities": map[string]interface{}{
					"tools": map[string]interface{}{},
				},
			})
			continue
		}

		switch req.Method {
		case "tools/list":
			list := make([]map[string]interface{}, 0, len(s.tools))
			for _, tool := range s.tools {
				list = append(list, map[string]interface{}{
					"name":        tool.Name,
					"description": tool.Description,
					"inputSchema": tool.InputSchema,
				})
			}
			_ = writeResult(writer, req.ID, map[string]interface{}{"tools": list})
		case "tools/call":
			params := map[string]interface{}{}
			if req.Params != nil {
				_ = json.Unmarshal(req.Params, &params)
			}
			name, _ := params["name"].(string)
			args, _ := params["arguments"].(map[string]interface{})
			tool, ok := s.tools[name]
			if !ok {
				_ = writeError(writer, req.ID, -32601, "tool not found")
				continue
			}

			ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
			result, err := tool.Handler(ctx, args)
			cancel()
			if err != nil {
				_ = writeError(writer, req.ID, -32000, err.Error())
				continue
			}
			_ = writeResult(writer, req.ID, map[string]interface{}{
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": fmt.Sprintf("tool %s executed", name),
					},
				},
				"structured": result,
			})
		case "shutdown":
			_ = writeResult(writer, req.ID, map[string]interface{}{})
		case "exit":
			return nil
		default:
			_ = writeError(writer, req.ID, -32601, "method not found")
		}
	}
}

type rpcRequest struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
}

type rpcResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *rpcError   `json:"error,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func writeResult(w *bufio.Writer, id interface{}, result interface{}) error {
	resp := rpcResponse{
		JSONRPC: "2.0",
		ID:      id,
		Result:  result,
	}
	data, _ := json.Marshal(resp)
	if _, err := w.Write(append(data, '\n')); err != nil {
		return err
	}
	return w.Flush()
}

func writeError(w *bufio.Writer, id interface{}, code int, msg string) error {
	resp := rpcResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error: &rpcError{
			Code:    code,
			Message: msg,
		},
	}
	data, _ := json.Marshal(resp)
	if _, err := w.Write(append(data, '\n')); err != nil {
		return err
	}
	return w.Flush()
}

func RunDefault(ctx context.Context, srv *Server) error {
	return srv.Serve(ctx, os.Stdin, os.Stdout)
}

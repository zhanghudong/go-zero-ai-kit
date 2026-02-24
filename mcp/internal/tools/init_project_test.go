package tools

import (
	"os"
	"path/filepath"
	"testing"
)

func TestInitProjectDryRunFallback(t *testing.T) {
	root := t.TempDir()
	tool := InitProjectTool()

	args := map[string]interface{}{
		"project_name": "demo",
		"module_path":  "example.com/demo",
		"service_name": "api",
		"output_dir":   root,
		"goctl_path":   "__missing_goctl__",
		"dry_run":      true,
	}

	result, err := tool.Handler(nil, args)
	if err != nil {
		t.Fatalf("handler failed: %v", err)
	}

	created := result["created_files"].([]string)
	if len(created) == 0 {
		t.Fatalf("expected created files in dry run")
	}

	if _, err := os.Stat(filepath.Join(root, "go.mod")); err == nil {
		t.Fatalf("go.mod should not exist in dry run")
	}
}

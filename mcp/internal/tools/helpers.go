package tools

import (
	"bufio"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func getString(args map[string]interface{}, key, def string) string {
	if v, ok := args[key]; ok {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return def
}

func getBool(args map[string]interface{}, key string, def bool) bool {
	if v, ok := args[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return def
}

func resolveTemplateRoot(root string) string {
	if root == "" {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, ".goctl")
	}
	if strings.HasPrefix(root, "~") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, strings.TrimPrefix(root, "~"))
	}
	return root
}

func lookPath(path string) (string, error) {
	if path == "" {
		return "", errors.New("empty path")
	}
	if strings.Contains(path, string(os.PathSeparator)) {
		return path, nil
	}
	return exec.LookPath(path)
}

func findModulePath(startDir string) (string, string) {
	dir := startDir
	for {
		goModPath := filepath.Join(dir, "go.mod")
		if module := readModulePath(goModPath); module != "" {
			return module, goModPath
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", ""
}

func readModulePath(goModPath string) string {
	file, err := os.Open(goModPath)
	if err != nil {
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}
	return ""
}

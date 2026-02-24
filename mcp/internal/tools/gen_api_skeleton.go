package tools

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"go-zero-ai-kit/mcp/internal/fsutil"
	"go-zero-ai-kit/mcp/internal/server"
)

func GenApiSkeletonTool() server.Tool {
	return server.Tool{
		Name:        "gen_api_skeleton",
		Description: "Generate .api and handler/logic/types skeleton",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"api_name":      map[string]interface{}{"type": "string"},
				"base_path":     map[string]interface{}{"type": "string"},
				"output_dir":    map[string]interface{}{"type": "string"},
				"template_root": map[string]interface{}{"type": "string"},
				"goctl_path":    map[string]interface{}{"type": "string"},
				"style":         map[string]interface{}{"type": "string"},
				"force":         map[string]interface{}{"type": "boolean"},
				"dry_run":       map[string]interface{}{"type": "boolean"},
			},
			"required": []string{"api_name", "base_path", "output_dir"},
		},
		Handler: func(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
			apiName := getString(args, "api_name", "example")
			basePath := getString(args, "base_path", "/api/v1")
			outputDir := getString(args, "output_dir", ".")
			templateRoot := resolveTemplateRoot(getString(args, "template_root", ""))
			goctlPath := getString(args, "goctl_path", "goctl")
			style := getString(args, "style", "")
			force := getBool(args, "force", false)
			dryRun := getBool(args, "dry_run", false)

			result := &GenerateResult{}

			modulePath, goModPath := findModulePath(outputDir)
			if modulePath == "" {
				modulePath = "go-zero-ai-kit"
				result.addWarning("module path not found; using default")
			} else {
				_ = goModPath
			}

			if !dryRun {
				used, warn := tryGoctlGenApi(ctx, goctlPath, templateRoot, apiName, basePath, outputDir, modulePath, style, force)
				if warn != "" {
					result.addWarning(warn)
				}
				if used {
					return mapResult(result), nil
				}
			}

			packageName := toPackageName(apiName)
			routesPath := filepath.Join(outputDir, "internal", "handler", "routes.go")
			includeRoutes := true
			if exists, _ := fsutil.FileExists(routesPath); exists {
				includeRoutes = false
				result.addWarning("routes.go already exists; fallback did not update routes")
			}

			data := TemplateData{
				ApiName:     apiName,
				BasePath:    basePath,
				ModulePath:  modulePath,
				PackageName: packageName,
			}
			opts := fsutil.WriteOptions{Force: force, DryRun: dryRun}
			if err := renderTemplateSet(genApiSkeletonTemplateSpecs(includeRoutes), data, opts, outputDir, result); err != nil {
				return nil, err
			}
			return mapResult(result), nil
		},
	}
}

func tryGoctlGenApi(ctx context.Context, goctlPath, templateRoot, apiName, basePath, outputDir, modulePath, style string, force bool) (bool, string) {
	path, err := lookPath(goctlPath)
	if err != nil {
		return false, "goctl not found; fallback templates used"
	}

	apiDir := filepath.Join(outputDir, "api")
	apiFile := filepath.Join(apiDir, fmt.Sprintf("%s.api", apiName))
	if err := os.MkdirAll(apiDir, 0o755); err != nil {
		return false, fmt.Sprintf("failed to create api dir: %s", err)
	}

	if !force {
		if exists, _ := fsutil.FileExists(apiFile); exists {
			return false, "api file exists and force=false; skipping goctl"
		}
	}

	data := TemplateData{ApiName: apiName, BasePath: basePath}
	content, _, err := renderTemplate(TemplateSpec{TemplatePath: "gen_api_skeleton/api.api.tmpl", DestPathPattern: "api"}, data)
	if err != nil {
		return false, fmt.Sprintf("failed to render api template: %s", err)
	}

	if err := os.WriteFile(apiFile, content, 0o644); err != nil {
		return false, fmt.Sprintf("failed to write api file: %s", err)
	}

	args := []string{"api", "go", "-api", apiFile, "-dir", outputDir}
	if templateRoot != "" {
		args = append(args, "--home", templateRoot)
	}
	if style != "" {
		args = append(args, "-style", style)
	}

	cmd := exec.CommandContext(ctx, path, args...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("GOCTL_HOME=%s", templateRoot))
	output, err := cmd.CombinedOutput()
	if err != nil {
		msg := strings.TrimSpace(string(output))
		if msg == "" {
			msg = err.Error()
		}
		return false, fmt.Sprintf("goctl failed (%s); fallback templates used", msg)
	}
	_ = modulePath
	return true, ""
}

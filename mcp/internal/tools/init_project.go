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

func InitProjectTool() server.Tool {
	return server.Tool{
		Name:        "init_project",
		Description: "Initialize a go-zero API project",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"project_name":  map[string]interface{}{"type": "string"},
				"module_path":   map[string]interface{}{"type": "string"},
				"service_name":  map[string]interface{}{"type": "string"},
				"output_dir":    map[string]interface{}{"type": "string"},
				"template_root": map[string]interface{}{"type": "string"},
				"goctl_path":    map[string]interface{}{"type": "string"},
				"force":         map[string]interface{}{"type": "boolean"},
				"dry_run":       map[string]interface{}{"type": "boolean"},
			},
			"required": []string{"project_name", "module_path", "service_name", "output_dir"},
		},
		Handler: func(ctx context.Context, args map[string]interface{}) (map[string]interface{}, error) {
			projectName := getString(args, "project_name", "gozero-app")
			modulePath := getString(args, "module_path", "gozero-app")
			serviceName := getString(args, "service_name", "api")
			outputDir := getString(args, "output_dir", ".")
			templateRoot := resolveTemplateRoot(getString(args, "template_root", ""))
			goctlPath := getString(args, "goctl_path", "goctl")
			force := getBool(args, "force", false)
			dryRun := getBool(args, "dry_run", false)

			result := &GenerateResult{}
			if !force {
				if info, err := os.Stat(outputDir); err == nil && info.IsDir() {
					entries, _ := os.ReadDir(outputDir)
					if len(entries) > 0 {
						result.addWarning("output_dir exists and is not empty; skipping")
						return mapResult(result), nil
					}
				}
			}

			if !dryRun {
				used, warn := tryGoctlInitProject(ctx, goctlPath, templateRoot, serviceName, outputDir)
				if warn != "" {
					result.addWarning(warn)
				}
				if used {
					return mapResult(result), nil
				}
			}

			data := TemplateData{
				ProjectName: projectName,
				ModulePath:  modulePath,
				ServiceName: serviceName,
			}
			opts := fsutil.WriteOptions{Force: force, DryRun: dryRun}
			if err := renderTemplateSet(initProjectTemplateSpecs(), data, opts, outputDir, result); err != nil {
				return nil, err
			}

			return mapResult(result), nil
		},
	}
}

func tryGoctlInitProject(ctx context.Context, goctlPath, templateRoot, serviceName, outputDir string) (bool, string) {
	path, err := lookPath(goctlPath)
	if err != nil {
		return false, "goctl not found; fallback templates used"
	}

	args := []string{"api", "new", serviceName, "--dir", outputDir}
	if templateRoot != "" {
		args = append(args, "--home", templateRoot)
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
	return true, ""
}

func mapResult(r *GenerateResult) map[string]interface{} {
	return map[string]interface{}{
		"created_files":     r.CreatedFiles,
		"skipped_files":     r.SkippedFiles,
		"overwritten_files": r.OverwrittenFiles,
		"warnings":          r.Warnings,
	}
}

package tools

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"go-zero-ai-kit/mcp/internal/fsutil"
	"go-zero-ai-kit/mcp/internal/templates"
)

type TemplateSpec struct {
	TemplatePath    string
	DestPathPattern string
	GoFormat        bool
}

type TemplateData struct {
	ProjectName string
	ModulePath  string
	ServiceName string
	ApiName     string
	BasePath    string
	PackageName string
}

func renderTemplate(spec TemplateSpec, data TemplateData) ([]byte, string, error) {
	tmplBytes, err := templates.FS.ReadFile(spec.TemplatePath)
	if err != nil {
		return nil, "", err
	}

	pathTmpl, err := template.New("path").Parse(spec.DestPathPattern)
	if err != nil {
		return nil, "", err
	}
	var pathBuf bytes.Buffer
	if err := pathTmpl.Execute(&pathBuf, data); err != nil {
		return nil, "", err
	}

	contentTmpl, err := template.New("content").Parse(string(tmplBytes))
	if err != nil {
		return nil, "", err
	}
	var contentBuf bytes.Buffer
	if err := contentTmpl.Execute(&contentBuf, data); err != nil {
		return nil, "", err
	}

	content := contentBuf.Bytes()
	if spec.GoFormat {
		content = fsutil.FormatGoSource(content)
	}

	return content, filepath.FromSlash(pathBuf.String()), nil
}

func renderTemplateSet(specs []TemplateSpec, data TemplateData, opts fsutil.WriteOptions, baseDir string, result *GenerateResult) error {
	for _, spec := range specs {
		content, relPath, err := renderTemplate(spec, data)
		if err != nil {
			return err
		}
		destPath := filepath.Join(baseDir, relPath)
		writeRes, err := fsutil.WriteFile(destPath, content, opts)
		if err != nil {
			return err
		}

		norm := fsutil.NormalizePath(destPath)
		switch writeRes.Status {
		case fsutil.StatusCreated:
			result.addCreated(norm)
		case fsutil.StatusOverwritten:
			result.addOverwritten(norm)
		case fsutil.StatusSkipped:
			result.addSkipped(norm)
		default:
			result.addWarning(fmt.Sprintf("unknown write status for %s", norm))
		}
	}
	return nil
}

func initProjectTemplateSpecs() []TemplateSpec {
	return []TemplateSpec{
		{TemplatePath: "init_project/go.mod.tmpl", DestPathPattern: "go.mod", GoFormat: false},
		{TemplatePath: "init_project/README.md.tmpl", DestPathPattern: "README.md", GoFormat: false},
		{TemplatePath: "init_project/cmd_main.go.tmpl", DestPathPattern: "cmd/{{.ServiceName}}/main.go", GoFormat: true},
		{TemplatePath: "init_project/config.go.tmpl", DestPathPattern: "internal/config/config.go", GoFormat: true},
		{TemplatePath: "init_project/service_context.go.tmpl", DestPathPattern: "internal/svc/service_context.go", GoFormat: true},
		{TemplatePath: "init_project/types.go.tmpl", DestPathPattern: "internal/types/types.go", GoFormat: true},
		{TemplatePath: "init_project/handler_routes.go.tmpl", DestPathPattern: "internal/handler/routes.go", GoFormat: true},
		{TemplatePath: "init_project/handler_ping.go.tmpl", DestPathPattern: "internal/handler/ping_handler.go", GoFormat: true},
		{TemplatePath: "init_project/logic_ping.go.tmpl", DestPathPattern: "internal/logic/ping_logic.go", GoFormat: true},
		{TemplatePath: "init_project/etc.yaml.tmpl", DestPathPattern: "etc/{{.ServiceName}}.yaml", GoFormat: false},
		{TemplatePath: "init_project/api.api.tmpl", DestPathPattern: "api/{{.ServiceName}}.api", GoFormat: false},
	}
}

func genApiSkeletonTemplateSpecs(includeRoutes bool) []TemplateSpec {
	specs := []TemplateSpec{
		{TemplatePath: "gen_api_skeleton/api.api.tmpl", DestPathPattern: "api/{{.ApiName}}.api", GoFormat: false},
		{TemplatePath: "gen_api_skeleton/handler_get.go.tmpl", DestPathPattern: "internal/handler/{{.PackageName}}/get_{{.ApiName}}_handler.go", GoFormat: true},
		{TemplatePath: "gen_api_skeleton/handler_post.go.tmpl", DestPathPattern: "internal/handler/{{.PackageName}}/create_{{.ApiName}}_handler.go", GoFormat: true},
		{TemplatePath: "gen_api_skeleton/logic_get.go.tmpl", DestPathPattern: "internal/logic/{{.PackageName}}/get_{{.ApiName}}_logic.go", GoFormat: true},
		{TemplatePath: "gen_api_skeleton/logic_post.go.tmpl", DestPathPattern: "internal/logic/{{.PackageName}}/create_{{.ApiName}}_logic.go", GoFormat: true},
		{TemplatePath: "gen_api_skeleton/types.go.tmpl", DestPathPattern: "internal/types/{{.ApiName}}.go", GoFormat: true},
	}
	if includeRoutes {
		specs = append(specs, TemplateSpec{TemplatePath: "gen_api_skeleton/routes.go.tmpl", DestPathPattern: "internal/handler/routes.go", GoFormat: true})
	}
	return specs
}

func toPackageName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ReplaceAll(name, "-", "_")
	name = strings.ReplaceAll(name, ".", "_")
	if name == "" {
		return "api"
	}
	return name
}

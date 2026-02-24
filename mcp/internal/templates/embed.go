package templates

import "embed"

//go:embed init_project/*.tmpl gen_api_skeleton/*.tmpl
var FS embed.FS

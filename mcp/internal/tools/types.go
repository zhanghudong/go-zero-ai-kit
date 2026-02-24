package tools

type GenerateResult struct {
	CreatedFiles     []string `json:"created_files"`
	SkippedFiles     []string `json:"skipped_files"`
	OverwrittenFiles []string `json:"overwritten_files"`
	Warnings         []string `json:"warnings"`
}

func (r *GenerateResult) addCreated(path string) {
	r.CreatedFiles = append(r.CreatedFiles, path)
}

func (r *GenerateResult) addSkipped(path string) {
	r.SkippedFiles = append(r.SkippedFiles, path)
}

func (r *GenerateResult) addOverwritten(path string) {
	r.OverwrittenFiles = append(r.OverwrittenFiles, path)
}

func (r *GenerateResult) addWarning(msg string) {
	r.Warnings = append(r.Warnings, msg)
}

package fsutil

import (
	"errors"
	"go/format"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type WriteOptions struct {
	Force  bool
	DryRun bool
}

type WriteStatus string

const (
	StatusCreated     WriteStatus = "created"
	StatusOverwritten WriteStatus = "overwritten"
	StatusSkipped     WriteStatus = "skipped"
)

type WriteResult struct {
	Path   string
	Status WriteStatus
}

func WriteFile(path string, data []byte, opts WriteOptions) (WriteResult, error) {
	exists, err := FileExists(path)
	if err != nil {
		return WriteResult{}, err
	}

	if exists && !opts.Force {
		return WriteResult{Path: path, Status: StatusSkipped}, nil
	}

	if opts.DryRun {
		status := StatusCreated
		if exists {
			status = StatusOverwritten
		}
		return WriteResult{Path: path, Status: status}, nil
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return WriteResult{}, err
	}

	if err := os.WriteFile(path, data, 0o644); err != nil {
		return WriteResult{}, err
	}

	status := StatusCreated
	if exists {
		status = StatusOverwritten
	}
	return WriteResult{Path: path, Status: status}, nil
}

func FileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return !info.IsDir(), nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func FormatGoSource(data []byte) []byte {
	formatted, err := format.Source(data)
	if err != nil {
		return data
	}
	return formatted
}

func NormalizePath(path string) string {
	clean := filepath.Clean(path)
	return strings.ReplaceAll(clean, string(os.PathSeparator), "/")
}

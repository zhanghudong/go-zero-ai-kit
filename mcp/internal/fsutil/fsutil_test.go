package fsutil

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFileStatuses(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "file.txt")

	res, err := WriteFile(path, []byte("one"), WriteOptions{})
	if err != nil {
		t.Fatalf("write failed: %v", err)
	}
	if res.Status != StatusCreated {
		t.Fatalf("expected created, got %s", res.Status)
	}

	res, err = WriteFile(path, []byte("two"), WriteOptions{})
	if err != nil {
		t.Fatalf("write failed: %v", err)
	}
	if res.Status != StatusSkipped {
		t.Fatalf("expected skipped, got %s", res.Status)
	}

	res, err = WriteFile(path, []byte("three"), WriteOptions{Force: true})
	if err != nil {
		t.Fatalf("write failed: %v", err)
	}
	if res.Status != StatusOverwritten {
		t.Fatalf("expected overwritten, got %s", res.Status)
	}
}

func TestWriteFileDryRun(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "dry.txt")

	res, err := WriteFile(path, []byte("data"), WriteOptions{DryRun: true})
	if err != nil {
		t.Fatalf("write failed: %v", err)
	}
	if res.Status != StatusCreated {
		t.Fatalf("expected created in dry run, got %s", res.Status)
	}
	if _, err := os.Stat(path); err == nil {
		t.Fatalf("file should not exist in dry run")
	}
}

package test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/madlambda/spells/assert"
)

// TempDir creates a temporary directory.
func TempDir(t *testing.T, base string) string {
	t.Helper()

	dir, err := ioutil.TempDir(base, "terrastack-test")
	assert.NoError(t, err, "creating temp directory")
	return dir
}

// WriteFile writes content to a filename inside dir directory.
// If dir is empty string then the file is created inside a temporary directory.
func WriteFile(t *testing.T, dir string, filename string, content string) string {
	t.Helper()

	if dir == "" {
		dir = TempDir(t, "")
	}

	path := filepath.Join(dir, filename)
	err := ioutil.WriteFile(path, []byte(content), 0700)
	assert.NoError(t, err, "writing test file %s", path)

	return path
}

// MkdirAll creates a temporary directory with default test permission bits.
func MkdirAll(t *testing.T, path string) {
	assert.NoError(t, os.MkdirAll(path, 0700), "failed to create temp directory")
}
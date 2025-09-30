package dl

import (
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func compileSo(ctx context.Context, path string) error {
	return exec.CommandContext(
		ctx,
		"gcc",
		"-fPIC",
		"-shared",
		"-o",
		filepath.Join(path, "testdata", "plugin.so"),
		filepath.Join(path, "testdata", "plugin.c"),
	).Run()
}

func TestCompile(t *testing.T) {
	path, _ := os.Getwd()
	err := compileSo(t.Context(), path)
	assert.NoError(t, err)

	lib, err := Open(filepath.Join(path, "testdata", "plugin.so"))
	assert.NoError(t, err)
	defer lib.Release()

	f, err := lib.Func("sample_buffer_function")
	assert.NoError(t, err)

	buf := make([]byte, 256)

	result := testFunc(f, buf)

	assert.Equal(t, 11, result)
	assert.Equal(t, "hello world", string(buf[:result]))
}
